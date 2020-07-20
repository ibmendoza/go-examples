**Stack**

Frontend
- Svelte 
- [Storybook](https://www.learnstorybook.com/)
- Svite, Snowpack
- [Zero](https://github.com/remoteinterview/zero)
- https://github.com/tjinauyeung/svelte-forms-lib

Backend
- [Oak](https://github.com/oakserver/oak)
- [Validator](https://github.com/icebob/fastest-validator)
- Go

Links

- https://stackoverflow.com/questions/54652683/how-to-access-secrets-in-node-js-with-hashicorp-vault
- https://medium.com/better-programming/the-battle-of-bundlers-6333a4e3eda9
- https://husio.github.io/blog/accessing-data-in-go/
- https://husobee.github.io/school/secure-programming/2017/08/28/secure-programming-1.html

**Separation of Concerns**

App logic should be [middleware](https://github.com/oakserver/oak) only, no more no less<br>
TL;DR Leave app logic alone

- Separate configuration from app logic
- [Separate deployment logic from app logic](https://github.com/crossplane/crossplane)
- [Separate authentication/authorization from app logic](https://github.com/dapr/samples/tree/master/7.middleware)
- [Separate security logic from app logic](https://github.com/dapr/samples/tree/master/9.secretstore)
- [Separate async logic from transactions](https://github.com/dapr/samples/tree/master/4.pub-sub)
- [Separate UI logic from app logic](https://svelte.dev/)
- [Separate database logic from app logic](https://husio.github.io/blog/accessing-data-in-go/)
- [Separate ops logic from app logic](https://github.com/dapr/samples/tree/master/8.observability)
- [Separate endpoint logic from app logic](https://zeroserver.io)
- [Separate internal RPC from app logic](https://github.com/dapr/docs/blob/master/concepts/service-invocation/README.md)
- [Separate external RPC from app logic](https://github.com/dapr/docs/blob/master/concepts/bindings/README.md)
- Separate end user error logic from app logic 
- Separate system error logic from app logic
- [Separate state management logic from app logic](https://github.com/dapr/docs/blob/master/concepts/state-management/README.md)
