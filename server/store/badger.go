package store

import (
	badger "github.com/dgraph-io/badger/v2"
)

type badgerTxn struct {
	store *badgerKeyValueStore
}

func (txn *badgerTxn) set(key, value []byte) error {

	return nil
}

func (txn *badgerTxn) get(key []byte) ([]byte, error) {
	return nil, nil
}

type badgerKeyValueStore struct {
	path string
	db   *badger.DB
}

func newBadgerKeyValueStore(path string) *badgerKeyValueStore {
	return &badgerKeyValueStore{
		path: path,
	}
}

func (s *badgerKeyValueStore) open() (err error) {
	s.db, err = badger.Open(badger.DefaultOptions(s.path))
	return err
}

func (s *badgerKeyValueStore) close() error {
	return s.db.Close()
}

func (s *badgerKeyValueStore) set(key, value []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (s *badgerKeyValueStore) get(key []byte) (value []byte, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == badger.ErrKeyNotFound {
			return nil
		} else if err != nil {
			return err
		}

		value, err = item.ValueCopy(nil)
		return err
	})

	return value, err
}

func (s *badgerKeyValueStore) view(txnFunc txnFunc) error {
	txn := s.db.NewTransaction(false)
	defer txn.Discard()

	if err := txnFunc(&badgerTxn{
		store: s,
	}); err != nil {
		return err
	}

	return txn.Commit()
}

func (s *badgerKeyValueStore) update(txnFunc txnFunc) error {
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	if err := txnFunc(&badgerTxn{
		store: s,
	}); err != nil {
		return err
	}

	return txn.Commit()
}
