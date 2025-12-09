package bblot_wrapper

import (
	"fmt"

	"go.etcd.io/bbolt"
)

// Put 将任意类型 T 写入指定 bucket
func Put[T any](w *Wrapper, bucket string, key string, value T) error {
	data, err := w.codec.Encode(value)
	if err != nil {
		return fmt.Errorf("encode failed: %w", err)
	}

	return w.engine.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("create bucket failed: %w", err)
		}
		return b.Put([]byte(key), data)
	})
}

// Delete 删除指定 key
func Delete(w *Wrapper, bucket string, key string) error {
	return w.engine.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		return b.Delete([]byte(key))
	})
}
