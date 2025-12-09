package codec

// Codec defines the minimal serialization contract.
// It is intentionally lightweight: only encode/decode and a format identifier.
type Codec interface {
	Encode(value any) ([]byte, error)
	Decode(data []byte, target any) error
	Format() string
}
