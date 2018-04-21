package node

import (
	"context"

	"github.com/the-anna-project/the-anna-project/spec/signal"
)

func (o *Object) Signals(ctx context.Context, sig signal.Interface) ([]signal.Interface, error) {
	var sigs []signal.Interface

	for _, p := range o.Peers() {
		s := sig.Copy()

		s.SetPeer(p)

		sigs = append(sigs, s)
	}

	return sigs, nil
}
