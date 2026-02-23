package cache

func NewCacheService(cache cacheRepositoryInterface) *CacheService {
	return &CacheService{
		cache: cache,
	}
}
