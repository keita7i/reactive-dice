package dice_test

import (
	"testing"

	"github.com/keitam913/reactive-dice/dice"
)

func TestRollHasValidLen(t *testing.T) {
	cs := []struct {
		faces   []string
		wantLen int
	}{
		{faces: []string{"a", "b", "c", "d", "e"}, wantLen: 5},
		{faces: []string{"a", "b", "c", "d", "e", "f"}, wantLen: 6},
		{faces: []string{"a", "b", "c", "d", "e", "f", "g"}, wantLen: 6},
	}
	for _, c := range cs {
		gotLen := len(dice.New(c.faces).Roll())
		if gotLen != c.wantLen {
			t.Errorf("len(dice.New(%v).Roll()) = %v; want %v", c.faces, gotLen, c.wantLen)
		}
	}
}
