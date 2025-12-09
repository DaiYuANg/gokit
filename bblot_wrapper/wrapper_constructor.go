package bblot_wrapper

import (
	"github.com/samber/lo"
	"github.com/samber/oops"
	"go.etcd.io/bbolt"
)

func NewWrapper(opts ...Option) (*Wrapper, error) {
	cfg := defaultConfig()

	lo.ForEach(opts, func(opt Option, index int) {
		opt(cfg)
	})

	db, err := bbolt.Open(cfg.path, cfg.mode, cfg.options)
	if err != nil {
		return nil, oops.With("path", cfg.path).Wrapf(err, "failed to open bbolt database")
	}

	return &Wrapper{
		engine: db,
		path:   cfg.path,
		codec:  cfg.codec,
	}, nil
}
