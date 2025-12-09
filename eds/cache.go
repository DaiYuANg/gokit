package eds

import "time"

// Cache 抽象接口（可用 gocache 或自定义）
type Cache interface {
	Get(key string) ([]Record, bool)
	Set(key string, value []Record, ttl time.Duration)
	Delete(key string)
}
