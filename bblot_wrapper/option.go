package bblot_wrapper

import (
	"os"
	"time"

	"github.com/DaiYuANg/gokit/codec"
	"github.com/DaiYuANg/gokit/codec/json_codec"
	"go.etcd.io/bbolt"
)

type Option func(*config)

// ===== 默认配置 =====
func defaultConfig() *config {
	return &config{
		path: "bbolt.db",
		mode: 0600,
		options: &bbolt.Options{
			Timeout: 1 * time.Second,
		},
		codec: json_codec.NewJSONCodec(), // 默认 JSON
	}
}

type config struct {
	path    string
	mode    os.FileMode
	options *bbolt.Options
	codec   codec.Codec
}

func WithPath(path string) Option {
	return func(c *config) {
		c.path = path
	}
}

func WithMode(mode os.FileMode) Option {
	return func(c *config) {
		c.mode = mode
	}
}

func WithOptions(opt *bbolt.Options) Option {
	return func(c *config) {
		c.options = opt
	}
}

func WithCodec(c codec.Codec) Option {
	return func(cfg *config) {
		cfg.codec = c
	}
}
