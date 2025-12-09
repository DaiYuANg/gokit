package badger_wrapper

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/goccy/go-json"
)

// Set 或 Update
func (r *Bucket[T]) Set(key string, value T) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.db.Update(func(txn *badger.Txn) error {
		return txn.Set(r.prefixed(key), raw)
	})
}

func (r *Bucket[T]) Delete(key string) error {
	rawKey := r.prefixed(key)
	return r.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(rawKey)
	})
}

// 支持批量写操作。
func (r *Bucket[T]) BatchSet(items map[string]T) error {
	return r.db.Update(func(txn *badger.Txn) error {
		for k, v := range items {
			raw, err := json.Marshal(v)
			if err != nil {
				return err
			}
			if err := txn.Set(r.prefixed(k), raw); err != nil {
				return err
			}
		}
		return nil
	})
}
