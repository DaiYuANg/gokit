package bblot_wrapper

import (
	"github.com/DaiYuANg/gokit/codec"
	"go.etcd.io/bbolt"
)

type Wrapper struct {
	engine *bbolt.DB
	path   string
	codec  codec.Codec
}

// ===== Close =====
func (w *Wrapper) Close() error {
	if w.engine == nil {
		return nil
	}
	return w.engine.Close()
}
