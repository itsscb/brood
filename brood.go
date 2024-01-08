package brood

import (
	"errors"

	"github.com/itsscb/brood/spore"
)

var (
	ErrNotImplemented = errors.New("brood: not implemented yet")
)

type Brood struct {
	inbound  <-chan spore.Spore
	outbound chan spore.Spore
	close    chan struct{}
	url      string
}

func New(in, out chan spore.Spore) *Brood {
	close := make(chan struct{}, 1)
	return &Brood{
		inbound:  in,
		outbound: out,
		close:    close,
	}
}

func (b *Brood) Publish(s spore.Spore) error {
	if b.url != "" {
		return ErrNotImplemented
	}

	b.outbound <- s

	return nil
}

func (b *Brood) Subscribe(hives string) (chan spore.Spore, error) {
	if b.url != "" {
		return b.outbound, ErrNotImplemented
	}

	return b.outbound, nil
}

func (b *Brood) ListenAndServe(url string) error {
	b.url = url

	// TODO: Implement Logic
	return ErrNotImplemented
}
