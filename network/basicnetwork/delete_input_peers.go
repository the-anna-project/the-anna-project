package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
)

func (o *Object) DeleteInputPeers(ctx context.Context, node node.Interface, peers []peer.Interface) error {
	for _, p := range peers {
		err := o.storage.Peer.Input.Delete(node.ID(), p.NodeID())
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
