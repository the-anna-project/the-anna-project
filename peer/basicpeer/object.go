package basicpeer

import (
	"sync"

	"github.com/giantswarm/microerror"
)

type Config struct {
	NodeID string
}

type Object struct {
	mutex sync.Mutex

	nodeID string
}

func New(config Config) (*Object, error) {
	if config.NodeID == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.NodeID must not be empty", config)
	}

	o := &Object{
		mutex: sync.Mutex{},

		nodeID: config.NodeID,
	}

	return o, nil
}
