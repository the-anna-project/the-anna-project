package node

import (
	"context"

	"github.com/the-anna-project/the-anna-project/signal"
	"github.com/the-anna-project/the-anna-project/signal/stopsignal"
)

// Activate checks if the node should be activated or not. The returned signal
// indicates the decision of the node. In case the node should not be activated
// a stop signal is returned, otherwhise the received signal is dispatched.
func (o *Object) Activate(ctx context.Context, sig signal.Interface) (signal.Interface, error) {
	if o.Energy() < o.Theshold() {
		return stopsignal.Object{}, nil
	}

	return sig, nil
}
