package redisstorage

import (
	"github.com/cenkalti/backoff"
	"github.com/giantswarm/microerror"
	"github.com/gomodule/redigo/redis"
)

func (o *Object) SearchSet(key string) ([]string, error) {
	var err error
	var result []string

	operation := func() error {
		conn := o.pool.Get()
		defer conn.Close()

		values, err := redis.Values(conn.Do("SMEMBERS", key))
		if err != nil {
			return microerror.Mask(err)
		}

		for _, v := range values {
			result = append(result, string(v.([]uint8)))
		}

		return nil
	}

	err = backoff.RetryNotify(operation, o.backoffFactory(), o.retryNotifier)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return result, nil
}
