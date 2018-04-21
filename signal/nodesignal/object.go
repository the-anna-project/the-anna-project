package nodesignal

import (
	"github.com/the-anna-project/the-anna-project/spec/peer"
	"github.com/the-anna-project/the-anna-project/spec/signal"
)

type Object struct {
}

func (o *Object) Copy() signal.Interface {
	return nil
}

func (o *Object) SetPeer(p peer.Interface) {
}
