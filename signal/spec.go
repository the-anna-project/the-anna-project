package signal

import "github.com/the-anna-project/the-anna-project/peer"

type Interface interface {
	Copy() Interface
	SetPeer(p peer.Interface)
}
