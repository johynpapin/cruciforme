package store

type keyValueStore interface {
	open() error
	close() error
	set(key, value []byte) error
	get(key []byte) ([]byte, error)
}

type Store struct {
	keyValueStore
}

func New() *Store {
	return &Store{
		keyValueStore: newBadgerKeyValueStore("./badger"),
	}
}

func (s *Store) Open() error {
	return s.open()
}

func (s *Store) Close() error {
	return s.close()
}
