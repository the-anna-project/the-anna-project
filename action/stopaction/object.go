package stopaction

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

type Config struct {
	Logger micrologger.Logger
}

type Object struct {
	logger micrologger.Logger
}

func New(config Config) (*Object, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	o := &Object{
		logger: config.Logger,
	}

	return o, nil
}
