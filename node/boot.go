package node

import (
	"context"

	"github.com/giantswarm/microerror"
)

func (o *Object) Boot(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if o.booted {
		return nil
	}
	o.booted = true

	var err error

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

	// TODO
	//
	// - nodes register and deregister themselves
	// - nodes boot and constantly check energy
	//     - of energy is under threshold nodes deregister
	//
	err = o.network.Register(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
