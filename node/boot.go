package node

import (
	"context"
	"time"

	"github.com/giantswarm/microerror"
)

const (
	oneMinute = 1 * time.Minute
)

func (o *Object) Boot(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if o.alreadyBooted {
		return nil
	}
	o.alreadyBooted = true

	var err error

	{
		o.energy, err = o.random.NewFloat64(ctx, 0, 1)
		if err != nil {
			return microerror.Mask(err)
		}
		o.threshold, err = o.random.NewFloat64(ctx, o.energy, 1)
		if err != nil {
			return microerror.Mask(err)
		}
		o.id, err = o.random.NewString(ctx, 32)
		if err != nil {
			return microerror.Mask(err)
		}
		o.peers, err = o.network.RandomPeers(ctx)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		err = o.network.Register(ctx, o, o.action)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		go func() {
			for {
				select {
				case <-ctx.Done():
					o.Shutdown(ctx)
				case <-o.shutdown:
					return
				case <-time.After(oneMinute):
					if o.Energy() > 0 {
						continue
					}

					o.Shutdown(ctx)
				}
			}
		}()
	}

	return nil
}
