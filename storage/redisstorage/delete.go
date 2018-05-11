package redisstorage

import (
	"github.com/cenkalti/backoff"
	"github.com/giantswarm/microerror"
	"github.com/gomodule/redigo/redis"
)

func (o *Object) Delete(key string) error {
	operation := func() error {
		conn := o.pool.Get()
		defer conn.Close()

		_, err := redis.Int64(conn.Do("DEL", key))
		if IsNotFound(err) {
			return backoff.Permanent(notFoundError)
		} else if err != nil {
			return microerror.Mask(err)
		}

		return nil
	}

	err := backoff.RetryNotify(operation, o.backoffFactory(), o.retryNotifier)
	if IsNotFound(err) {
		// fall through
	} else if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
