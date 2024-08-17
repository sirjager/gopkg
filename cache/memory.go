package cache

import (
	"context"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type inMemoryCache struct {
	data sync.Map
}

// item represents a cache item with a value and an expiration timestamp.
type item struct {
	value      interface{}
	expiration int64
}

// NewInMemoryCache creates a new in-memory cache instance.
func NewInMemoryCache() Cache {
	return &inMemoryCache{}
}

func (c *inMemoryCache) GetKeys(ctx context.Context, pattern string) ([]string, error) {
	var keys []string
	c.data.Range(func(key, value interface{}) bool {
		k := key.(string)
		if matchPattern(k, pattern) {
			keys = append(keys, k)
		}
		return true
	})
	return keys, nil
}

func (c *inMemoryCache) GetKeysWithPrefix(ctx context.Context, prefix string) ([]string, error) {
	var keys []string
	c.data.Range(func(key, value interface{}) bool {
		k := key.(string)
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
		return true
	})
	return keys, nil
}

func (c *inMemoryCache) Get(ctx context.Context, key string, value interface{}) error {
	v, ok := c.data.Load(key)
	if !ok {
		return ErrNoRecord
	}

	it := v.(item)
	if it.expiration > 0 && time.Now().UnixNano() > it.expiration {
		c.data.Delete(key)
		return ErrNoRecord
	}

	valuePtr := value.(*interface{})
	*valuePtr = it.value
	return nil
}

func (c *inMemoryCache) GetWithPrefix(ctx context.Context, prefix string, values any) error {
	var results []interface{}
	c.data.Range(func(key, value interface{}) bool {
		k := key.(string)
		if strings.HasPrefix(k, prefix) {
			it := value.(item)
			if it.expiration == 0 || time.Now().UnixNano() <= it.expiration {
				results = append(results, it.value)
			}
		}
		return true
	})

	valuePtr := values.(*[]interface{})
	*valuePtr = results
	return nil
}

func (c *inMemoryCache) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration ...time.Duration,
) error {
	var exp int64
	if len(expiration) > 0 {
		exp = time.Now().Add(expiration[0]).UnixNano()
	}
	c.data.Store(key, item{
		value:      value,
		expiration: exp,
	})
	return nil
}

func (c *inMemoryCache) Delete(ctx context.Context, keys ...string) error {
	for _, key := range keys {
		c.data.Delete(key)
	}
	return nil
}

func (c *inMemoryCache) DeleteWithPrefix(ctx context.Context, prefix string) error {
	c.data.Range(func(key, value interface{}) bool {
		k := key.(string)
		if strings.HasPrefix(k, prefix) {
			c.data.Delete(k)
		}
		return true
	})
	return nil
}

func (c *inMemoryCache) Flush(ctx context.Context) error {
	c.data.Range(func(key, value interface{}) bool {
		c.data.Delete(key)
		return true
	})
	return nil
}

// matchPattern checks if the key matches the pattern.
func matchPattern(key, pattern string) bool {
	matched, _ := filepath.Match(pattern, key)
	return matched
}
