package noopaction

import (
	"context"

	"github.com/the-anna-project/the-anna-project/signal"
)

func (o *Object) Execute(ctx context.Context, sigs []signal.Interface) ([]signal.Interface, error) {
	return sigs, nil
}
