package storage

type Interface interface {
	Create(key string, val string) error
	SearchRandom() (string, error)
}
