package json_codec

type JSONOption func(*JSONCodec)

// Strict 模式：不允许未知字段、检测 trailing garbage
func WithStrictJSON() JSONOption {
	return func(c *JSONCodec) {
		c.disallowUnknown = true
	}
}

// Pretty 打印：仅建议 Debug 时使用
func WithPrettyJSON() JSONOption {
	return func(c *JSONCodec) {
		c.pretty = true
	}
}
