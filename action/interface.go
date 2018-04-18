package action

import (
	"context"

	"github.com/the-anna-project/the-anna-project/signal"
)

type Interface interface {
	Execute(ctx context.Context, sig signal.Interface) (signal.Interface, error)
}
