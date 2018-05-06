package storageset

import (
	"github.com/giantswarm/microerror"
	"github.com/the-anna-project/the-anna-project/storage"
)

type StorageSetPeer struct {
	Input  storage.Interface
	Output storage.Interface
}

func (s StorageSetPeer) Validate() error {
	if s.Input == nil {
		return microerror.Maskf(invalidConfigError, "%T.Input must not be empty", s)
	}
	if s.Output == nil {
		return microerror.Maskf(invalidConfigError, "%T.Output must not be empty", s)
	}

	return nil
}
