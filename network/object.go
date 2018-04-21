package network

import (
	"github.com/giantswarm/microerror"

	"github.com/the-anna-project/the-anna-project/spec/storage"
)

type Config struct {
	Storage storage.Interface
}

type Object struct {
	storage storage.Interface
}

func New(config Config) (*Object, error) {
	if config.Storage == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Storage must not be empty", config)
	}

	o := &Object{
		storage: config.Storage,
	}

	return o, nil
}
