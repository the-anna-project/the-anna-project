package redisstorage

import (
	"github.com/giantswarm/microerror"
	"github.com/gomodule/redigo/redis"
)

var invalidConfigError = microerror.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var notFoundError = microerror.New("not found")

// IsNotFound asserts notFoundError. It also checks whether a redis response was
// empty. Therefore it checks for redigo.ErrNil and notFoundError.
//
//     ErrNil indicates that a reply value is nil.
//
func IsNotFound(err error) bool {
	c := microerror.Cause(err)

	if c == notFoundError {
		return true
	}

	if c == redis.ErrNil {
		return true
	}

	return false
}
