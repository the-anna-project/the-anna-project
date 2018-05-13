package basicnode

import "github.com/the-anna-project/the-anna-project/port"

// OutputPorts returns the information of the nodes the current node forwards
// its signal to.
func (o *Object) OutputPorts() []port.Interface {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.outputPorts
}
