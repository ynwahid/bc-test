package cache

import (
	"context"

	"github.com/ynwahid/bc-test/internal/entity"
)

func (s *CacheService) GetFlights(ctx context.Context, origin, destination, departureDate string) []entity.Flight {
	return s.cache.GetFlightsCache(ctx, origin, destination, departureDate)
}

func (s *CacheService) SetFlightsCache(ctx context.Context, origin, destination, departureDate string, data []entity.Flight) {
	s.cache.SetFlightsCache(ctx, origin, destination, departureDate, data)
}
