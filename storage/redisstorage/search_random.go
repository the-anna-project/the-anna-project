package redisstorage

import (
	"github.com/cenkalti/backoff"
	"github.com/giantswarm/microerror"
	"github.com/gomodule/redigo/redis"
)

func (o *Object) SearchRandom() (string, error) {
	var err error
	var result string

	operation := func() error {
		conn := o.pool.Get()
		defer conn.Close()

		result, err = redis.String(conn.Do("RANDOMKEY"))
		if IsNotFound(err) {
			return backoff.Permanent(notFoundError)
		} else if err != nil {
			return microerror.Mask(err)
		}

		return nil
	}

	err = backoff.RetryNotify(operation, o.backoffFactory(), o.retryNotifier)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return result, nil
}
