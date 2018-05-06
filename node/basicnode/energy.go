package basicnode

import "context"

// Energy returns the level of vitality a node currently has. A negative value
// being returned results in the node being put down.
func (o *Object) Energy() float64 {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.energy
}

// TODO decrease energy
func (o *Object) decreaseEnergy(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return nil
}

// TODO increase energy
func (o *Object) increaseEnergy(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return nil
}
