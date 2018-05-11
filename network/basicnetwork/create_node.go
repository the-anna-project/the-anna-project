package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/node"
)

func (o *Object) CreateNode(ctx context.Context, node node.Interface) error {
	err := o.storage.Node.Create(node.ID(), node.Action().ID())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
