package dice

import (
	"math/rand"
	"sync/atomic"
)

type Dice struct {
	faces *atomic.Value
}

func New(faces []string) Dice {
	fs := &atomic.Value{}
	fs.Store(faces)
	return Dice{
		faces: fs,
	}
}

func (d Dice) Ref() []string {
	fs := append([]string{}, d.faces.Load().([]string)...)
	end := 6
	if end > len(fs) {
		end = len(fs)
	}
	return fs[:end]
}

func (d Dice) Roll() {
	fs := append([]string{}, d.faces.Load().([]string)...)
	rand.Shuffle(len(fs), func(i, j int) {
		fs[i], fs[j] = fs[j], fs[i]
	})
	d.faces.Store(fs)
}
