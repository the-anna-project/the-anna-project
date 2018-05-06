package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/peer"
	"github.com/the-anna-project/the-anna-project/peer/basicpeer"
)

// TODO implement CreateInputPeers
func (o *Object) CreateInputPeers(ctx context.Context, n node.Interface) error {
	// Decide how many peers should be created.
	var inputPeers map[string]peer.Interface
	{
		peerCount, err := o.random.NewFloat64(ctx, 0, 10)
		if err != nil {
			return microerror.Mask(err)
		}

		if peerCount == 0 {
			return nil
		}

		var repeatCount int

		for {
			nodeID, err := o.storage.Node.SearchRandom()
			if err != nil {
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

	return nil
}
