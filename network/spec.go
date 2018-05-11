package network

import (
	"context"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
)

type Interface interface {
	// CreateInputPeers creates input peers for the given node based on random
	// guesses. It also creates output peers implicitly on the other side. This is
	// because an input peer for one node is an output peer for the other.
	CreateInputPeers(ctx context.Context, node node.Interface) error
	// CreateNode adds shared knowledge about the given node and its associated
	// action.
	CreateNode(ctx context.Context, node node.Interface) error
	// DeleteInputPeers deletes the given references of input peers from the given
	// node.
	DeleteInputPeers(ctx context.Context, node node.Interface, peers []peer.Interface) error
	// DeleteNode removes shared knowledge about the given node and the given
	// action from the configured storage.
	DeleteNode(ctx context.Context, node node.Interface) error
	// DeleteOutputPeers deletes the given references of output peers from the
	// given node.
	DeleteOutputPeers(ctx context.Context, node node.Interface, peers []peer.Interface) error
	// SearchInputPeers returns the shared knowledge about the peers providing
	// signals as input for the given node.
	SearchInputPeers(ctx context.Context, node node.Interface) ([]peer.Interface, error)
	// SearchOutputPeers returns the shared knowledge about the peers the given
	// node provides signals as output for.
	SearchOutputPeers(ctx context.Context, node node.Interface) ([]peer.Interface, error)
}
