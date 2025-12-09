package eds

type Record struct {
	Name  string
	Type  uint16
	Value string
	TTL   uint32
}
