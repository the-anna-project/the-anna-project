package basicnetwork

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/the-anna-project/the-anna-project/node"
	"github.com/the-anna-project/the-anna-project/port"
	"github.com/the-anna-project/the-anna-project/port/basicport"
)

func (o *Object) SearchInputPorts(ctx context.Context, node node.Interface) ([]port.Interface, error) {
	result, err := o.storage.Port.Input.SearchSet(node.ID())
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var ports []port.Interface

	for _, r := range result {
		c := basicport.Config{
			NodeID: r,
		}

		p, err := basicport.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		ports = append(ports, p)
	}

	return ports, nil
}
