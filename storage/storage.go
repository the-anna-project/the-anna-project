package storage

type Interface interface {
	AddToSet(key string, val string) error
	Create(key string, val string) error
	Delete(key string) error
	Random() (string, error)
	RemoveFromSet(key string, val string) error
	SearchSet(key string) ([]string, error)
}
