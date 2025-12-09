package codec

// Middleware defines hooks around serialization.
// Designed to be optional and composable (compression, encryption, checksum, logging, etc.).
type Middleware interface {
	BeforeEncode(data []byte) ([]byte, error)
	AfterDecode(data []byte) ([]byte, error)
}
