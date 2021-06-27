FROM golang:1.16 as builder

WORKDIR /app

# better to do this first before copying code base as code can change
# so all layers below will be rebuild which is not necessary
COPY go.mod .
RUN go mod download

# copy codebase into workdir
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server *.go

FROM alpine:3.14.0 as runner

ARG APP_PORT

EXPOSE $APP_PORT

COPY --from=builder /app/server /server

CMD ["sh","-c", "./server"]
