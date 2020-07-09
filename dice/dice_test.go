package dice_test

import (
	"reflect"
	"testing"

	"github.com/keita7i/reactive-dice/dice"
)

func TestRef(t *testing.T) {
	cs := []struct {
		faces []string
		want  []string
	}{
		{faces: []string{"a", "b", "c", "d", "e"}, want: []string{"a", "b", "c", "d", "e"}},
		{faces: []string{"a", "b", "c", "d", "e", "f"}, want: []string{"a", "b", "c", "d", "e", "f"}},
		{faces: []string{"a", "b", "c", "d", "e", "f", "g"}, want: []string{"a", "b", "c", "d", "e", "f"}},
	}
	for _, c := range cs {
		got := dice.New("name", c.faces).Ref()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("len(dice.New(%v).Ref()) = %v; want %v", c.faces, got, c.want)
		}
	}
}
