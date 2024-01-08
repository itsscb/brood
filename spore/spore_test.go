package spore_test

import (
	"testing"

	"github.com/itsscb/brood/spore"
)

func TestChannel(t *testing.T) {
	for i := -10; i < 100; i++ {
		want := i
		ch := spore.Channel(i)

		if i < 0 {
			want = 0
		}
		if cap(ch) != want {
			t.Errorf("spore.Channel: want=%d, got=%d", want, cap(ch))
		}
	}
}
