build:
	go mod download && go build -o server *.go

test:
	go test ./... --cover

run:
	SENTRY_DSN=$(SENTRY_DSN) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) \
	DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) \
	./server
