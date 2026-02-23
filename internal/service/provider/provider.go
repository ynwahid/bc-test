package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ynwahid/bc-test/internal/entity"
	"github.com/ynwahid/bc-test/internal/repository/provider"
)

const (
	timeLayout        = `2006-01-02T15:04:05-07:00`
	timeLayoutLionAir = `2006-01-02T15:04:05`
)

func (p *ProviderService) GetAirAsiaFlights(ctx context.Context, origin, destination, departureDate string) ([]entity.Flight, error) {
	res, err := p.provider.GetAirAsiaFlights(ctx, origin, destination, departureDate)
	if err != nil {
		return []entity.Flight{}, err
	}
	return p.transformAirAsiaResponseToFlights(res)
}

func (p *ProviderService) GetBatikAirFlights(ctx context.Context, origin, destination, departureDate string) ([]entity.Flight, error) {
	res, err := p.provider.GetBatikAirFlights(ctx, origin, destination, departureDate)
	if err != nil {
		return []entity.Flight{}, err
	}
	return p.transformBatikAirResponseToFlights(res)
}

func (p *ProviderService) GetGarudaIndonesiaFlights(ctx context.Context, origin, destination, departureDate string) ([]entity.Flight, error) {
	res, err := p.provider.GetGarudaIndonesiaFlights(ctx, origin, destination, departureDate)
	if err != nil {
		return []entity.Flight{}, err
	}
	return p.transformGarudaIndonesiaResponseToFlights(res)
}

func (p *ProviderService) GetLionAirFlights(ctx context.Context, origin, destination, departureDate string) ([]entity.Flight, error) {
	res, err := p.provider.GetLionAirFlights(ctx, origin, destination, departureDate)
	if err != nil {
		return []entity.Flight{}, err
	}
	return p.transformLionAirResponseToFlights(res)
}

func (p *ProviderService) transformAirAsiaResponseToFlights(response provider.AirAsiaResponse) ([]entity.Flight, error) {
	flights := make([]entity.Flight, len(response.Flights))

	for i, flight := range response.Flights {
		departureTime, err := time.Parse(timeLayout, flight.DepartTime)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse AirAsia departure time: %s", err.Error())
		}

		arrivalTime, err := time.Parse(timeLayout, flight.ArriveTime)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse AirAsia arrival time: %s", err.Error())
		}

		duration, err := time.ParseDuration(fmt.Sprintf("%fh", flight.DurationHours))
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse AirAsia duration hours: %s", err.Error())
		}
		formatted := fmt.Sprintf(
			"%dh %dm",
			int(duration.Hours()),
			int(duration.Minutes())%60,
		)
		flights[i] = entity.Flight{
			ID:       fmt.Sprintf("%s_%s", flight.FlightCode, flight.Airline),
			Provider: "AirAsia",
			Airline: entity.Airline{
				Name: flight.Airline,
				Code: "QZ",
			},
			FlightNumber: flight.FlightCode,
			Departure: entity.Departure{
				Airport:   flight.FromAirport,
				City:      entity.AirportMap[flight.FromAirport],
				Datetime:  flight.DepartTime,
				Timestamp: int(departureTime.Unix()),
			},
			Arrival: entity.Arrival{
				Airport:   flight.ToAirport,
				City:      entity.AirportMap[flight.ToAirport],
				Datetime:  flight.ArriveTime,
				Timestamp: int(arrivalTime.Unix()),
			},
			Duration: entity.Duration{
				TotalMinutes: int(duration.Minutes()),
				Formatted:    formatted,
			},
			Stops: len(flight.Stops),
			Price: entity.Price{
				Amount:   flight.PriceIDR,
				Currency: "IDR",
			},
			AvailableSeats: flight.Seats,
			CabinClass:     flight.CabinClass,
			Aircraft:       nil,
			Amenities:      []string{},
			Baggage: entity.Baggage{
				CarryOn: "Cabin baggage only",
				Checked: "Additional fee",
			},
		}
	}
	return flights, nil
}

func (p *ProviderService) transformBatikAirResponseToFlights(response provider.BatikAirResponse) ([]entity.Flight, error) {
	flights := make([]entity.Flight, len(response.Results))

	for i, flight := range response.Results {
		departureTime, err := time.Parse(timeLayout, flight.DepartureDateTime)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse Batik Air departure time: %s", err.Error())
		}

		arrivalTime, err := time.Parse(timeLayout, flight.ArrivalDateTime)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse Batik Air arrival time: %s", err.Error())
		}

		duration, err := time.ParseDuration(strings.ReplaceAll(flight.TravelTime, " ", ""))
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse Batik Air duration hours: %s", err.Error())
		}

		carryOn := "7kg cabin"
		checked := "20kg checked"
		baggageInfo := strings.Split(strings.ReplaceAll(flight.BaggageInfo, " ", ""), ",")
		if len(baggageInfo) == 2 {
			carryOn = baggageInfo[0]
			checked = baggageInfo[1]
		}

		flights[i] = entity.Flight{
			ID:       fmt.Sprintf("%s_%s", flight.FlightNumber, flight.AirlineName),
			Provider: "Batik Air",
			Airline: entity.Airline{
				Name: flight.AirlineName,
				Code: flight.AirlineIATA,
			},
			FlightNumber: flight.FlightNumber,
			Departure: entity.Departure{
				Airport:   flight.Origin,
				City:      entity.AirportMap[flight.Origin],
				Datetime:  flight.DepartureDateTime,
				Timestamp: int(departureTime.Unix()),
			},
			Arrival: entity.Arrival{
				Airport:   flight.Destination,
				City:      entity.AirportMap[flight.Destination],
				Datetime:  flight.ArrivalDateTime,
				Timestamp: int(arrivalTime.Unix()),
			},
			Duration: entity.Duration{
				TotalMinutes: int(duration.Minutes()),
				Formatted:    flight.TravelTime,
			},
			Stops: flight.NumberOfStops,
			Price: entity.Price{
				Amount:   flight.Fare.TotalPrice,
				Currency: flight.Fare.CurrencyCode,
			},
			AvailableSeats: flight.SeatsAvailable,
			CabinClass:     entity.ClassMap[flight.Fare.Class],
			Aircraft:       flight.AircraftModel,
			Amenities:      flight.OnboardServices,
			Baggage: entity.Baggage{
				CarryOn: carryOn,
				Checked: checked,
			},
		}
	}
	return flights, nil
}

func (p *ProviderService) transformGarudaIndonesiaResponseToFlights(response provider.GarudaIndonesiaResponse) ([]entity.Flight, error) {
	flights := make([]entity.Flight, len(response.Flights))

	for i, flight := range response.Flights {
		departureTime, err := time.Parse(timeLayout, flight.Departure.Time)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse Garuda Indonesia departure time: %s", err.Error())
		}

		arrivalTime, err := time.Parse(timeLayout, flight.Arrival.Time)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse Garuda Indonesia arrival time: %s", err.Error())
		}

		duration := time.Minute * time.Duration(flight.DurationMinutes)
		formatted := fmt.Sprintf(
			"%dh %dm",
			int(duration.Hours()),
			int(duration.Minutes())%60,
		)

		flights[i] = entity.Flight{
			ID:       fmt.Sprintf("%s_%s", flight.FlightID, flight.Airline),
			Provider: "Garuda Indonesia",
			Airline: entity.Airline{
				Name: flight.Airline,
				Code: flight.AirlineCode,
			},
			FlightNumber: flight.FlightID,
			Departure: entity.Departure{
				Airport:   flight.Departure.Airport,
				City:      flight.Departure.City,
				Datetime:  flight.Departure.Time,
				Timestamp: int(departureTime.Unix()),
			},
			Arrival: entity.Arrival{
				Airport:   flight.Arrival.Airport,
				City:      flight.Arrival.City,
				Datetime:  flight.Arrival.Time,
				Timestamp: int(arrivalTime.Unix()),
			},
			Duration: entity.Duration{
				TotalMinutes: flight.DurationMinutes,
				Formatted:    formatted,
			},
			Stops: len(flight.Segments),
			Price: entity.Price{
				Amount:   flight.Price.Amount,
				Currency: flight.Price.Currency,
			},
			AvailableSeats: flight.AvailableSeats,
			CabinClass:     flight.FareClass,
			Aircraft:       flight.Aircraft,
			Amenities:      flight.Amenities,
			Baggage: entity.Baggage{
				CarryOn: fmt.Sprintf("%d", (flight.Baggage.CarryOn)),
				Checked: fmt.Sprintf("%d", (flight.Baggage.Checked)),
			},
		}
	}
	return flights, nil
}

func (p *ProviderService) transformLionAirResponseToFlights(response provider.LionAirResponse) ([]entity.Flight, error) {
	flights := make([]entity.Flight, len(response.Data.AvailableFlights))

	for i, flight := range response.Data.AvailableFlights {
		departureTZLocation, err := time.LoadLocation(flight.Schedule.DepartureTimezone)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to load Lion Air departure timezone: %s", err.Error())
		}

		departureTimeWithoutTZ, err := time.Parse(timeLayoutLionAir, flight.Schedule.Departure)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse Lion Air departure time: %s", err.Error())
		}

		departureTime := time.Date(
			departureTimeWithoutTZ.Year(),
			departureTimeWithoutTZ.Month(),
			departureTimeWithoutTZ.Day(),
			departureTimeWithoutTZ.Hour(),
			departureTimeWithoutTZ.Minute(),
			departureTimeWithoutTZ.Second(),
			departureTimeWithoutTZ.Nanosecond(),
			departureTZLocation,
		)

		arrivalTZLocation, err := time.LoadLocation(flight.Schedule.ArrivalTimezone)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to load Lion Air arrival timezone: %s", err.Error())
		}

		arrivalTimeWithoutTZ, err := time.Parse(timeLayoutLionAir, flight.Schedule.Arrival)
		if err != nil {
			return []entity.Flight{}, fmt.Errorf("failed to parse Lion Air arrival time: %s", err.Error())
		}

		arrivalTime := time.Date(
			arrivalTimeWithoutTZ.Year(),
			arrivalTimeWithoutTZ.Month(),
			arrivalTimeWithoutTZ.Day(),
			arrivalTimeWithoutTZ.Hour(),
			arrivalTimeWithoutTZ.Minute(),
			arrivalTimeWithoutTZ.Second(),
			arrivalTimeWithoutTZ.Nanosecond(),
			arrivalTZLocation,
		)

		duration := time.Minute * time.Duration(flight.FlightTime)
		formatted := fmt.Sprintf(
			"%dh %dm",
			int(duration.Hours()),
			int(duration.Minutes())%60,
		)

		var amenities []string
		if flight.Services.WifiAvailable {
			amenities = append(amenities, "wifi")
		}
		if flight.Services.MealsIncluded {
			amenities = append(amenities, "meals")
		}

		flights[i] = entity.Flight{
			ID:       fmt.Sprintf("%s_%s", flight.ID, flight.Carrier.Name),
			Provider: "Lion Air",
			Airline: entity.Airline{
				Name: flight.Carrier.Name,
				Code: flight.Carrier.Iata,
			},
			FlightNumber: flight.ID,
			Departure: entity.Departure{
				Airport:   flight.Route.From.Code,
				City:      flight.Route.From.City,
				Datetime:  departureTime.Format(timeLayout),
				Timestamp: int(departureTime.Unix()),
			},
			Arrival: entity.Arrival{
				Airport:   flight.Route.To.Code,
				City:      flight.Route.To.City,
				Datetime:  arrivalTime.Format(timeLayout),
				Timestamp: int(arrivalTime.Unix()),
			},
			Duration: entity.Duration{
				TotalMinutes: flight.FlightTime,
				Formatted:    formatted,
			},
			Stops: flight.StopCount,
			Price: entity.Price{
				Amount:   flight.Pricing.Total,
				Currency: flight.Pricing.Currency,
			},
			AvailableSeats: flight.SeatsLeft,
			CabinClass:     flight.Pricing.FareType,
			Aircraft:       flight.PlaneType,
			Amenities:      amenities,
			Baggage: entity.Baggage{
				CarryOn: flight.Services.BaggageAllowance.Cabin,
				Checked: flight.Services.BaggageAllowance.Hold,
			},
		}
	}
	return flights, nil
}
