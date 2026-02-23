package cache

import (
	"context"

	"github.com/ynwahid/bc-test/internal/entity"
)

type cacheRepositoryInterface interface {
	GetFlightsCache(ctx context.Context, origin, destination, departureDate string) []entity.Flight
	SetFlightsCache(ctx context.Context, origin, destination, departureDate string, data []entity.Flight)
}

type CacheService struct {
	cache cacheRepositoryInterface
}
