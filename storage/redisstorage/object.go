package redisstorage

import (
	"github.com/cenkalti/backoff"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/gomodule/redigo/redis"
)

type Config struct {
	BackoffFactory func() backoff.BackOff
	Logger         micrologger.Logger
	Pool           *redis.Pool
}

type Object struct {
	backoffFactory func() backoff.BackOff
	logger         micrologger.Logger
	pool           *redis.Pool
}

func New(config Config) (*Object, error) {
	if config.BackoffFactory == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.BackoffFactory must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Pool == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Pool must not be empty", config)
	}

	o := &Object{
		backoffFactory: config.BackoffFactory,
		logger:         config.Logger,
		pool:           config.Pool,
	}

	return o, nil
}
