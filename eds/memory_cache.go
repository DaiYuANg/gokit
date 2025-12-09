package eds

import (
	"time"

	"github.com/eko/gocache/lib/v4/cache"
)

type MemoryCache struct {
	cache *cache.Cache[[]Record]
}

func NewMemoryCache() {

}

func (m MemoryCache) Get(key string) ([]Record, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MemoryCache) Set(key string, value []Record, ttl time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (m MemoryCache) Delete(key string) {
	//TODO implement me
	panic("implement me")
}
