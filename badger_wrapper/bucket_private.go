package badger_wrapper

func (r *Bucket[T]) prefixed(key string) []byte {
	if len(r.keyPrefix) == 0 {
		return []byte(key)
	}
	return append(r.keyPrefix, key...)
}
