package basicnetwork

import (
	"context"

	"github.com/the-anna-project/the-anna-project/action"
	"github.com/the-anna-project/the-anna-project/node"
)

// TODO implement node and action deregistration
func (o *Object) DeregisterNode(ctx context.Context, n node.Interface, a action.Interface) error {
	return nil
}
