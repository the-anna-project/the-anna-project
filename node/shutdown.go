package node

import (
	"context"

	"github.com/giantswarm/microerror"
)

func (o *Object) Shutdown(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if o.alreadyShutDown {
		return nil
	}
	o.alreadyShutDown = true

	var err error

	close(o.shutdown)

	err = o.network.Deregister(ctx, o, o.action)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
