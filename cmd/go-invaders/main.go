package main

import (
	"flag"
	"log"

	"github.com/faiface/pixel/pixelgl"

	"github.com/danmrichards/go-invaders/internal/machine"
	"github.com/danmrichards/go-invaders/internal/memory"
)

var (
	dir   string
	debug bool
)

func main() {
	flag.StringVar(&dir, "dir", "roms", "Path to directory containing ROM files")
	flag.BoolVar(&debug, "debug", false, "Run the emulator in debug mode")
	flag.Parse()

	// Instantiate 64K of memory.
	mem := make(memory.Basic, 65536)
	if err := mem.LoadROM(dir); err != nil {
		log.Fatal(err)
	}

	var opts []machine.Option
	if debug {
		opts = append(opts, machine.WithDebugEnabled())
	}

	// Instantiate the Space Invaders machine.
	m := machine.New(mem, opts...)

	pixelgl.Run(m.Run)
}
