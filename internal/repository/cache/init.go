package cache

func NewCacheRepository(cache CacheInterface) *CacheRepository {
	return &CacheRepository{
		cache: cache,
	}
}
