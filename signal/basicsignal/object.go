package basicsignal

import (
	"github.com/the-anna-project/the-anna-project/peer"
	"github.com/the-anna-project/the-anna-project/signal"
)

type Object struct {
}

func (o *Object) Copy() signal.Interface {
	return nil
}

func (o *Object) SetPeer(p peer.Interface) {
}
