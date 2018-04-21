package basicnode

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/signal"
)

// Execute processes the business logic of the node. This is to update the
// energy and threshold as well as executing the node's action, if any.
func (o *Object) Execute(ctx context.Context, sig signal.Interface) (signal.Interface, error) {
	var err error

	{
		err = o.updateEnergy(ctx)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		err = o.updateThreshold(ctx)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	if o.action != nil {
		sig, err = o.action.Execute(ctx, sig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return sig, nil
}
