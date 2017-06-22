https://github.com/ibmendoza/go-examples/tree/master/docker

**Docker Swarm Mode Workflow with Golang**

- Write code
- Test
- Build (result: executable)
- Build executable into a container image
- Push container image to container registry
- From any Docker Swarm Mode manager node, pull container image from registry
- Run the container image using [docker service](https://blog.nimbleci.com/2016/10/05/how-to-automate-the-deploy-of-docker-services/) from manager node
- Done
