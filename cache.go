package camp

import (
	"github.com/coocood/freecache"
	"github.com/spf13/cast"
	"github.com/xilloader/camp/interfaces"
)

type cache struct{ cache *freecache.Cache }

func NewFreeCache(size int) interfaces.Cache {
	return &cache{cache: freecache.NewCache(size)}
}

func (c *cache) Set(key string, value interface{}, expireSeconds int64) error {
	var val []byte
	switch v := value.(type) {
	case []byte:
		val = v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		val = IntToBytes(v)
	default:
		s, err := cast.ToStringE(v)
		if err != nil {
			return err
		}
		val = []byte(s)
	}
	return c.cache.Set([]byte(key), val, int(expireSeconds))
}

func (c *cache) Get(key string) (interface{}, error) {
	return c.cache.Get([]byte(key))
}

func (c *cache) Del(key string) bool {
	return c.cache.Del([]byte(key))
}

func (c *cache) Exist(key string) bool {
	timeLeft, err := c.cache.TTL([]byte(key))
	if err != nil || timeLeft <= 0 {
		return false
	}
	return true
}
