package main

import (
	"flag"
	"fmt"
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

	// TODO: Implement configuration for colours.

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

	fmt.Println("*****************************")
	fmt.Println("* Welcome to Space Invaders *")
	fmt.Println("*                           *")
	fmt.Println("* Insert coin = C           *")
	fmt.Println("* 1P start = 1              *")
	fmt.Println("* 2P start = 2              *")
	fmt.Println("* 1P shoot = W              *")
	fmt.Println("* 1P left = Q               *")
	fmt.Println("* 1P right = E              *")
	fmt.Println("* 2P shoot = O              *")
	fmt.Println("* 2P left = I               *")
	fmt.Println("* 2P right = P              *")
	fmt.Println("* Tilt = T                  *")
	fmt.Println("*                           *")
	fmt.Println("*****************************")

	pixelgl.Run(m.Run)
}
