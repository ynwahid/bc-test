package cache

import (
	"context"
	"fmt"

	"github.com/ynwahid/bc-test/internal/entity"
)

const flightKey = `%s%s%s`

func (repo *CacheRepository) GetFlightsCache(ctx context.Context, origin, destination, departureDate string) []entity.Flight {
	cachedData, found := repo.cache.Get(fmt.Sprintf(flightKey, origin, destination, departureDate))
	if !found {
		return nil
	}

	result := cachedData.([]entity.Flight)
	if len(result) == 0 {
		return nil
	}
	return result
}

func (repo *CacheRepository) SetFlightsCache(ctx context.Context, origin, destination, departureDate string, data []entity.Flight) {
	repo.cache.SetDefault(fmt.Sprintf(flightKey, origin, destination, departureDate), data)
}
