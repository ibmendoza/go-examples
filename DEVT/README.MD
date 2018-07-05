**Software Services**

*Invocation*

- one-time invocation - call a function/script, process, optionally return result, exit
- multiple invocation - call a function, service, API multiple times

*Lifetime*

- recurring (cron jobs) - multiple invocation, exit
- [long-running](https://medium.com/@itmarketplace.net/web-services-vs-microservices-7d2230480329) - multiple invocation, respawn on exit - Web services, microservices

**control plane vs data plane**

- error handling (cp), idempotent functions (dp)
- error handling at the control plane (either return error to end-users or to dev/ops), error handling at the data plane return [errors as values](https://blog.golang.org/errors-are-values) or handle them as exceptions (throw...catch thing) but within an idempotent function only (for example, goroutine, actor, etc)
- control plane is stateful, data plane is stateless
- if an error needs to bubble up to the end-user, return an error to control plane for handling otherwise log to STDOUT, etc

