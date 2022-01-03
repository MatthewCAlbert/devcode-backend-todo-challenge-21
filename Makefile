.PHONY: run
run: ## run the API server
	go run ${LDFLAGS} cmd/server/main.go

.PHONY: run-dev
run-dev: ## run the API server dev
	docker-compose -f ./docker-compose.yml -f ./docker-compose.dev.yml up

.PHONY: stop-dev
stop-dev: ## stop the API server dev
	docker-compose -f ./docker-compose.yml -f ./docker-compose.dev.yml down

.PHONY: run-prod
run-prod: ## run docker-compose with production mode
	docker-compose -f ./docker-compose.yml -f ./docker-compose.prod.yml up -d

.PHONY: stop-prod
stop-prod: ## stop docker-compose with production mode
	docker-compose -f ./docker-compose.yml -f ./docker-compose.prod.yml down

.PHONY: start-db
start-db: ## start db
	docker-compose -f ./docker-compose.yml up -d db

.PHONY: reset-db
reset-db: ## reset db
	docker exec -d devcode-backend-todo-challenge-mysql /reset-db.sh

.PHONY: stop-db
stop-db: ## stop db
	docker-compose -f ./docker-compose.yml down db

.PHONY: run-restart
run-restart: ## restart the API server
	@pkill -P `cat $(PID_FILE)` || true
	@printf '%*s\n' "80" '' | tr ' ' -
	@echo "Source file changed. Restarting server..."
	@go run ${LDFLAGS} cmd/server/main.go & echo $$! > $(PID_FILE)
	@printf '%*s\n' "80" '' | tr ' ' -

.PHONY: build
build:  ## build the API server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o server ./cmd/server

.PHONY: build-docker
build-docker: ## build the API server as a docker image
	docker build -f cmd/server/Dockerfile.prod -t devcode-backend-todo-challenge .

.PHONY: push-docker
push-docker: ## push built docker image
	docker tag devcode-backend-todo-challenge:latest matthewcalbert/devcode-backend-todo-challenge:latest
	docker push matthewcalbert/devcode-backend-todo-challenge:latest

.PHONY: clean
clean: ## remove temporary files
	rm -rf server coverage.out coverage-all.out

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out