package network

import (
	"context"

	"github.com/the-anna-project/the-anna-project/peer"
)

// TODO
//
//     network creates new nodes
//
type Interface interface {
	// RandomPeers returns a random, arbitrary list of peer information
	//
	//     TODO 0 to 10 nodes are picked as peers
	//
	RandomPeers(ctx context.Context) ([]peer.Interface, error)
}
