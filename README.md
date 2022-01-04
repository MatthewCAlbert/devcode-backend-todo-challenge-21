# Devode Backend Challenge: TODO List API

## Info

- This is stripped down version ✂️
- Docker Hub 🐳 [matthewcalbert/devcode-backend-todo-challenge](https://hub.docker.com/r/matthewcalbert/devcode-backend-todo-challenge)
- Find proper version (non-stripped and non-custom cached) at ["vanilla" 🍦 branch](https://github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/tree/vanilla).

## Built using

- Go 1.16 🐹
- Fiber v2
- GORM
- MySQL

## Notes

- Dev Server now is working 🤟

## Development Local

```sh
# In a terminal (to dev with fresh db)
./scripts/run-dev-reset.sh

# Open new terminal (to start testing, don't forget to rerun script above)
make test
```

## Reference

- https://github.com/qiangxue/go-rest-api (Structure Sample)
- https://github.com/golang-standards/project-layout
- https://firehydrant.io/blog/develop-a-go-app-with-docker-compose/ (Dev Mode)
- https://gist.github.com/derlin/0be53d0d7f38db181198aada024269b8