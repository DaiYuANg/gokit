package json_codec

import (
	"bytes"
	"encoding/json"

	"github.com/DaiYuANg/gokit/codec"
	"github.com/samber/oops"
)

type JSONCodec struct {
	// 严格模式：禁止未知字段
	// 建议在配置文件或协议数据的场景中开启
	disallowUnknown bool

	// pretty 输出，仅在 Debug/测试使用
	pretty bool
}

func NewJSONCodec(opts ...JSONOption) codec.Codec {
	c := &JSONCodec{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *JSONCodec) Format() string {
	return "json"
}

func (c *JSONCodec) Encode(value any) ([]byte, error) {
	var (
		data []byte
		err  error
	)

	if c.pretty {
		data, err = json.MarshalIndent(value, "", "  ")
	} else {
		data, err = json.Marshal(value)
	}

	if err != nil {
		return nil, oops.Wrapf(err, "json: marshal failed (type=%T)", value)
	}

	return data, nil
}

func (c *JSONCodec) Decode(data []byte, value any) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	if c.disallowUnknown {
		dec.DisallowUnknownFields()
	}

	if err := dec.Decode(value); err != nil {
		return oops.
			With("data", string(data)).
			Wrapf(err, "json: decode failed (type=%T)", value)
	}

	// 再读一次，确认没有多余数据（防止 trailing garbage）
	if c.disallowUnknown && dec.More() {
		return oops.Errorf("json: trailing data detected during strict decoding")
	}

	return nil
}
