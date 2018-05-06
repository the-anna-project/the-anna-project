package node

import (
	"context"

	"github.com/the-anna-project/the-anna-project/action"
	"github.com/the-anna-project/the-anna-project/signal"
)

type Interface interface {
	Action() action.Interface
	Execute(ctx context.Context, sigs []signal.Interface) ([]signal.Interface, error)
	ID() string
}
