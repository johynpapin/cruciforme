package store

type txn interface {
	set(key, value []byte) error
	get(key []byte) ([]byte, error)
}

type txnFunc func(txn) error

type keyValueStore interface {
	open() error
	close() error
	set(key, value []byte) error
	get(key []byte) ([]byte, error)
	view(txnFunc txnFunc) error
	update(txnFunc txnFunc) error
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
