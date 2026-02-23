package sort

import (
	"sort"

	"github.com/ynwahid/bc-test/internal/entity"
)

func (s *SortService) SortFlights(param SortParameter, flights []entity.Flight) {
	s.sortByPrice(param.SortByPrice, flights)
	s.sortByDuration(param.SortByDuration, flights)
	s.sortByDeparture(param.SortByDeparture, flights)
	s.sortByArrival(param.SortByArrival, flights)
}

func (s *SortService) sortByPrice(mode int, flights []entity.Flight) {
	switch mode {
	case 1:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Amount < flights[j].Amount
		})
	case 2:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Amount > flights[j].Amount
		})
	default:
	}
}

func (s *SortService) sortByDuration(mode int, flights []entity.Flight) {
	switch mode {
	case 1:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Duration.TotalMinutes < flights[j].Duration.TotalMinutes
		})
	case 2:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Duration.TotalMinutes > flights[j].Duration.TotalMinutes
		})
	default:
	}
}

func (s *SortService) sortByDeparture(mode int, flights []entity.Flight) {
	switch mode {
	case 1:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Departure.Timestamp < flights[j].Departure.Timestamp
		})
	case 2:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Departure.Timestamp > flights[j].Departure.Timestamp
		})
	default:
	}
}

func (s *SortService) sortByArrival(mode int, flights []entity.Flight) {
	switch mode {
	case 1:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Arrival.Timestamp < flights[j].Arrival.Timestamp
		})
	case 2:
		sort.SliceStable(flights, func(i, j int) bool {
			return flights[i].Arrival.Timestamp > flights[j].Arrival.Timestamp
		})
	default:
	}
}
