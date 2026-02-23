package provider

import (
	"context"

	"github.com/ynwahid/bc-test/internal/repository/provider"
)

type ProviderInterface interface {
	GetAirAsiaFlights(ctx context.Context, origin, destination, departureDate string) (provider.AirAsiaResponse, error)
	GetBatikAirFlights(ctx context.Context, origin, destination, departureDate string) (provider.BatikAirResponse, error)
	GetGarudaIndonesiaFlights(ctx context.Context, origin, destination, departureDate string) (provider.GarudaIndonesiaResponse, error)
	GetLionAirFlights(ctx context.Context, origin, destination, departureDate string) (provider.LionAirResponse, error)
}

type ProviderService struct {
	provider ProviderInterface
}
