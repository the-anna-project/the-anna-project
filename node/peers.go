package node

import (
	"github.com/the-anna-project/the-anna-project/peer"
)

// Peers returns the information of the nodes the current node is connected to.
func (o *Object) Peers() []peer.Interface {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.peers
}
