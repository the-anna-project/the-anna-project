package redisstorage

import (
	"github.com/cenkalti/backoff"
	"github.com/giantswarm/microerror"
	"github.com/gomodule/redigo/redis"
)

func (o *Object) RemoveFromSet(key string, val string) error {
	operation := func() error {
		conn := o.pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SREM", key, val))
		if err != nil {
			return microerror.Mask(err)
		}

		return nil
	}

	err := backoff.RetryNotify(operation, o.backoffFactory(), o.retryNotifier)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
