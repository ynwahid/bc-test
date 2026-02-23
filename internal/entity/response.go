package entity

type ResultResponse struct {
	SearchCriteria `json:"search_criteria"`
	Metadata       `json:"metadata"`
	Flights        []Flight `json:"flights"`
}

type SearchCriteria struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departure_date"`
	Passengers    int    `json:"passengers"`
	CabinClass    string `json:"cabin_class"`
}

type Metadata struct {
	TotalResults       int  `json:"total_results"`
	ProvidersQueried   int  `json:"providers_queried"`
	ProvidersSucceeded int  `json:"providers_succeeded"`
	ProvidersFailed    int  `json:"providers_failed"`
	SearchTimeMs       int  `json:"search_time_ms"`
	CacheHit           bool `json:"cache_hit"`
}

type Flight struct {
	ID             string `json:"id"`
	Provider       string `json:"provider"`
	Airline        `json:"airline"`
	FlightNumber   string `json:"flight_number"`
	Departure      `json:"departure"`
	Arrival        `json:"arrival"`
	Duration       `json:"duration"`
	Stops          int `json:"stops"`
	Price          `json:"price"`
	AvailableSeats int      `json:"available_seats"`
	CabinClass     string   `json:"cabin_class"`
	Aircraft       any      `json:"aircraft"`
	Amenities      []string `json:"amenities"`
	Baggage        `json:"baggage"`
}

type Airline struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Departure struct {
	Airport   string `json:"airport"`
	City      string `json:"city"`
	Datetime  string `json:"datetime"`
	Timestamp int    `json:"timestamp"`
}

type Arrival struct {
	Airport   string `json:"airport"`
	City      string `json:"city"`
	Datetime  string `json:"datetime"`
	Timestamp int    `json:"timestamp"`
}

type Duration struct {
	TotalMinutes int    `json:"total_minutes"`
	Formatted    string `json:"formatted"`
}

type Price struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type Baggage struct {
	CarryOn string `json:"carry_on"`
	Checked string `json:"checked"`
}
