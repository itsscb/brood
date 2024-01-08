package brood

import (
	"testing"

	"github.com/itsscb/brood/spore"
)

func TestNew(t *testing.T) {
	in := spore.Channel(0)
	out := spore.Channel(1)
	got := New(in, out)
	want := Brood{
		inbound:  in,
		outbound: out,
		close:    make(chan struct{}, 1),
		// TODO: Implement and Test brood.url & brood.ListenAndServe()
		url: "",
	}

	if cap(got.inbound) != cap(want.inbound) {
		t.Fatalf("brood.inbound cap: got=%d, want=%d", cap(got.inbound), cap(want.inbound))
	}

	if cap(got.outbound) != cap(want.outbound) {
		t.Fatalf("brood.outbound cap: got=%d, want=%d", cap(got.outbound), cap(want.outbound))
	}

	if cap(got.close) != cap(want.close) {
		t.Fatalf("brood.close cap: got=%d, want=%d", cap(got.close), cap(want.close))
	}
}
