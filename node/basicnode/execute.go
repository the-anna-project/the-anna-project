package basicnode

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/the-anna-project/the-anna-project/nodeidcontext"
	"github.com/the-anna-project/the-anna-project/signal"
	"github.com/the-anna-project/the-anna-project/signal/stopsignal"
)

// Execute checks if the node should be activated or not. In case the node is
// being activated it will increase its energy level. When this is done the
// node's action will be executed. The returned signals indicate eventual action
// results and decisions of the node. In case the node should not be activated,
// the returned list only contains a stop signal, otherwhise the node returns
// signals supposed to be dispatched.
func (o *Object) Execute(ctx context.Context, sigs []signal.Interface) ([]signal.Interface, error) {
	var err error

	if o.Energy() < o.Threshold() {
		return []signal.Interface{&stopsignal.Object{}}, nil
	}

	ctx = nodeidcontext.NewContext(ctx, o.ID())

	{
		err = o.increaseEnergy(ctx)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	sigs, err = o.Action().Execute(ctx, sigs)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return sigs, nil
}
