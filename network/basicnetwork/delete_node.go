package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/the-anna-project/the-anna-project/node"
)

func (o *Object) DeleteNode(ctx context.Context, node node.Interface) error {
	err := o.storage.Node.Delete(node.ID())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
