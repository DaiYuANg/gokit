package codec

import (
	"github.com/samber/lo"
	"github.com/samber/oops"
)

type CompositeCodec struct {
	base       Codec
	middleware []Middleware
}

func New(base Codec, middleware ...Middleware) Codec {
	if base == nil {
		panic("codec: base codec is nil")
	}
	return &CompositeCodec{
		base:       base,
		middleware: middleware,
	}
}

func (c *CompositeCodec) Format() string {
	return c.base.Format()
}

func (c *CompositeCodec) Encode(value any) ([]byte, error) {
	data, err := c.base.Encode(value)
	if err != nil {
		return nil, oops.Wrapf(err, "base.Encode failed (format=%s)", c.Format())
	}

	for _, mw := range c.middleware {
		data, err = mw.BeforeEncode(data)
		if err != nil {
			return nil, oops.Wrapf(err, "middleware.BeforeEncode failed")
		}
	}

	return data, nil
}

func (c *CompositeCodec) Decode(data []byte, value any) error {
	// reverse middleware execution using lo
	reversed := lo.Reverse(c.middleware)

	var err error
	for _, mw := range reversed {
		data, err = mw.AfterDecode(data)
		if err != nil {
			return oops.Wrapf(err, "middleware.AfterDecode failed")
		}
	}

	if err := c.base.Decode(data, value); err != nil {
		return oops.Wrapf(err, "base.Decode failed (format=%s)", c.Format())
	}

	return nil
}
