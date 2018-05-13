package network

import (
	"context"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/port"
)

type Interface interface {
	// CreateInputPorts creates input ports for the given node based on random
	// guesses. It also creates output ports implicitly on the other side. This is
	// because an input port for one node is an output port for the other.
	CreateInputPorts(ctx context.Context, node node.Interface) error
	// CreateNode adds shared knowledge about the given node and its associated
	// action.
	CreateNode(ctx context.Context, node node.Interface) error
	// DeleteInputPorts deletes the given references of input ports from the given
	// node.
	DeleteInputPorts(ctx context.Context, node node.Interface, ports []port.Interface) error
	// DeleteNode removes shared knowledge about the given node and the given
	// action from the configured storage.
	DeleteNode(ctx context.Context, node node.Interface) error
	// DeleteOutputPorts deletes the given references of output ports from the
	// given node.
	DeleteOutputPorts(ctx context.Context, node node.Interface, ports []port.Interface) error
	// SearchInputPorts returns the shared knowledge about the ports providing
	// signals as input for the given node.
	SearchInputPorts(ctx context.Context, node node.Interface) ([]port.Interface, error)
	// SearchOutputPorts returns the shared knowledge about the ports the given
	// node provides signals as output for.
	SearchOutputPorts(ctx context.Context, node node.Interface) ([]port.Interface, error)
}
