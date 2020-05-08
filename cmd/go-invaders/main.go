package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/danmrichards/go-invaders/internal/cpu"
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

	// Instantiate the Intel 8080.
	var opts []cpu.Option
	if debug {
		opts = append(opts, cpu.WithDebugEnabled())
	}
	i80 := cpu.NewIntel8080(mem, opts...)

	done := make(chan struct{})

	// Instantiate the Space Invaders machine.
	m := machine.New(i80, mem, done)

	// Basic shutdown handler.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		close(done)
	}()

	// Run the machine until it errors or we shutdown.
	if err := m.Run(); err != nil {
		log.Fatal(err)
	}
}
