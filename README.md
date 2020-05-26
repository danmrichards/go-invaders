# Go Invaders
A [Space Invaders][1] emulator implemented in Go

## Building From Source
### Pre-requisites
The emulator uses the following packages which have requirements of their own
before we can build with them. Follow the instructions for each:

* [Pixel][2]
* [ModVendor][3]
* [Packr][4]

### Building
Clone this repo and build the binary:
```bash
$ make
```

#### Windows
To build from source on Windows run this command: 
```bash
go build -ldflags="-s -w" -o bin/go-invaders-windows-amd64.exe ./cmd/go-invaders
```
> Swap out `amd64` for the relevant architecture. See `go env GOARCH`

## Installation
```bash
$ go get -u github.com/danmrichards/go-invaders/cmd/go-invaders/...
```

[1]: https://en.wikipedia.org/wiki/Space_Invaders
[2]: https://github.com/faiface/pixel#requirements
[3]: https://github.com/goware/modvendor
[4]: https://github.com/gobuffalo/packr/tree/master/v2
