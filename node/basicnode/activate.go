package basicnode

import (
	"context"

	"github.com/the-anna-project/the-anna-project/signal"
	"github.com/the-anna-project/the-anna-project/signal/stopsignal"
)

// Activate checks if the node should be activated or not. The returned signal
// indicates the decision of the node. In case the node should not be activated
// a stop signal is returned, otherwhise the received signal is dispatched.
func (o *Object) Activate(ctx context.Context, sig signal.Interface) (signal.Interface, error) {
	// TODO consider random resistance for activation calculation
	// TODO respect signal injection
	if o.Energy() < o.Threshold() {
		return &stopsignal.Object{}, nil
	}

	return sig, nil
}
