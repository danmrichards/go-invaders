package sound

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gobuffalo/packr/v2"
)

const sr = beep.SampleRate(11025)

var (
	box  = packr.New("sound", "./data")
	wavs = []string{
		"0.wav",
		"1.wav",
		"2.wav",
		"3.wav",
		"4.wav",
		"5.wav",
		"6.wav",
		"7.wav",
		"8.wav",
	}
)

func init() {
	speaker.Init(sr, sr.N(time.Second/10)) //nolint:errcheck
}

type sound struct {
	fmt   beep.Format
	strmr beep.StreamCloser
}

// Player is responsible for playing buffered sounds.
type Player struct {
	bufs []*beep.Buffer
}

// NewPlayer returns an instantiated sound player with all sounds buffered.
func NewPlayer() (*Player, error) {
	snds, err := prepareSounds()
	if err != nil {
		return nil, err
	}

	p := &Player{
		bufs: make([]*beep.Buffer, 0, len(wavs)),
	}

	for _, s := range snds {
		b := beep.NewBuffer(s.fmt)
		b.Append(s.strmr)
		s.strmr.Close()

		p.bufs = append(p.bufs, b)
	}

	return p, nil
}

// Play plays the given sound.
func (p *Player) Play(name string) {
	var i int
	for wi, wn := range wavs {
		if wn == name {
			i = wi
		}
	}

	speaker.Play(p.bufs[i].Streamer(i, p.bufs[i].Len()))
}

// prepareSounds returns a list of decoded wave files.
func prepareSounds() ([]sound, error) {
	snds := make([]sound, 0, len(wavs))

	for _, w := range wavs {
		b, err := box.Find(w)
		if err != nil {
			return nil, err
		}

		ws, wf, err := wav.Decode(ioutil.NopCloser(bytes.NewReader(b)))
		if err != nil {
			return nil, err
		}

		snds = append(snds, sound{
			fmt:   wf,
			strmr: ws,
		})
	}

	return snds, nil
}
