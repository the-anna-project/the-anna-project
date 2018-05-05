package action

import (
	"context"

	"github.com/the-anna-project/the-anna-project/signal"
)

type Interface interface {
	Execute(ctx context.Context, sigs []signal.Interface) ([]signal.Interface, error)
	ID() string
}
