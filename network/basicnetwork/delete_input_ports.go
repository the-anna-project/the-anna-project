package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/port"
)

func (o *Object) DeleteInputPorts(ctx context.Context, node node.Interface, ports []port.Interface) error {
	for _, p := range ports {
		err := o.storage.Port.Input.RemoveFromSet(node.ID(), p.NodeID())
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
