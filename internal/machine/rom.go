package machine

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// romOffsets represents the memory offsets that each ROM part begins it's data
// range at.
var romOffsets = map[string]uint32{
	"invaders.h": 0,
	"invaders.g": 0x800,
	"invaders.f": 0x1000,
	"invaders.e": 0x1800,
}

// LoadROM loads the Space Invaders ROM into memory.
//
// The ROM itself is broken down into 4 parts with the following address ranges:
//
// $0000-$07ff: invaders.h
// $0800-$0fff: invaders.g
// $1000-$17ff: invaders.f
// $1800-$1fff: invaders.e
func (m *Machine) LoadROM(dir string) error {
	if dir == "" {
		return errors.New("ROM directory cannot be empty")
	}
	if _, err := os.Stat(dir); err != nil {
		return err
	}

	if err := m.loadROMPart(dir, "invaders.h"); err != nil {
		return err
	}
	if err := m.loadROMPart(dir, "invaders.g"); err != nil {
		return err
	}
	if err := m.loadROMPart(dir, "invaders.f"); err != nil {
		return err
	}
	if err := m.loadROMPart(dir, "invaders.e"); err != nil {
		return err
	}

	return nil
}

func (m *Machine) loadROMPart(dir, part string) error {
	path := filepath.Join(dir, part)

	// Get the correct memory offset for this ROM part.
	offset, ok := romOffsets[part]
	if !ok {
		return fmt.Errorf(
			"could not determine ROM offset: unknown part: %q", part,
		)
	}

	rf, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open ROM part (%q): %w", path, err)
	}
	defer rf.Close()

	if _, err = rf.Read(m.mem.Dump()[offset:]); err != nil {
		return fmt.Errorf("could not read ROM part (%q): %w", path, err)
	}

	return nil
}
