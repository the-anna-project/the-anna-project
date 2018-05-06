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

// Boot is responsible for setting up the node's internal states and keeping
// them in sync. The following processes are also ensured periodically.
//
//     Initializing the node's energy and threshold level.
//
//     Initializing the node's ID.
//
//     Registering the node in the network for general awareness.
//
//     Creating the node's input peers ones so it can receive signals.
//
//     Keeping the node's input peers in sync. Input peers might die and go away
//     but once initialized never increase.
//
//     Keeping the node's output peers in sync. Output peers might die and go
//     away or get added during runtime.
//
//     Checking the node's energy level continuously. The node naturally decays
//     if not being used. In case the energy level goes below 0 the node dies
//     by initiating the shutdown process.
//
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
		err = o.network.CreateNode(ctx, o)
		if err != nil {
			return microerror.Mask(err)
		}

		err = o.network.CreateInputPeers(ctx, o)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-o.shutdown:
				return
			case <-time.After(oneMinute):
				peers, err := o.network.SearchInputPeers(context.Background(), o)
				if err != nil {
					o.logger.Log("level", "warning", "message", "searching input peers failed", "stack", fmt.Sprintf("%#v", err))
					continue
				}

				o.mutex.Lock()
				o.inputPeers = peers
				o.mutex.Unlock()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-o.shutdown:
				return
			case <-time.After(oneMinute):
				peers, err := o.network.SearchOutputPeers(context.Background(), o)
				if err != nil {
					o.logger.Log("level", "warning", "message", "searching output peers failed", "stack", fmt.Sprintf("%#v", err))
					continue
				}

				o.mutex.Lock()
				o.outputPeers = peers
				o.mutex.Unlock()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				o.Shutdown(context.Background())
			case <-o.shutdown:
				return
			case <-time.After(oneMinute):
				o.decreaseEnergy(context.Background())

				if o.Energy() > 0 {
					continue
				}

				o.Shutdown(context.Background())
			}
		}
	}()

	{
		o.mutex.Lock()
		o.alreadyBooted = true
		o.mutex.Unlock()
	}

	return nil
}
