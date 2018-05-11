package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
	"github.com/the-anna-project/the-anna-project/peer/basicpeer"
)

func (o *Object) SearchOutputPeers(ctx context.Context, node node.Interface) ([]peer.Interface, error) {
	result, err := o.storage.Peer.Output.SearchSet(node.ID())
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var peers []peer.Interface

	for _, r := range result {
		c := basicpeer.Config{
			NodeID: r,
		}

		p, err := basicpeer.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		peers = append(peers, p)
	}

	return peers, nil
}
