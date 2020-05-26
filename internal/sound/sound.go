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

var box = packr.New("sound", "./data")

// Play plays the sound with the given name.
func Play(name string) error {
	b, err := box.Find(name)
	if err != nil {
		return err
	}

	s, format, err := wav.Decode(ioutil.NopCloser(bytes.NewReader(b)))
	if err != nil {
		return err
	}

	if err = speaker.Init(
		format.SampleRate,
		format.SampleRate.N(time.Second/10),
	); err != nil {
		return err
	}

	done := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	<-done

	return nil
}
