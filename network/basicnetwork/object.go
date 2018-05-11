package basicnetwork

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/the-anna-project/the-anna-project/random"
	"github.com/the-anna-project/the-anna-project/storageset"
)

type Config struct {
	Logger  micrologger.Logger
	Storage storageset.StorageSet
	Random  random.Interface
}

type Object struct {
	logger  micrologger.Logger
	storage storageset.StorageSet
	random  random.Interface
}

func New(config Config) (*Object, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	{
		err := config.Storage.Validate()
		if err == nil {
			return nil, microerror.Mask(err)
		}
	}
	if config.Random == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Random must not be empty", config)
	}

	o := &Object{
		logger:  config.Logger,
		storage: config.Storage,
		random:  config.Random,
	}

	return o, nil
}
