Reposted from **https://joeshaw.org/revisiting-context-and-http-handler-for-go-17/**

Go 1.7 was released earlier this month, and the thing I’m most excited about is the incorporation of the context package into the Go standard library. Previously it lived in the ```golang.org/x/net/context``` package.

With the move, other packages within the standard library can now use it. The net package’s Dialer and os/exec package’s Command can now utilize contexts for easy cancelation. More on this can be found in the Go 1.7 release notes.

Go 1.7 also brings contexts to the net/http package’s Request type for both HTTP clients and servers. Last year I wrote a post about using context.Context with http.Handler when it lived outside the standard library, but Go 1.7 makes things much simpler and thankfully renders all of the approaches from that post obsolete.
A quick recap

I suggest reading my original post for more background, but one of the main uses of context.Context is to pass around request-scoped data. Things like request IDs, authenticated user information, and other data useful for handlers and middleware to examine in the scope of a single HTTP request.

In that post I examined three different approaches for incorporating context into requests. Since contexts are now attached to http.Request values, this is no longer necessary. As long as you’re willing to require at least Go 1.7, it’s now possible to use the standard http.Handler interface and common middleware patterns with ```context.Context```!

The new approach

Recall that the http.Handler interface is defined as:

```go
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}
```

Go 1.7 adds new context-related methods on the ```*http.Request``` type.

```go
func (r *Request) Context() context.Context
func (r *Request) WithContext(ctx context.Context) *Request
```

The Context method returns the current context associated with the request. The WithContext method creates a new Request value with the provided context.

Suppose we want each request to have an associated ID, pulling it from the X-Request-ID HTTP header if present, and generating it if not. We might implement the context functions like this:

```go
type key int
const requestIDKey key = 0

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
    reqID := req.Header.Get("X-Request-ID")
    if reqID == "" {
        reqID = generateRandomID()
    }

    return context.WithValue(ctx, requestIDKey, reqID)
}

func requestIDFromContext(ctx context.Context) string {
    return ctx.Value(requestIDKey).(string)
}
```

We can implement middleware that derives a new context with a request ID, create a new Request value from it, and pass it onto the next handler in the chain.

```go
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        ctx := newContextWithRequestID(req.Context(), req)
        next.ServeHTTP(rw, req.WithContext(ctx))
    })
}
```

The final handler and any middleware lower in the chain have access to all the previously request-scoped data set in middleware above it.

```go
func handler(rw http.ResponseWriter, req *http.Request) {
    reqID := requestIDFromContext(req.Context())
    fmt.Fprintf(rw, "Hello request ID %v\n", reqID)
}
```
And that’s it! It’s no longer necessary to implement custom context handlers, adapters to standard http.Handler implementations, or hackily wrap http.ResponseWriter. Everything you need is in the standard library, and right there on the ```*http.Request``` type.
