package store

import (
	"github.com/charmbracelet/charm/client"
	"github.com/charmbracelet/charm/kv"
	"github.com/dgraph-io/badger/v3"
	"github.com/gabriel-vasile/mimetype"
)

type Store struct {
	db *kv.KV
}

func NewStore(name string) (*Store, error) {
	db, err := openKV(name)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Link(code string) error {
	cc, err := client.NewClientWithDefaults()
	if err != nil {
		return err
	}
	lh := newLinkHandler()
	cc.Link(lh, code)
	return nil
}

func (s *Store) GetFibers(name string) ([]Fiber, error) {
	err := s.db.Sync()
	if err != nil {
		return nil, err
	}

	var fibers []Fiber
	err = s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				mtype := mimetype.Detect(v)
				fibers = append(fibers, Fiber{TS: string(k), Data: v, Mimetype: mtype.String()})
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return fibers, nil
}

func (s *Store) GetDbOptions(name string) badger.Options {
	return s.db.DB.Opts()
}
