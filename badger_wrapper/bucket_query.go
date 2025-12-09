package badger_wrapper

import (
	"errors"

	"github.com/dgraph-io/badger/v4"
	"github.com/goccy/go-json"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

// Get 根据 key 获取值
func (r *Bucket[T]) Get(key string) (T, error) {
	var zero T
	rawKey := r.prefixed(key)

	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(rawKey)
		if err != nil {
			return oops.Wrap(err)
		}

		valBytes, err := item.ValueCopy(nil)
		if err != nil {
			return oops.Wrap(err)
		}

		return oops.Wrap(json.Unmarshal(valBytes, &zero))
	})

	return zero, err
}

func (r *Bucket[T]) Exists(key string) (bool, error) {
	rawKey := r.prefixed(key)
	err := r.db.View(func(txn *badger.Txn) error {
		_, e := txn.Get(rawKey)
		if e != nil && !errors.Is(e, badger.ErrKeyNotFound) {
			return oops.Wrap(e)
		}
		return nil
	})
	return err == nil, nil
}

// 返回整个 DB 中（或限定前缀后）的所有 keys。
// 注意：对于整库扫描，性能差，建议配合分页或 stream API。

// ListKeys 返回所有 key（仅 key，不返回 value）
// 注意：大数据量情况下性能差，建议配合分页或流式处理
func (r *Bucket[T]) ListKeys() ([]string, error) {
	var keys []string
	err := r.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(r.keyPrefix); it.ValidForPrefix(r.keyPrefix); it.Next() {
			rawKey := it.Item().KeyCopy(nil)
			trimKey := string(rawKey[len(r.keyPrefix):])
			keys = append(keys, trimKey)
		}
		return nil
	})

	// 可用 lo.Map 对 keys 做额外处理或转换
	return lo.Map(keys, func(k string, _ int) string { return k }), oops.Wrap(err)
}

// 遍历 key/value 并反序列化为 T。
// ForEach 遍历所有 key/value，并反序列化为 T
func (r *Bucket[T]) ForEach(fn func(key string, val T) error) error {
	return oops.Wrap(r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(r.keyPrefix); it.ValidForPrefix(r.keyPrefix); it.Next() {
			item := it.Item()
			valBytes, err := item.ValueCopy(nil)
			if err != nil {
				return oops.Wrap(err)
			}

			var parsed T
			if err := json.Unmarshal(valBytes, &parsed); err != nil {
				return oops.Wrap(err)
			}

			trimKey := string(item.KeyCopy(nil)[len(r.keyPrefix):])
			if err := fn(trimKey, parsed); err != nil {
				return oops.Wrap(err)
			}
		}
		return nil
	}))
}

// 返回条目数（仅做 rough count）。
func (r *Bucket[T]) Count() (int, error) {
	count := 0
	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(r.keyPrefix); it.ValidForPrefix(r.keyPrefix); it.Next() {
			count++
		}
		return nil
	})
	return count, oops.Wrap(err)
}

// 支持前缀扫描，分页处理数据。
func (r *Bucket[T]) PrefixScan(
	prefixStr string,
	callback func(key string, val T) (stop bool, err error),
) error {
	rawPrefix := append(r.keyPrefix, []byte(prefixStr)...)

	return oops.Wrap(r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(rawPrefix); it.ValidForPrefix(rawPrefix); it.Next() {
			item := it.Item()
			valBytes, err := item.ValueCopy(nil)
			if err != nil {
				return oops.Wrap(err)
			}

			var parsed T
			if err := json.Unmarshal(valBytes, &parsed); err != nil {
				return oops.Wrap(err)
			}

			trimKey := string(item.KeyCopy(nil)[len(r.keyPrefix):])
			stop, err := callback(trimKey, parsed)
			if err != nil {
				return oops.Wrap(err)
			}
			if stop {
				break
			}
		}
		return nil
	}))
}
