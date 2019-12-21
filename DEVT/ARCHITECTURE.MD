**Front-end**

- [SvelteJS](https://svelte.dev)

**Backend**

pic courtesy of [LogRocket](https://blog.logrocket.com/the-perfect-architecture-flow-for-your-next-node-js-project/)

<img src="https://itjumpstart.files.wordpress.com/2019/12/business-logic-api-routes.png">

1. Controller 

- API routes
- error handling 
- view - [SapperJS](https://sapper.svelte.dev)
- data - [JSON API](https://dri.es/headless-cms-rest-vs-jsonapi-vs-graphql)

2. Service Layer (aka Business Logic Layer)

- [Workflow](https://news.ycombinator.com/item?id=19732447) (state management) and Task orchestration (can be [long running](https://blog.bernd-ruecker.com/what-are-long-running-processes-b3ee769f0a27)) aka control plane - [Uber Cadence](https://cadenceworkflow.io) 
- Data plane (Using [Dapr](https://dapr.io), pull data from RPC or push data to asynchronous task queue like NATS, Kafka )

3. Data Access Layer (via [Vault](https://www.vaultproject.io/docs/secrets/databases/index.html)-backed DAL proxy)

- read data from SQL/NoSQL stores
- write data to SQL/NoSQL stores 
