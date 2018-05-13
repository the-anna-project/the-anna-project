package basicnode

import "context"

const (
	DefaultResistence float64 = 0.001
)

// Energy returns the level of vitality a node currently has. A negative value
// being returned results in the node being put down.
func (o *Object) Energy() float64 {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.energy
}

func (o *Object) decreaseEnergy(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.energy -= DefaultResistence

	return nil
}

func (o *Object) increaseEnergy(ctx context.Context) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.energy += DefaultResistence

	return nil
}
