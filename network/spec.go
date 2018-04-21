package network

import (
	"context"

	"github.com/the-anna-project/the-anna-project/action"
	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
)

type Interface interface {
	// CreateNode adds shared knowledge about the given node and the given
	// action to the configured storage.
	CreateNode(ctx context.Context, n node.Interface, a action.Interface) error
	// DeleteInputPeers TODO
	DeleteInputPeers(ctx context.Context, n node.Interface, p []peer.Interface) error
	// DeleteOutputPeers TODO
	DeleteOutputPeers(ctx context.Context, n node.Interface, p []peer.Interface) error
	// DeleteNode removes shared knowledge about the given node and the given
	// action from the configured storage.
	DeleteNode(ctx context.Context, n node.Interface, a action.Interface) error
	// SearchInputPeers returns the shared knowledge about the peers providing
	// signals as input for the given node.
	SearchInputPeers(ctx context.Context, n node.Interface) ([]peer.Interface, error)
	// SearchOutputPeers returns the shared knowledge about the peers the given node provides
	// signals as output for.
	SearchOutputPeers(ctx context.Context, n node.Interface) ([]peer.Interface, error)
}
