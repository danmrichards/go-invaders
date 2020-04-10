package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/danmrichards/go-invaders/internal/machine"
)

var (
	dir   string
	debug bool
)

func main() {
	flag.StringVar(&dir, "dir", "roms", "Path to directory containing ROM files")
	flag.BoolVar(&debug, "debug", false, "Run the emulator in debug mode")
	flag.Parse()

	var opts []machine.Option
	if debug {
		opts = append(opts, machine.WithDebugEnabled())
	}

	m := machine.New(opts...)

	if err := m.LoadROM(dir); err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	// Basic shutdown handler.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	for {
		select {
		case <-done:
			break
		default:
		}

		// Emulate a cycle.
		if err := m.Cycle(); err != nil {
			log.Fatal(err)
		}
	}
}
