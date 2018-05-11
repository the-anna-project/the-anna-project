package redisstorage

import (
	"github.com/cenkalti/backoff"
	"github.com/giantswarm/microerror"
	"github.com/gomodule/redigo/redis"
)

func (o *Object) Create(key string, val string) error {
	operation := func() error {
		conn := o.pool.Get()
		defer conn.Close()

		reply, err := redis.String(conn.Do("SET", key, val))
		if err != nil {
			return microerror.Mask(err)
		}

		if reply != "OK" {
			return backoff.Permanent(executionFailedError)
		}

		return nil
	}

	err := backoff.RetryNotify(operation, o.backoffFactory(), o.retryNotifier)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
