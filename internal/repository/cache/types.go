package cache

import (
	"time"
)

type CacheInterface interface {
	// Get an item from the cache. Returns the item or nil, and a bool indicating
	// whether the key was found.
	Get(k string) (any, bool)

	// Add an item to the cache, replacing any existing item, using the default
	// expiration.
	SetDefault(k string, x any)
}

type CacheRepository struct {
	defaultTimeout time.Duration
	purgeTimeout   time.Duration
	cache          CacheInterface
}
