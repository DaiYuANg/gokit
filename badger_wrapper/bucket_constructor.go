package badger_wrapper

import (
	"github.com/DaiYuANg/gokit/codec"
	"github.com/dgraph-io/badger/v4"
)

type Bucket[T any] struct {
	db        *badger.DB
	keyPrefix []byte
	codec     codec.Codec
}

func NewBucket[T any](db *badger.DB, keyPrefix string) *Bucket[T] {
	prefix := keyPrefix
	if prefix != "" && prefix[len(prefix)-1] != ':' {
		prefix += ":"
	}
	return &Bucket[T]{
		db:        db,
		keyPrefix: []byte(prefix),
	}
}
