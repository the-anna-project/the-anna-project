package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
	"github.com/the-anna-project/the-anna-project/peer/basicpeer"
	"github.com/the-anna-project/the-anna-project/storage/redisstorage"
)

func (o *Object) CreateInputPeers(ctx context.Context, node node.Interface) error {
	// Decide how many peers should be created and create the desired amount if
	// possible.
	var inputPeers map[string]peer.Interface
	{
		peerCount, err := o.random.NewInt(ctx, 0, 10)
		if err != nil {
			return microerror.Mask(err)
		}

		if peerCount == 0 {
			return nil
		}

		var repeatCount int

		for {
			nodeID, err := o.storage.Node.Random()
			if redisstorage.IsNotFound(err) {
				break
			} else if err != nil {
				return microerror.Mask(err)
			}

			c := basicpeer.Config{
				NodeID: nodeID,
			}

			p, err := basicpeer.New(c)
			if err != nil {
				return microerror.Mask(err)
			}

			_, ok := inputPeers[p.NodeID()]
			if ok {
				repeatCount++
				if repeatCount == 3 {
					break
				}

				continue
			}

			inputPeers[p.NodeID()] = p

			if len(inputPeers) == peerCount {
				break
			}
		}
	}

	// Persist input peer and output peer relations relative to the current node.
	{
		for _, p := range inputPeers {
			err := o.storage.Peer.Input.Create(node.ID(), p.NodeID())
			if err != nil {
				return microerror.Mask(err)
			}
		}

		for _, p := range inputPeers {
			err := o.storage.Peer.Output.Create(p.NodeID(), node.ID())
			if err != nil {
				return microerror.Mask(err)
			}
		}
	}

	return nil
}
