package filter

import (
	"slices"
	"time"

	"github.com/ynwahid/bc-test/internal/entity"
)

// noStopFilter indicates that the users will obtain
// all numbers of stops, both direct and indirect flights.
const noStopFilter = -1

func (s *FilterService) FilterFlights(param FilterParameter, flights []entity.Flight) []entity.Flight {
	if param.FilterByOrigin != "" {
		flights = s.filterByOrigin(param.FilterByOrigin, flights)
	}

	if param.FilterByDestination != "" {
		flights = s.filterByDestination(param.FilterByDestination, flights)
	}

	if param.FilterByDate != "" {
		flights = s.filterByDate(param.FilterByDate, flights)
	}

	if param.FilterByMinPrice > 0 &&
		param.FilterByMaxPrice > 0 &&
		param.FilterByMaxPrice > param.FilterByMinPrice {
		flights = s.filterByPriceRange(
			param.FilterByMinPrice,
			param.FilterByMaxPrice,
			flights,
		)
	}

	if param.FilterByNumberOfStops != noStopFilter {
		flights = s.filterByNumberOfStops(param.FilterByNumberOfStops, flights)
	}

	if !(param.FilterByDepartureTimeMin.IsZero() &&
		param.FilterByDepartureTimeMax.IsZero()) {
		flights = s.filterByDepartureTime(param.FilterByDepartureTimeMin, param.FilterByDepartureTimeMax, flights)
	}

	if !(param.FilterByArrivalTimeMin.IsZero() &&
		param.FilterByArrivalTimeMax.IsZero()) {
		flights = s.filterByArrivalTime(param.FilterByArrivalTimeMin, param.FilterByArrivalTimeMax, flights)
	}

	if len(param.FilterByAirlinesMap) > 0 {
		flights = s.filterByAirlinesMap(param.FilterByAirlinesMap, flights)
	}

	if param.FilterByDurationMaxInMinutes > 0 &&
		param.FilterByDurationMinInMinutes <= param.FilterByDurationMaxInMinutes {
		flights = s.filterByDuration(param.FilterByDurationMinInMinutes, param.FilterByDurationMaxInMinutes, flights)
	}
	return flights
}

func (s *FilterService) filterByOrigin(origin string, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if flight.Departure.Airport == origin {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByDestination(destination string, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if flight.Arrival.Airport == destination {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByDate(date string, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if flight.Departure.Datetime[:10] == date {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByPriceRange(minPrice, maxPrice int, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if minPrice <= flight.Amount && flight.Amount <= maxPrice {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByNumberOfStops(numberOfStops int, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if flight.Stops == numberOfStops {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByDepartureTime(departureMin, departureMax time.Time, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if isWithinTimeRange(flight.Departure.Datetime, departureMin, departureMax) {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByArrivalTime(arrivalMin, arrivalMax time.Time, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if isWithinTimeRange(flight.Arrival.Datetime, arrivalMin, arrivalMax) {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByAirlinesMap(airlinesMap map[string]bool, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if airlinesMap[flight.Airline.Name] {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func (s *FilterService) filterByDuration(durationMinInMinutes, durationMaxInMinutes int, flights []entity.Flight) []entity.Flight {
	return slices.Collect(
		func(yield func(entity.Flight) bool) {
			for _, flight := range flights {
				if durationMinInMinutes <= flight.TotalMinutes &&
					flight.TotalMinutes <= durationMaxInMinutes {
					if !yield(flight) {
						return
					}
				}
			}
		},
	)
}

func isWithinTimeRange(flightTime string, startTime, endTime time.Time) bool {
	parsedFlightTime, err := time.Parse(time.RFC3339, flightTime)
	if err != nil {
		return false
	}

	year, month, day := parsedFlightTime.Date()
	location := parsedFlightTime.Location()

	startDateTime := time.Date(
		year,
		month,
		day,
		startTime.Hour(),
		startTime.Minute(),
		0,
		0,
		location,
	)

	endDateTime := time.Date(
		year,
		month,
		day,
		endTime.Hour(),
		endTime.Minute(),
		0,
		0,
		location,
	)

	// Handle crossing midnight, example: 23:30 - 03:00
	if endDateTime.Before(startDateTime) ||
		endDateTime.Equal(startDateTime) {
		return (parsedFlightTime.After(startDateTime) ||
			parsedFlightTime.Equal(startDateTime)) ||
			parsedFlightTime.Before(endDateTime)
	}

	// Normal case
	return (parsedFlightTime.After(startDateTime) ||
		parsedFlightTime.Equal(startDateTime)) ||
		parsedFlightTime.Before(endDateTime)
}
