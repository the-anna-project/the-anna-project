package basicnode

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/the-anna-project/the-anna-project/signal"
	"github.com/the-anna-project/the-anna-project/signal/stopsignal"
	"github.com/the-anna-project/the-anna-project/signal/waitsignal"
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

	{
		peers, err := o.network.SearchInputPeers(ctx, o)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		for _, p := range peers {
			o.mutex.Lock()
			_, ok := o.bufferedSignals[p.NodeID()]
			if !ok {
				o.bufferedSignals[p.NodeID()] = sig
			}
			o.mutex.Unlock()
		}

		if len(peers) != len(o.bufferedSignals) {
			return &waitsignal.Object{}, nil
		}

		for _, p := range peers {
			o.mutex.Lock()
			s := o.bufferedSignals[p.NodeID()]
			o.mutex.Unlock()
			sig.AppendData(s.Data())
		}

		o.mutex.Lock()
		o.bufferedSignals = map[string]signal.Interface{}
		o.mutex.Unlock()
	}

	return sig, nil
}
