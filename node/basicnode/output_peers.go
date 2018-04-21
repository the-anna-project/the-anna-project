package basicnode

import "github.com/the-anna-project/the-anna-project/peer"

// OutputPeers returns the information of the nodes the current node forwards
// its signal to.
func (o *Object) OutputPeers() []peer.Interface {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.outputPeers
}
