package network

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/peer"
)

// TODO
func (o *Object) RandomPeers(ctx context.Context) ([]peer.Interface, error) {
	var peers []peer.Interface

	for i := 0; i < peers; i++ {
		o.peers, err = o.peer.NewFloat64(ctx, o.energy, 1)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return peers, nil
}
