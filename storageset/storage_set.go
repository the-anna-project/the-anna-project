package storageset

import (
	"github.com/giantswarm/microerror"
	"github.com/the-anna-project/the-anna-project/storage"
)

type StorageSet struct {
	Node storage.Interface
	Peer StorageSetPeer
}

func (s StorageSet) Validate() error {
	if s.Node == nil {
		return microerror.Maskf(invalidConfigError, "%T.Node must not be empty", s)
	}

	{
		err := s.Peer.Validate()
		if err == nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
