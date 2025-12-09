package msgpack_codec

import (
	"github.com/hashicorp/go-msgpack/codec"
	"github.com/samber/oops"
)

type MsgPackCodec struct {
	// 你可以加一些配置字段，比如是否排序 map key, 扩展 tag, etc.
	handle *codec.MsgpackHandle
}

func NewMsgPackCodec() *MsgPackCodec {
	h := &codec.MsgpackHandle{
		// 根据需要你可以设置 MapType, SliceType, StructAsArray, RawToString 等选项
		// 例如: h.MapType = reflect.TypeOf(map[string]interface{}{})
	}
	return &MsgPackCodec{handle: h}
}

func (c *MsgPackCodec) Format() string {
	return "msgpack"
}

func (c *MsgPackCodec) Encode(value any) ([]byte, error) {
	var buf []byte
	// EncodeBytes will append to buf
	err := codec.NewEncoderBytes(&buf, c.handle).Encode(value)
	if err != nil {
		return nil, oops.Wrapf(err, "msgpack: encode failed (type=%T)", value)
	}
	return buf, nil
}

func (c *MsgPackCodec) Decode(data []byte, target any) error {
	err := codec.NewDecoderBytes(data, c.handle).Decode(target)
	if err != nil {
		return oops.Wrapf(err, "msgpack: decode failed (type=%T)", target)
	}
	return nil
}
