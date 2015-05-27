//Courtesy: Alexey Kachayev
//https://gist.github.com/kachayev/21e7fe149bc5ae0bd878

// The Task:
// get twitter username
// read users' tweets (all of them or with given limit)
// calculate "audience" for each tweet: number of uniq users how saw each tweet
// (set union of users' followers and all retweeters' followers)

// The Idea:
// pool of tweeter users (we need auth data for each request)
// with shared "queues" for requests: read timeline, get retweets, find followers
// 3 actors: TimelineReader, RetweetsReader, FollowersReader
// shared channels for communication: tweet ids, user screen names
// TimelineReader: sequentially read tweets, emit IDs to RetweetsReader
// RetweetsReader: for each ID fetch list of retweets, emit all mentioned user to FollowersReader
// FollowersReader: for each user name fetch followers, keep cache of already seen users

// The ugly part: how to find that everything is done?
// wait group + stop channel + waiting true|false flag
// why use single "done" channel for actor instead of separated
// channel for each request? cause we handle map[string]strig for
// results. and it means that only top-level actor can normally deal
// with responses from other actors.
// You can find more ugly parts in comments:
// * how to implement abstract type-safe Pool of Worker(s)?
// * how to implement abstract type-safe Future of something?
// * how to avoid code duplicates for Request/Reply pattern without getting type checker cry?
// * how to solve backpressure problem without deadlocks and suicides?

package main

import (
	"runtime"
	"log"
	"flag"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"fmt"
	"strings"
	"time"
	"sort"
	"sync"
)

var (
	tkey = "<KEY>"
	tsecret = "<SECRET>"
	tokens = []string{
           "<TOKEN1>:<SECRET1>"
	}
)

type AnalyzedTweet struct {
	Id int64
	Text string
	Retweets int
	Favorites int
	Audience int
}

type Reader struct {
	Index int
	Client *anaconda.TwitterApi
	Limit int
	Requests chan Request
}

// xxx: note, that there is no way how to describe
// list of concrete types. something similar to
// type Request = QueryRequest | RetweetRequest | FollowerRequest
type Request interface {
	Do(*Reader)
}

type QueryRequest struct {
	Username string
	MaxId int64
	Left int
	Requester chan []anaconda.Tweet
}

func (rq QueryRequest) Do(r *Reader) {
	log.Printf("perform query @%s from id %d", rq.Username, rq.MaxId)
	// xxx: need to limit number of requests somehow
	// xxx: note, that you probably do not need to
	// use goroutine here, you need to create special
	// queue procedure to read and keep all requests,
	// before worker is ready to take new task
	go func(){
		v := url.Values{}
		v.Set("count", fmt.Sprintf("%d", min(rq.Left, 1000)))
		if rq.MaxId != 0 {
			v.Set("max_id", fmt.Sprintf("%d", rq.MaxId))
		}
		v.Set("screen_name", strings.TrimPrefix(rq.Username, "@"))
		results, err := r.Client.GetUserTimeline(v)

		if err != nil {
			log.Printf("Error (%d): %s", r.Index, err.Error())
		}

		rq.Requester <- results
	}()
}

type RetweetRequest struct {
	Id int64
	Requester chan *RetweetResponse
}

func (rq RetweetRequest) Do(r *Reader) {
	log.Printf("looking retweets for #%d", rq.Id)

	// xxx: need to limit number of requests somehow
	go func() {
		v := url.Values{}
		v.Set("count", "1000")
		result, err := r.Client.GetRetweets(rq.Id, v)

		if err != nil {
			log.Printf("Error (%d): %s", r.Index, err.Error())
		}

		log.Printf("fetch %d retweets for #%d", len(result), rq.Id)

		rq.Requester <- &RetweetResponse{rq.Id, result}
	}()
}

type RetweetResponse struct {
	Id int64
	Tweets []anaconda.Tweet
}

type FollowerRequest struct {
	Name string
	Requester chan *FollowerResponse
}

func (rq FollowerRequest) Do(r *Reader) {
	log.Printf("looking followers for @%s", rq.Name)

	// xxx: need to limit number of requests somehow
	go func() {
		v := url.Values{}
		v.Set("screen_name", rq.Name)
		v.Set("count", "200")
		result := r.Client.GetFollowersListAll(v)

		followers := []anaconda.User{}
		// xxx: not the best approach... reading too many pages
		// we will block current worker cause of API calls
		// limitation, so it's better to expose channel outside
		// and send tasks to worker pool to read independent pages
		// (we know number of followers, so easily can do this)
		for page := range result {
			followers = append(followers, page.Followers...)
		}

		log.Printf("fetch %d followers for %s", len(followers), rq.Name)

		rq.Requester <- &FollowerResponse{rq.Name, followers}
	}()
}

type FollowerResponse struct {
	Name string
	Users []anaconda.User
}

type TwitterGateway struct {
	Key string
	Secret string
	Readers []*Reader
	Requests chan Request
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func (r *Reader) Run() {
	for rq := range r.Requests {
		rq.Do(r)
	}
}

func NewTwitterGateway(key string, secret string, users []string) *TwitterGateway {
	anaconda.SetConsumerKey(key)
	anaconda.SetConsumerSecret(secret)

	queue := make(chan Request, 5000)

	readers := []*Reader{}
	for i, fullToken := range users {
		parts := strings.Split(fullToken, ":")
		apiClient := anaconda.NewTwitterApi(parts[0], parts[1])
		r := &Reader{i+1, apiClient, 450, queue}
		go r.Run()
		readers = append(readers, r)
	}

	log.Println("==> initialized ", len(readers), " readers")

	return &TwitterGateway{key, secret, readers, queue}
}

func minId(tweets []anaconda.Tweet) int64 {
	m := int64(0)
	for _, t := range tweets {
		if t.Id < m || m == 0 {
			m = t.Id
		}
	}

	return m
}

func (tg *TwitterGateway) TimelineReader(username string, limit int, rtc chan int64) []anaconda.Tweet {
	userTweets := []anaconda.Tweet{}
	maxId := int64(0)

	for {
		if limit != 0 && len(userTweets) >= limit {
			break
		}

		resp := make(chan []anaconda.Tweet)
		tg.Requests <- QueryRequest{username, maxId, limit-len(userTweets), resp}
		tweets := <- resp
		if len(tweets) == 0 {
			break
		}

		cursor := maxId
		id := minId(tweets)
		if id < maxId || maxId == 0 {
			maxId = id
		}

		if cursor == maxId {
			break
		}

		for _, t := range tweets {
			if !strings.HasPrefix(t.Text, "RT ") {
				// ownRetweets += t.RetweetCount
				// ownFavorites += t.FavoriteCount
				rtc <- t.Id
			}
			// if t.RetweetCount > 5 && !strings.HasPrefix(t.Text, "RT ") {
			// 	log.Printf("%d retweets: %s", t.RetweetCount, t.Text)
			// }
		}

		userTweets = append(userTweets, tweets...)
	}

	close(rtc)
	log.Println("==> closing retweets channel")
	return userTweets
}

func (tg *TwitterGateway) RetweetsReader(rtc chan int64, fc chan string) map[int64][]anaconda.Tweet {
	retweets := map[int64][]anaconda.Tweet{}
	done := make(chan *RetweetResponse)
	stop := make(chan bool)
	waiting := false
	var tasks sync.WaitGroup

loop:
	for {
		select {
		case id, ok := <- rtc:
			if ok {
				tasks.Add(1)
				// xxx: off course it's much easier to use separated
				// channel for results instead of single "done" with
				// the same select loop. but the problem is shared
				// retweets map that we can't safely update from
				// child gorouting (at least without locks)
				tg.Requests <- RetweetRequest{id, done}
			} else if !waiting {
				waiting = true
				go func(){
					tasks.Wait()
					close(stop)
				}()
			}
		case rt := <- done:
			retweets[rt.Id] = rt.Tweets
			for _, t := range rt.Tweets {
				// xxx: slow client / slow server?
				log.Printf("ask for followers for @%s", t.User.ScreenName)
				fc <- t.User.ScreenName
			}
			tasks.Done()
		case <-stop:
			close(fc)
			log.Println("==> closing followers channel")
			break loop
		}
	}

	return retweets
}

func (tg *TwitterGateway) FollowersReader(username string, fc chan string) map[string][]string {
	followers := map[string][]string{username: []string{}}
	done := make(chan *FollowerResponse)
	stop := make(chan bool)
	waiting := false

	var tasks sync.WaitGroup
	tasks.Add(1)
	tg.Requests <- FollowerRequest{username, done}

loop:
	for {
		select {
		case name, ok := <- fc:
			if ok {
				if _, present := followers[name]; !present {
					tasks.Add(1)
					followers[name] = []string{}
					tg.Requests <- FollowerRequest{name, done}
				}
			} else if !waiting {
				waiting = true
				go func() {
					tasks.Wait()
					close(stop)
				}()
			}
		case fls := <- done:
			users := []string{}
			for _, u := range fls.Users {
				users = append(users, u.ScreenName)
			}

			followers[fls.Name] = users
			tasks.Done()
		case <-stop:
			break loop
		}
	}

	return followers
}

type Pair struct {
	Key string
	Value  int
}

type ByValue []Pair

// xxx: please, kill me
func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value > a[j].Value }

// xxx: please, kill me twice
type ByAudience []AnalyzedTweet

func (a ByAudience) Len() int           { return len(a) }
func (a ByAudience) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAudience) Less(i, j int) bool { return a[i].Audience > a[j].Audience }


func (tg *TwitterGateway) AnalyzeUser(username string, limit int) {
	start := time.Now()

	// xxx: actually it's an actor, we just don't try to
	// attract someone's attention. there are 3 actors:
	// 1) TimelineReader (emit tweet IDs to RetweetsReader)
	// 2) RetweetsReader (accept IDs and emits Names to FollowerReader)
	// 3) FollowersReader (accept ScreenNames and keep local cache)
	// use sync.WaitGroup to track when to stop
	var (
		allActors sync.WaitGroup
		userTweets []anaconda.Tweet
		retweets map[int64][]anaconda.Tweet
		followers map[string][]string
	)

	retweetsChan := make(chan int64)
	followersChan := make(chan string)

	allActors.Add(3)
	go func() {
		// never use waitgroup inside "actor", it leads to high coupling
		defer allActors.Done()
		userTweets = tg.TimelineReader(username, limit, retweetsChan)
	}()

	go func() {
		defer allActors.Done()
		retweets = tg.RetweetsReader(retweetsChan, followersChan)
	}()

	go func() {
		defer allActors.Done()
		followers = tg.FollowersReader(username, followersChan)
	}()

	allActors.Wait()

	totalRetweets := 0
	ownRetweets := 0
	totalFavorites := 0
	retweetFrequency := map[string]int{} // user -> number of retweets

	analyzedTweets := []AnalyzedTweet{}

	for _, t := range userTweets {
		totalRetweets += t.RetweetCount
		totalFavorites += t.FavoriteCount

		if !strings.HasPrefix(t.Text, "RT ") {
			ownRetweets += t.RetweetCount

			// calculate audience
			audience := map[string]bool{}
			for _, f := range followers[username] {
				audience[f] = true
			}

			for _, frt := range retweets[t.Id] {
				c, ok := retweetFrequency[frt.User.ScreenName]
				if !ok {
					c = 0
				}

				retweetFrequency[frt.User.ScreenName] = c + 1

				for _, f := range followers[frt.User.ScreenName] {
					audience[f] = true
				}
			}

			text := t.Text
			if len(text) > 100 {
				text = text[0:100] + "..."
			}

			at := AnalyzedTweet{
				t.Id,
				text,
				t.RetweetCount,
				t.FavoriteCount,
				len(audience),
			}

			analyzedTweets = append(analyzedTweets, at)
		}
	}

	pairFrequency := []Pair{}
	for name, count := range retweetFrequency {
		pairFrequency = append(pairFrequency, Pair{name, count})
	}
	// xxx: please, kill me once again
	sort.Sort(ByValue(pairFrequency))
	// xxx: now!
	sort.Sort(ByAudience(analyzedTweets))

	fmt.Println("-----------------------------")
	fmt.Println("Elapsed time:", time.Since(start))
	fmt.Println("Tweets:", len(userTweets))
	fmt.Println("Retweets:", totalRetweets)
	fmt.Println("Own retweets:", ownRetweets)
	fmt.Println("Favorites:", totalFavorites)
	fmt.Println("Analyzed tweets:", len(analyzedTweets))
	fmt.Println("-----------------------------")
	fmt.Println("Retweeters:")
	for _, p := range pairFrequency {
		fmt.Printf("%d\t@%s\t%d\n", p.Value, p.Key, len(followers[p.Key]))
	}
	fmt.Println("-----------------------------")
	fmt.Println("Tweets:")
	for _, at := range analyzedTweets {
		if at.Retweets > 0 {
			fmt.Printf("#%d\trt %d\tfav %d\taud %d\t%s\n", at.Id, at.Retweets, at.Favorites, at.Audience, at.Text)
		}
	}
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	log.Printf("Set GOMAXPROCS to: %d", num)

	tg := NewTwitterGateway(tkey, tsecret, tokens)

	username := ""
	limit := 0
	flag.StringVar(&username, "username", username, "twitter user screen name")
	flag.IntVar(&limit, "limit", limit, "when to stop")
	flag.Parse()

	if username == "" {
		log.Fatal("Error: provide username to work with")
	}

	log.Printf("==> collecting tweets for @%s", username)
	tg.AnalyzeUser(username, limit)
}
