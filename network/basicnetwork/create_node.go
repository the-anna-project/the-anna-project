package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/the-anna-project/the-anna-project/node"
)

func (o *Object) CreateNode(ctx context.Context, n node.Interface) error {
	err := o.storage.Node.Create(n.ID(), n.Action().ID())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
