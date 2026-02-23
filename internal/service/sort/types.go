package sort

type SortService struct{}

// The kind of sort:
//
// 1 -> Ascending
//
// 2 -> Descending
type SortParameter struct {
	SortByPrice     int
	SortByDuration  int
	SortByDeparture int
	SortByArrival   int
}
