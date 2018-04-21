package basicnode

import (
	"context"
)

// Threshold returns the level of resistance a node currently has. A value being
// bigger than the level of vitality as returned by Energy results in the node
// being activated.
func (o *Object) Threshold() float64 {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.threshold
}

// TODO update threshold using node resistance
// TODO random node resistance has to be assigned at node boot
func (o *Object) updateThreshold(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return nil
}
