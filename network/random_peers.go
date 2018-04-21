package network

import (
	"context"

	"github.com/the-anna-project/the-anna-project/spec/peer"
)

// TODO lookup random peers from storage
func (o *Object) RandomPeers(ctx context.Context) ([]peer.Interface, error) {
	var peers []peer.Interface

	return peers, nil
}
