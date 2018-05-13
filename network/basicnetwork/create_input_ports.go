package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/port"
	"github.com/the-anna-project/the-anna-project/port/basicport"
	"github.com/the-anna-project/the-anna-project/storage/redisstorage"
)

func (o *Object) CreateInputPorts(ctx context.Context, node node.Interface) error {
	// Decide how many ports should be created and create the desired amount if
	// possible.
	var inputPorts map[string]port.Interface
	{
		portCount, err := o.random.NewInt(ctx, 0, 10)
		if err != nil {
			return microerror.Mask(err)
		}

		if portCount == 0 {
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

			c := basicport.Config{
				NodeID: nodeID,
			}

			p, err := basicport.New(c)
			if err != nil {
				return microerror.Mask(err)
			}

			_, ok := inputPorts[p.NodeID()]
			if ok {
				repeatCount++
				if repeatCount == 3 {
					break
				}

				continue
			}

			inputPorts[p.NodeID()] = p

			if len(inputPorts) == portCount {
				break
			}
		}
	}

	// Persist input port and output port relations relative to the current node.
	{
		for _, p := range inputPorts {
			err := o.storage.Port.Input.AddToSet(node.ID(), p.NodeID())
			if err != nil {
				return microerror.Mask(err)
			}
		}

		for _, p := range inputPorts {
			err := o.storage.Port.Output.AddToSet(p.NodeID(), node.ID())
			if err != nil {
				return microerror.Mask(err)
			}
		}
	}

	return nil
}
