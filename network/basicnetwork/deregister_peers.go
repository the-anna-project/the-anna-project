package basicnetwork

import (
	"context"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
)

// TODO implement node and action deregistration
func (o *Object) DeregisterPeers(ctx context.Context, n node.Interface, p []peer.Interface) error {
	return nil
}
