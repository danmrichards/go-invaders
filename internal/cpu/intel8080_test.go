package cpu

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/danmrichards/go-invaders/internal/memory"
)

var debug = flag.Bool("debug", false, "Run the emulator in debug mode")

func TestCPU(t *testing.T) {
	// Instantiate 64K of memory.
	mem := make(memory.Basic, 65536)

	// Load the test ROM.
	rf, err := os.Open(filepath.Join("testdata", "8080PRE.COM"))
	if err != nil {
		t.Fatal(err)
	}
	defer rf.Close()

	// The test ROM assumes the program code starts at 0x100. So read the ROM
	// into memory with this as an offset.
	if _, err = rf.Read(mem[0x100:]); err != nil {
		t.Fatal(err)
	}

	// Manually set the first instruction in the memory to be a JMP to 0x100.
	// This will force the emulation to start at the point where the ROM expects.
	mem.Write(0, 0xc3)
	mem.Write(1, 0)
	mem.Write(2, 0x01)

	var opts []Option
	if *debug {
		opts = append(opts, WithDebugEnabled())
	}
	i80 := NewIntel8080(mem, opts...)

	for {
		if i80.halted {
			return
		}

		if err := i80.Step(); err != nil {
			t.Fatal(err)
		}

		// TODO: Why does the printing work like this?
		if i80.pc == 0x05 {
			if i80.c == 0x09 {
				addr := uint16(i80.d)<<8 | uint16(i80.e)

				for {
					c := mem.Read(addr)

					if fmt.Sprintf("%c", c) == "$" {
						break
					} else {
						addr++
					}

					fmt.Printf("%c", c)
				}
			}
			if i80.c == 0x02 {
				//fmt.Printf("%c", i80.e)
			}
		}

		// TODO: Error detection?
		if i80.pc == 0x00 {
			fmt.Println()
			fmt.Println()
			break
		}

		//time.Sleep(50 * time.Millisecond)
	}
}
