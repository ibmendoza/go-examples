**Front-end**

- SvelteJS

**Backend**

pic courtesy of [LogRocket]()

https://itjumpstart.files.wordpress.com/2019/12/business-logic-api-routes.png

1. Controller (e.g. SapperJS)
2. Service Layer

- Workflow (state management) and Task orchestration (can be [long running](https://blog.bernd-ruecker.com/what-are-long-running-processes-b3ee769f0a27)) aka control plane - [Uber Cadence](https://cadenceworkflow.io)
- Data plane (request/reply - Dapr or async - task queue like NATS)
