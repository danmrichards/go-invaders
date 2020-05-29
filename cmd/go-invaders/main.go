package main

import (
	"flag"
	"log"

	"github.com/danmrichards/go-invaders/internal/machine"
	"github.com/danmrichards/go-invaders/internal/memory"
	"github.com/faiface/pixel/pixelgl"
)

var (
	dir         string
	debug       bool
	scaleFactor int
)

func main() {
	flag.StringVar(&dir, "dir", "roms", "Path to directory containing ROM files")
	flag.BoolVar(&debug, "debug", false, "Run the emulator in debug mode")
	flag.IntVar(&scaleFactor, "scale-factor", 2, "Scales the original video resolution (224x256)")
	flag.Parse()

	// Instantiate 64K of memory.
	mem := make(memory.Basic, 65536)
	if err := mem.LoadROM(dir); err != nil {
		log.Fatal(err)
	}

	opts := []machine.Option{
		machine.WithScaleFactor(scaleFactor),
	}
	if debug {
		opts = append(opts, machine.WithDebugEnabled())
	}

	// Instantiate the Space Invaders machine.
	m, err := machine.New(mem, opts...)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Log the controls.

	pixelgl.Run(m.Run)
}
