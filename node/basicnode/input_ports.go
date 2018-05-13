package basicnode

import "github.com/the-anna-project/the-anna-project/port"

// InputPorts returns the information of the nodes the current node receives
// signals from.
func (o *Object) InputPorts() []port.Interface {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.inputPorts
}
