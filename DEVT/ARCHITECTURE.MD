- https://google.github.io/eng-practices/

**Front-end**

- [SvelteJS](https://svelte.dev)

**Backend**

pic courtesy of [LogRocket](https://blog.logrocket.com/the-perfect-architecture-flow-for-your-next-node-js-project/)

<img src="https://itjumpstart.files.wordpress.com/2019/12/business-logic-api-routes.png">

Backend architecture works regardless if you are using Kubernetes or not

**1. Controller**

- API gateway - Ambassador (k8s-specific), Traefik
- API routes/views - [SapperJS](https://sapper.svelte.dev)
- error handling 
- data - [JSON API](https://dri.es/headless-cms-rest-vs-jsonapi-vs-graphql)

**2. Service Layer (aka Business Logic Layer)**

- [Workflow](https://news.ycombinator.com/item?id=19732447) (state management) and Task orchestration (can be [long running](https://blog.bernd-ruecker.com/what-are-long-running-processes-b3ee769f0a27)) aka control plane - [Uber Cadence](https://cadenceworkflow.io) - tasks included here are only those necessary to complete the user transaction otherwise delegate to the data plane
- Data plane (Using [Dapr](https://dapr.io), pull data from RPC or push data to asynchronous task queue like NATS, Kafka, etc )

**3. Data Access Layer** (via [Vault](https://www.vaultproject.io/docs/secrets/databases/index.html)-backed DAL proxy)

- read data from SQL/NoSQL stores
- write data to SQL/NoSQL stores 
