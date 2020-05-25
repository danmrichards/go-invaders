GOARCH=amd64
BINARY=go-invaders

build:
	GO111MODULE=on GOARCH=${GOARCH} GOOS=linux go build -ldflags="-s -w" -o bin/${BINARY}-linux-${GOARCH} ./cmd/go-invaders/main.go

lint:
	golangci-lint run ./cmd/... ./internal/...

deps:
	go mod verify && \
	go mod tidy && \
	go mod vendor && \
	modvendor -copy="**/*.c **/*.h **/*.m"

.PHONY: build