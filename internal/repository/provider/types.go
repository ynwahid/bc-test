package provider

type ProviderRepository struct{}

type AirAsiaResponse struct {
	Status  string `json:"status"`
	Flights []struct {
		FlightCode    string  `json:"flight_code"`
		Airline       string  `json:"airline"`
		FromAirport   string  `json:"from_airport"`
		ToAirport     string  `json:"to_airport"`
		DepartTime    string  `json:"depart_time"`
		ArriveTime    string  `json:"arrive_time"`
		DurationHours float64 `json:"duration_hours"`
		DirectFlight  bool    `json:"direct_flight"`
		Stops         []struct {
			Airport         string `json:"airport"`
			WaitTimeMinutes int    `json:"wait_time_minutes"`
		} `json:"stops,omitempty"`
		PriceIDR    int    `json:"price_idr"`
		Seats       int    `json:"seats"`
		CabinClass  string `json:"cabin_class"`
		BaggageNote string `json:"baggage_note"`
	} `json:"flights"`
}

type BatikAirResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Results []struct {
		FlightNumber      string `json:"flightNumber"`
		AirlineName       string `json:"airlineName"`
		AirlineIATA       string `json:"airlineIATA"`
		Origin            string `json:"origin"`
		Destination       string `json:"destination"`
		DepartureDateTime string `json:"departureDateTime"`
		ArrivalDateTime   string `json:"arrivalDateTime"`
		TravelTime        string `json:"travelTime"`
		NumberOfStops     int    `json:"numberOfStops"`
		Connections       []struct {
			StopAirport  string `json:"stopAirport"`
			StopDuration string `json:"stopDuration"`
		} `json:"connections,omitempty"`
		Fare struct {
			BasePrice    int    `json:"basePrice"`
			Taxes        int    `json:"taxes"`
			TotalPrice   int    `json:"totalPrice"`
			CurrencyCode string `json:"currencyCode"`
			Class        string `json:"class"`
		} `json:"fare"`
		SeatsAvailable  int      `json:"seatsAvailable"`
		AircraftModel   string   `json:"aircraftModel"`
		BaggageInfo     string   `json:"baggageInfo"`
		OnboardServices []string `json:"onboardServices"`
	} `json:"results"`
}

type GarudaIndonesiaResponse struct {
	Status  string `json:"status"`
	Flights []struct {
		FlightID    string `json:"flight_id"`
		Airline     string `json:"airline"`
		AirlineCode string `json:"airline_code"`
		Departure   struct {
			Airport  string `json:"airport"`
			City     string `json:"city"`
			Time     string `json:"time"`
			Terminal string `json:"terminal"`
		} `json:"departure"`
		Arrival struct {
			Airport  string `json:"airport"`
			City     string `json:"city"`
			Time     string `json:"time"`
			Terminal string `json:"terminal"`
		} `json:"arrival"`
		DurationMinutes int    `json:"duration_minutes"`
		Stops           int    `json:"stops"`
		Aircraft        string `json:"aircraft"`
		Price           struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"price"`
		Segments []struct {
			FlightNumber string `json:"flight_number"`
			Departure    struct {
				Airport string `json:"airport"`
				Time    string `json:"time"`
			} `json:"departure"`
			Arrival struct {
				Airport string `json:"airport"`
				Time    string `json:"time"`
			} `json:"arrival"`
			DurationMinutes int `json:"duration_minutes"`
			LayoverMinutes  int `json:"layover_minutes,omitempty"`
		} `json:"segments,omitempty"`
		AvailableSeats int    `json:"available_seats"`
		FareClass      string `json:"fare_class"`
		Baggage        struct {
			CarryOn int `json:"carry_on"`
			Checked int `json:"checked"`
		} `json:"baggage"`
		Amenities []string `json:"amenities,omitempty"`
	} `json:"flights"`
}

type LionAirResponse struct {
	Success bool `json:"success"`
	Data    struct {
		AvailableFlights []struct {
			ID      string `json:"id"`
			Carrier struct {
				Name string `json:"name"`
				Iata string `json:"iata"`
			} `json:"carrier"`
			Route struct {
				From struct {
					Code string `json:"code"`
					Name string `json:"name"`
					City string `json:"city"`
				} `json:"from"`
				To struct {
					Code string `json:"code"`
					Name string `json:"name"`
					City string `json:"city"`
				} `json:"to"`
			} `json:"route"`
			Schedule struct {
				Departure         string `json:"departure"`
				DepartureTimezone string `json:"departure_timezone"`
				Arrival           string `json:"arrival"`
				ArrivalTimezone   string `json:"arrival_timezone"`
			} `json:"schedule"`
			FlightTime int  `json:"flight_time"`
			IsDirect   bool `json:"is_direct"`
			StopCount  int  `json:"stop_count,omitempty"`
			Layovers   []struct {
				Airport         string `json:"airport"`
				DurationMinutes int    `json:"duration_minutes"`
			} `json:"layovers,omitempty"`
			Pricing struct {
				Total    int    `json:"total"`
				Currency string `json:"currency"`
				FareType string `json:"fare_type"`
			} `json:"pricing"`
			SeatsLeft int    `json:"seats_left"`
			PlaneType string `json:"plane_type"`
			Services  struct {
				WifiAvailable    bool `json:"wifi_available"`
				MealsIncluded    bool `json:"meals_included"`
				BaggageAllowance struct {
					Cabin string `json:"cabin"`
					Hold  string `json:"hold"`
				} `json:"baggage_allowance"`
			} `json:"services"`
		} `json:"available_flights"`
	} `json:"data"`
}
