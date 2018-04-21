package node

import "context"

// Energy returns the level of vitality a node currently has. A negative value
// being returned results in the node being put down.
func (o *Object) Energy() float64 {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.energy
}

// TODO update energy using node resistance
// TODO random node resistance has to be assigned at node boot
func (o *Object) updateEnergy(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return nil
}
