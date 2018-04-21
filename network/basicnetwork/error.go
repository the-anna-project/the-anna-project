package basicnetwork

import "github.com/giantswarm/microerror"

var alreadyDeregisteredError = microerror.New("already deregistered")

// IsAlreadyDeregistered asserts alreadyDeregisteredError.
func IsAlreadyDeregistered(err error) bool {
	return microerror.Cause(err) == alreadyDeregisteredError
}

var alreadyRegisteredError = microerror.New("already registered")

// IsAlreadyRegistered asserts alreadyRegisteredError.
func IsAlreadyRegistered(err error) bool {
	return microerror.Cause(err) == alreadyRegisteredError
}

var invalidConfigError = microerror.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}
