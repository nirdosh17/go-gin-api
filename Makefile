SERVICE_NAME:=go-rest-api
IMAGE_VERSION:=$(shell cat .image_version)
APP_PORT:=8080

DOCKER_REPO_ADDR:=nirdoshgautam/$(SERVICE_NAME)

build:
	docker build -t $(SERVICE_NAME) .

run:
	docker run -it -p $(APP_PORT):$(APP_PORT) \
		-e SENTRY_DSN=$(SENTRY_DSN) \
		-e DB_HOST=host.docker.internal \
		-e DB_PORT=$(DB_PORT) \
		-e DB_USER=$(DB_USER) \
		-e DB_PASSWORD=$(DB_PASSWORD) \
		-e DB_NAME=$(DB_NAME) \
		$(SERVICE_NAME):latest

build-binary:
	go mod download && go build -o server *.go

run-local: build-binary
	SENTRY_DSN=$(SENTRY_DSN) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) \
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) \
	./server

push-image:
	docker tag $(SERVICE_NAME):latest $(DOCKER_REPO_ADDR):$(IMAGE_VERSION)
	docker push $(DOCKER_REPO_ADDR):$(IMAGE_VERSION)

test:
	go test ./... --cover
