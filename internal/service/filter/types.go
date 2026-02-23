package filter

import "time"

type FilterService struct{}

type FilterParameter struct {
	// Search
	FilterByOrigin      string
	FilterByDestination string
	FilterByDate        string

	// Filter
	FilterByMinPrice             int
	FilterByMaxPrice             int
	FilterByNumberOfStops        int
	FilterByDepartureTimeMin     time.Time
	FilterByDepartureTimeMax     time.Time
	FilterByArrivalTimeMin       time.Time
	FilterByArrivalTimeMax       time.Time
	FilterByAirlinesMap          map[string]bool
	FilterByDurationMinInMinutes int
	FilterByDurationMaxInMinutes int
}

type SortParameter struct {
	// Sort
	SortByPrice     int
	SortByDuration  int
	SortByDeparture int
	SortByArrival   int
}
