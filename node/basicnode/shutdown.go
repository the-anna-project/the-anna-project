package basicnode

import (
	"context"

	"github.com/giantswarm/microerror"
)

func (o *Object) Shutdown(ctx context.Context) error {
	var err error

	{
		o.mutex.Lock()
		if o.alreadyShutDown {
			o.mutex.Unlock()
			return nil
		}
		o.mutex.Unlock()
	}

	{
		err = o.network.DeleteNode(ctx, o, o.action)
		if err != nil {
			return microerror.Mask(err)
		}
		err = o.network.DeleteInputPeers(ctx, o, o.inputPeers)
		if err != nil {
			return microerror.Mask(err)
		}
		err = o.network.DeleteOutputPeers(ctx, o, o.outputPeers)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		o.mutex.Lock()
		o.alreadyShutDown = true
		o.mutex.Unlock()
	}

	close(o.shutdown)

	return nil
}
