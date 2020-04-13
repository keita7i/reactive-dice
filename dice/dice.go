package dice

import "math/rand"

type Dice struct {
	faces []string
}

func New(faces []string) Dice {
	return Dice{
		faces: faces,
	}
}

func (d Dice) Roll() []string {
	var faces []string
	faces = append(faces, d.faces...)
	rand.Shuffle(len(faces), func(i, j int) {
		faces[i], faces[j] = faces[j], faces[i]
	})
	end := 6
	if end > len(faces) {
		end = len(faces)
	}
	return faces[:end]
}
