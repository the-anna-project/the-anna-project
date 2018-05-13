package stopaction

import (
	"context"

	"github.com/the-anna-project/the-anna-project/signal"
)

func (o *Object) Execute(ctx context.Context, sigs []signal.Interface) ([]signal.Interface, error) {
	// TODO check if config exists
	// TODO read config if exists
	// TODO check if config matches signal input
	// TODO return stop signal if config matches signal input
	return sigs, nil
}
