# Go Invaders
A [Space Invaders][1] emulator implemented in Go

## Installation
```bash
$ go get -u github.com/danmrichards/go-invaders/cmd/go-invaders/...
```
## Usage
In order to play Space Invaders you will need to supply the ROM files. For
obvious reasons they are not included in this repo.
```
  -debug
        Run the emulator in debug mode
  -dir string
        Path to directory containing ROM files (default "roms")
  -scale-factor int
        Scales the original video resolution (224x256) (default 2)
```

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
$ make build
```

#### Cross compile for Windows
If you are running on Linux, it is possible to cross-compile the application for Windows.

In order to do this you will need a GCC compiler that targets Windows, such as mingw. Once installed, you can cross compile like so:
```
$ GOOS=windows GOARCH=${GOARCH} CGO_ENABLED=1 CC=${GCC} go build -ldflags="-s -w" -o bin/go-invaders-windows-${GOARCH}.exe ./cmd/go-invaders
```
> Replace `${GOARCH}` with your target architecture (e.g. amd64) and replace ${GCC} with the name of your Windows GCC (e.g. x86_64-w64-mingw32-gcc)

[1]: https://en.wikipedia.org/wiki/Space_Invaders
[2]: https://github.com/faiface/pixel#requirements
[3]: https://github.com/goware/modvendor
[4]: https://github.com/gobuffalo/packr/tree/master/v2
