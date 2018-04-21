package storage

type Interface interface {
	SearchRandom() (string, error)
}
