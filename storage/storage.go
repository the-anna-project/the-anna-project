package storage

type Interface interface {
	Create(key string, val string) error
	Delete(key string, val string) error
	Random() (string, error)
}
