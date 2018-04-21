package basicnode

import "github.com/the-anna-project/the-anna-project/peer"

// InputPeers returns the information of the nodes the current node receives
// signals from.
func (o *Object) InputPeers() []peer.Interface {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.inputPeers
}
