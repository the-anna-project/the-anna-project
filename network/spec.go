package network

import (
	"context"

	"github.com/the-anna-project/the-anna-project/action"
	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
)

// TODO network creates new nodes
type Interface interface {
	Deregister(ctx context.Context, n node.Interface, a action.Interface) error
	// RandomPeers returns a random, arbitrary list of peer information
	//
	//     TODO 0 to 10 nodes are picked as peers
	//
	RandomPeers(ctx context.Context) ([]peer.Interface, error)
	Register(ctx context.Context, n node.Interface, a action.Interface) error
}
