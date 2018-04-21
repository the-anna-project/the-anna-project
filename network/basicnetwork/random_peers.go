package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/peer"
	"github.com/the-anna-project/the-anna-project/peer/basicpeer"
)

func (o *Object) RandomPeers(ctx context.Context) ([]peer.Interface, error) {
	var max int
	// TODO get random number
	// TODO     get most successful maximum boundary
	// TODO     use default value if no known successful maximum boundary is known

	var peers []peer.Interface

	for i := 0; i < max; i++ {
		var err error
		var result string

		result, err = o.storage.SearchRandom()
		if err != nil {
			return nil, microerror.Mask(err)
		}

		var p peer.Interface
		{
			c := basicpeer.Config{
				NodeID: result,
			}

			p, err = basicpeer.New(c)
			if err != nil {
				return nil, microerror.Mask(err)
			}
		}

		peers = append(peers, p)
	}

	return peers, nil
}
