package basicnode

import (
	"context"
	"fmt"
	"time"

	"github.com/giantswarm/microerror"
)

const (
	oneMinute = 1 * time.Minute
)

func (o *Object) Boot(ctx context.Context) error {
	var err error

	{
		o.mutex.Lock()
		if o.alreadyBooted {
			o.mutex.Unlock()
			return nil
		}
		o.mutex.Unlock()
	}

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
	}

	{
		err = o.network.CreateNode(ctx, o, o.action)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-o.shutdown:
					return
				case <-time.After(oneMinute):
					peers, err := o.network.SearchInputPeers(ctx, o)
					if err != nil {
						o.logger.Log("level", "warning", "message", "searching input peers failed", "stack", fmt.Sprintf("%#v", err))
					}

					o.mutex.Lock()
					o.inputPeers = peers
					o.mutex.Unlock()
				}
			}
		}()
	}

	{
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-o.shutdown:
					return
				case <-time.After(oneMinute):
					peers, err := o.network.SearchOutputPeers(ctx, o)
					if err != nil {
						o.logger.Log("level", "warning", "message", "searching output peers failed", "stack", fmt.Sprintf("%#v", err))
					}

					o.mutex.Lock()
					o.outputPeers = peers
					o.mutex.Unlock()
				}
			}
		}()
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

	{
		o.mutex.Lock()
		o.alreadyBooted = true
		o.mutex.Unlock()
	}

	return nil
}
