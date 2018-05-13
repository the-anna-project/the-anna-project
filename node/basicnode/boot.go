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
//     Creating the node's input ports ones so it can receive signals.
//
//     Keeping the node's input ports in sync. Input ports might die and go away
//     but once initialized never increase.
//
//     Keeping the node's output ports in sync. Output ports might die and go
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

		err = o.network.CreateInputPorts(ctx, o)
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
				ports, err := o.network.SearchInputPorts(context.Background(), o)
				if err != nil {
					o.logger.Log("level", "warning", "message", "searching input ports failed", "stack", fmt.Sprintf("%#v", err))
					continue
				}

				o.mutex.Lock()
				o.inputPorts = ports
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
				ports, err := o.network.SearchOutputPorts(context.Background(), o)
				if err != nil {
					o.logger.Log("level", "warning", "message", "searching output ports failed", "stack", fmt.Sprintf("%#v", err))
					continue
				}

				o.mutex.Lock()
				o.outputPorts = ports
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
