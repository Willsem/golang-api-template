APP_NAME := golang-api-template

ifeq (${DOCKER_LOGIN},)
IMAGE := $(APP_NAME)
else
IMAGE := ${DOCKER_LOGIN}/$(APP_NAME)
endif

TAG := $$(\
if [[ $$(git describe --tags --abbrev=0) ]]; then \
    echo $$(git describe --tags --abbrev=0); \
else \
    echo "v0.0.0"; \
fi \
)

COMMIT := $$(git rev-parse HEAD)
DATE := $$(date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS := "\
	-X github.com/Willsem/golang-api-template/internal/app/build.Version=$(TAG) \
	-X github.com/Willsem/golang-api-template/internal/app/build.VersionCommit=$(COMMIT) \
	-X github.com/Willsem/golang-api-template/internal/app/build.BuildDate=$(DATE) \
	"

all: lint test build

run:
	@dotenv -e ./configs/.env.example -- go run ./cmd/api/main.go

lint:
	golangci-lint run -v

build:
	@go build -ldflags=$(LDFLAGS) -mod readonly -o ./bin/golang-api-template ./cmd/api/main.go

test:
	@go test -race -cover ./... -v

swag:
	@swag fmt && \
		swag init --parseDependency --parseDepth 5 -g ./cmd/api/main.go -o ./api

generate:
	@go generate ./...

docker-build:
	@docker build \
		--build-arg LDFLAGS=$(LDFLAGS) \
		-t $(IMAGE):$(TAG)-$(COMMIT) \
		.

docker-push:
	@docker push $(IMAGE):$(TAG)-$(COMMIT)

release-build:
	@docker build \
		--build-arg LDFLAGS=$(LDFLAGS) \
		-t $(IMAGE):$(TAG) \
		.

release-push:
	@docker push $(IMAGE):$(TAG)
