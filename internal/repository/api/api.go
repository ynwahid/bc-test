package api

import (
	"context"
	"encoding/json"
)

func (repo *APIRepository) GetAirAsiaFlights(ctx context.Context, origin, destination, departureDate string) (AirAsiaResponse, error) {
	var response AirAsiaResponse
	err := json.Unmarshal([]byte(airAsiaMockResponse), &response)
	if err != nil {
		return AirAsiaResponse{}, err
	}
	return response, nil
}

func (repo *APIRepository) GetBatikAirFlights(ctx context.Context, origin, destination, departureDate string) (BatikAirResponse, error) {
	var response BatikAirResponse
	err := json.Unmarshal([]byte(batikAirMockResponse), &response)
	if err != nil {
		return BatikAirResponse{}, err
	}
	return response, nil
}

func (repo *APIRepository) GetGarudaIndonesiaFlights(ctx context.Context, origin, destination, departureDate string) (GarudaIndonesiaResponse, error) {
	var response GarudaIndonesiaResponse
	err := json.Unmarshal([]byte(garudaIndonesiaMockResponse), &response)
	if err != nil {
		return GarudaIndonesiaResponse{}, err
	}
	return response, nil
}

func (repo *APIRepository) GetLionAirFlights(ctx context.Context, origin, destination, departureDate string) (LionAirResponse, error) {
	var response LionAirResponse
	err := json.Unmarshal([]byte(lionAirMockResponse), &response)
	if err != nil {
		return LionAirResponse{}, err
	}
	return response, nil
}
