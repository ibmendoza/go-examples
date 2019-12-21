**Front-end**

- SvelteJS

**Backend**

pic courtesy of [LogRocket](https://blog.logrocket.com/the-perfect-architecture-flow-for-your-next-node-js-project/)

<img src="https://itjumpstart.files.wordpress.com/2019/12/business-logic-api-routes.png">

1. Controller (e.g. SapperJS) - API routes
2. Service Layer (aka Business Logic Layer)

- Workflow (state management) and Task orchestration (can be [long running](https://blog.bernd-ruecker.com/what-are-long-running-processes-b3ee769f0a27)) aka control plane - [Uber Cadence](https://cadenceworkflow.io) 
- Data plane (pull data from RPC (Dapr) or push data to asynchronous task queue like NATS)

3. Data Access Layer
- read data from SQL/NoSQL stores
- write data to SQL/NoSQL stores
