## Go build tags
- following [YouTube](https://www.youtube.com/watch?v=eqvDSkuBihs) tutorial
- using Makefile [YouTube](https://www.youtube.com/watch?v=XlobWOgcK7Y) tutorial


### Usage:
> Prerequisites:
- Go v1.22
```bash
# run Go in Docker container
make dev
```

> Run:
```bash
make run
```

### Requests:
> Root
```bash
curl http://localhost:8080/
```

> Creating users:
```bash
curl -i -X POST -d '{"name":"John"}' http://localhost:8080/users
```

> Get user by ID:
```bash
curl -i http://localhost:8080/users/1
```
