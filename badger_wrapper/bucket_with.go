package badger_wrapper

import "github.com/dgraph-io/badger/v4"

func (r *Bucket[T]) WithReadTx(fn func(txn *badger.Txn) error) error {
	return r.db.View(fn)
}

// 支持读写事务嵌套调用。
func (r *Bucket[T]) WithTx(fn func(txn *badger.Txn) error) error {
	return r.db.Update(fn)
}
