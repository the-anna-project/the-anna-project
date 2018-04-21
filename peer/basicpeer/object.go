package basicpeer

import "github.com/giantswarm/microerror"

type Config struct {
	NodeID string
}

type Object struct {
	nodeID string
}

func New(config Config) (*Object, error) {
	if config.NodeID == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.NodeID must not be empty", config)
	}

	o := &Object{
		nodeID: config.NodeID,
	}

	return o, nil
}
