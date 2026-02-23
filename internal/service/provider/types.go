package provider

import (
	"context"

	"github.com/ynwahid/bc-test/internal/repository/api"
)

type ProviderInterface interface {
	GetAirAsiaFlights(ctx context.Context, origin, destination, departureDate string) (api.AirAsiaResponse, error)
	GetBatikAirFlights(ctx context.Context, origin, destination, departureDate string) (api.BatikAirResponse, error)
	GetGarudaIndonesiaFlights(ctx context.Context, origin, destination, departureDate string) (api.GarudaIndonesiaResponse, error)
	GetLionAirFlights(ctx context.Context, origin, destination, departureDate string) (api.LionAirResponse, error)
}

type ProviderService struct {
	api ProviderInterface
}
