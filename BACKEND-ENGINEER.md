**So you want to be a back-end engineer?**

Courtesy of Segment (https://segment.com/jobs/descriptions/backend-engineer)

If you’re a fan of distributed systems and like to stress-test a new database every other week, you’ll fit right in. And it won’t hurt if you like thinking about queuing topologies and love open-sourcing your work.

Our infrastructure runs on Go, with a sprinkling of Redis, Mongo, and NSQ. We’re building a service oriented system, built atop Docker containers and a fleet of microservices. We’re using ECS in production, and scripting our entire setup using Terraform.

Think that might tickle your fancy? We’ve got a few projects in the pipeline that you can sink your teeth into right away:

**Strong event ordering**

Right now, Segment is stateless. Our users send data to us and we pass it through. Most events get sent through with timestamps, but sometimes ordering is important.

Keeping a global ordering isn’t an easy problem. What happens if an event gets dropped here or there? How big of a window does it make sense to keep? What if an integration goes down for hours at a time?

**Enrichment integrations**

We’re collecting a lot of data passing through our system, and one of the top requests from our users is how to enrich the information they’re sending with external data. It’s common to say: “I want to detect the geolocation from the ip address.” or “How can I parse the user agent and detect the mobile referrer?”.

It ends up being a bunch of transforms pulled from 3rd party services into a gigantic DAG.

But there’s a lot of complication there. The network isn’t reliable, services fail, and processing can take minutes or even hours. Sometimes the information to enrich a call is necessary, but sometimes it’s not.

You’ll help design the architecture that takes all these requirements and handles them elegantly to get both good throughput and correctness.

