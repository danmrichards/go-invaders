GOARCH=amd64
BINARY=go-invaders

build:
	GO111MODULE=on GOARCH=${GOARCH} GOOS=linux go build -ldflags="-s -w" -o bin/${BINARY}-linux-${GOARCH} ./cmd/go-invaders/main.go

.PHONY: build