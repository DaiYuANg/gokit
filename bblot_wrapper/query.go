package bblot_wrapper

import (
	"fmt"

	"go.etcd.io/bbolt"
)

// Get 从指定 bucket 读取 key 并解码为 T
func Get[T any](w *Wrapper, bucket string, key string) (T, error) {
	var zero T
	err := w.engine.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		data := b.Get([]byte(key))
		if data == nil {
			return fmt.Errorf("key %s not found", key)
		}
		return w.codec.Decode(data, &zero)
	})
	return zero, err
}

// Exists 判断 key 是否存在
func Exists(w *Wrapper, bucket string, key string) (bool, error) {
	var exists bool
	err := w.engine.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			exists = false
			return nil
		}
		data := b.Get([]byte(key))
		exists = data != nil
		return nil
	})
	return exists, err
}
