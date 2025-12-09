package json_codec

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/samber/oops"
	"github.com/stretchr/testify/assert"
)

// 测试结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestJSONCodec_Format(t *testing.T) {
	c := NewJSONCodec()
	assert.Equal(t, "json", c.Format())
}

func TestJSONCodec_Encode_Normal(t *testing.T) {
	c := NewJSONCodec()

	u := User{ID: 1, Name: "Alice"}
	data, err := c.Encode(u)
	assert.NoError(t, err)

	var v User
	err = json.Unmarshal(data, &v)
	assert.NoError(t, err)
	assert.Equal(t, u, v)
}

func TestJSONCodec_Encode_Pretty(t *testing.T) {
	c := NewJSONCodec(WithPrettyJSON())

	u := User{ID: 2, Name: "Pretty"}
	data, err := c.Encode(u)
	assert.NoError(t, err)

	assert.Contains(t, string(data), "\n")
}

func TestJSONCodec_Decode_Normal(t *testing.T) {
	c := NewJSONCodec()

	jsonStr := `{"id": 10, "name": "Bob"}`
	var u User
	err := c.Decode([]byte(jsonStr), &u)
	assert.NoError(t, err)

	assert.Equal(t, 10, u.ID)
	assert.Equal(t, "Bob", u.Name)
}

func TestJSONCodec_Decode_InvalidJSON(t *testing.T) {
	c := NewJSONCodec()

	var u User
	err := c.Decode([]byte(`{"id":10,`), &u)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &oops.OopsError{}) // 包装成 oops 错误
}

func TestJSONCodec_Decode_DisallowUnknown(t *testing.T) {
	c := NewJSONCodec(WithStrictJSON())

	jsonStr := `{"id": 1, "name": "Charlie", "extra": 999}`
	var u User

	err := c.Decode([]byte(jsonStr), &u)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown field")
}

func TestJSONCodec_Decode_TrailingGarbage(t *testing.T) {
	c := NewJSONCodec(WithStrictJSON())

	jsonStr := `{"id": 1, "name": "Tail"} xyz`

	var u User
	err := c.Decode([]byte(jsonStr), &u)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "trailing data")
}

func TestJSONCodec_Decode_TypeMismatch(t *testing.T) {
	c := NewJSONCodec()

	var u User
	err := c.Decode([]byte(`123`), &u)
	assert.Error(t, err)
}

func TestJSONCodec_Decode_EmptyData(t *testing.T) {
	c := NewJSONCodec()

	var u User
	err := c.Decode([]byte(``), &u)
	assert.Error(t, err)
}

func TestJSONCodec_Encode_ErrorCase(t *testing.T) {
	c := NewJSONCodec()

	// json.Marshal 的错误场景：chan / func 无法被序列化
	ch := make(chan int)
	_, err := c.Encode(ch)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &oops.OopsError{})
}

func TestJSONCodec_ConcurrentSafety(t *testing.T) {
	c := NewJSONCodec(WithStrictJSON())

	wg := sync.WaitGroup{}

	for i := 0; i < 200; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			u := User{ID: i, Name: "U"}

			data, err := c.Encode(u)
			assert.NoError(t, err)

			var out User
			err = c.Decode(data, &out)
			assert.NoError(t, err)
			assert.Equal(t, u, out)
		}(i)
	}

	wg.Wait()
}
