package api_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ynwahid/bc-test/internal/repository/api"
)

func TestAPIRepository_GetAirAsiaFlights(t *testing.T) {
	const timeLayout = `2006-01-02T15:04:05-07:00`
	departureQZ520, err := time.Parse(timeLayout, "2025-12-15T04:45:00+07:00")
	require.NoError(t, err)
	arriveQZ520, err := time.Parse(timeLayout, "2025-12-15T07:25:00+08:00")
	require.NoError(t, err)
	departureQZ524, err := time.Parse(timeLayout, "2025-12-15T10:00:00+07:00")
	require.NoError(t, err)
	arriveQZ524, err := time.Parse(timeLayout, "2025-12-15T12:45:00+08:00")
	require.NoError(t, err)
	departureQZ532, err := time.Parse(timeLayout, "2025-12-15T19:30:00+07:00")
	require.NoError(t, err)
	arriveQZ532, err := time.Parse(timeLayout, "2025-12-15T22:10:00+08:00")
	require.NoError(t, err)
	departureQZ7250, err := time.Parse(timeLayout, "2025-12-15T15:15:00+07:00")
	require.NoError(t, err)
	arriveQZ7250, err := time.Parse(timeLayout, "2025-12-15T20:35:00+08:00")
	require.NoError(t, err)

	tests := []struct {
		name          string
		origin        string
		destination   string
		departureDate string
		want          api.AirAsiaResponse
		wantErr       bool
	}{
		{
			name:          "succeed",
			origin:        "CGK",
			destination:   "DPS",
			departureDate: "2025-12-15",
			want: api.AirAsiaResponse{
				Status: "ok",
				Flights: []struct {
					FlightCode    string    "json:\"flight_code\""
					Airline       string    "json:\"airline\""
					FromAirport   string    "json:\"from_airport\""
					ToAirport     string    "json:\"to_airport\""
					DepartTime    time.Time "json:\"depart_time\""
					ArriveTime    time.Time "json:\"arrive_time\""
					DurationHours float64   "json:\"duration_hours\""
					DirectFlight  bool      "json:\"direct_flight\""
					Stops         []struct {
						Airport         string "json:\"airport\""
						WaitTimeMinutes int    "json:\"wait_time_minutes\""
					} "json:\"stops,omitempty\""
					PriceIDR    int    "json:\"price_idr\""
					Seats       int    "json:\"seats\""
					CabinClass  string "json:\"cabin_class\""
					BaggageNote string "json:\"baggage_note\""
				}{
					{
						FlightCode:    "QZ520",
						Airline:       "AirAsia",
						FromAirport:   "CGK",
						ToAirport:     "DPS",
						DepartTime:    departureQZ520,
						ArriveTime:    arriveQZ520,
						DurationHours: 1.67,
						DirectFlight:  true,
						PriceIDR:      650000,
						Seats:         67,
						CabinClass:    "economy",
						BaggageNote:   "Cabin baggage only, checked bags additional fee",
					},
					{
						FlightCode:    "QZ524",
						Airline:       "AirAsia",
						FromAirport:   "CGK",
						ToAirport:     "DPS",
						DepartTime:    departureQZ524,
						ArriveTime:    arriveQZ524,
						DurationHours: 1.75,
						DirectFlight:  true,
						PriceIDR:      720000,
						Seats:         54,
						CabinClass:    "economy",
						BaggageNote:   "Cabin baggage only, checked bags additional fee",
					},
					{
						FlightCode:    "QZ532",
						Airline:       "AirAsia",
						FromAirport:   "CGK",
						ToAirport:     "DPS",
						DepartTime:    departureQZ532,
						ArriveTime:    arriveQZ532,
						DurationHours: 1.67,
						DirectFlight:  true,
						PriceIDR:      595000,
						Seats:         72,
						CabinClass:    "economy",
						BaggageNote:   "Cabin baggage only, checked bags additional fee",
					},
					{
						FlightCode:    "QZ7250",
						Airline:       "AirAsia",
						FromAirport:   "CGK",
						ToAirport:     "DPS",
						DepartTime:    departureQZ7250,
						ArriveTime:    arriveQZ7250,
						DurationHours: 4.33,
						DirectFlight:  false,
						Stops: []struct {
							Airport         string "json:\"airport\""
							WaitTimeMinutes int    "json:\"wait_time_minutes\""
						}{
							{
								Airport:         "SOC",
								WaitTimeMinutes: 95,
							},
						},
						PriceIDR:    485000,
						Seats:       88,
						CabinClass:  "economy",
						BaggageNote: "Cabin baggage only, checked bags additional fee",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := api.NewAPIRepository()
			got, err := repo.GetAirAsiaFlights(
				context.Background(),
				test.origin,
				test.destination,
				test.departureDate,
			)
			if err != nil {
				require.True(t, test.wantErr)
			} else {
				require.False(t, test.wantErr)
			}
			require.Equal(t, test.want, got)
		})
	}
}

func TestAPIRepository_GetBatikAirFlights(t *testing.T) {
	tests := []struct {
		name          string
		origin        string
		destination   string
		departureDate string
		want          api.BatikAirResponse
		wantErr       bool
	}{
		{
			name:          "succeed",
			origin:        "CGK",
			destination:   "DPS",
			departureDate: "2025-12-15",
			want: api.BatikAirResponse{
				Code:    200,
				Message: "OK",
				Results: []struct {
					FlightNumber      string "json:\"flightNumber\""
					AirlineName       string "json:\"airlineName\""
					AirlineIATA       string "json:\"airlineIATA\""
					Origin            string "json:\"origin\""
					Destination       string "json:\"destination\""
					DepartureDateTime string "json:\"departureDateTime\""
					ArrivalDateTime   string "json:\"arrivalDateTime\""
					TravelTime        string "json:\"travelTime\""
					NumberOfStops     int    "json:\"numberOfStops\""
					Connections       []struct {
						StopAirport  string "json:\"stopAirport\""
						StopDuration string "json:\"stopDuration\""
					} "json:\"connections,omitempty\""
					Fare struct {
						BasePrice    int    "json:\"basePrice\""
						Taxes        int    "json:\"taxes\""
						TotalPrice   int    "json:\"totalPrice\""
						CurrencyCode string "json:\"currencyCode\""
						Class        string "json:\"class\""
					} "json:\"fare\""
					SeatsAvailable  int      "json:\"seatsAvailable\""
					AircraftModel   string   "json:\"aircraftModel\""
					BaggageInfo     string   "json:\"baggageInfo\""
					OnboardServices []string "json:\"onboardServices\""
				}{
					{
						FlightNumber:      "ID6514",
						AirlineName:       "Batik Air",
						AirlineIATA:       "ID",
						Origin:            "CGK",
						Destination:       "DPS",
						DepartureDateTime: "2025-12-15T07:15:00+0700",
						ArrivalDateTime:   "2025-12-15T10:00:00+0800",
						TravelTime:        "1h 45m",
						NumberOfStops:     0,
						Fare: struct {
							BasePrice    int    "json:\"basePrice\""
							Taxes        int    "json:\"taxes\""
							TotalPrice   int    "json:\"totalPrice\""
							CurrencyCode string "json:\"currencyCode\""
							Class        string "json:\"class\""
						}{
							BasePrice:    980000,
							Taxes:        120000,
							TotalPrice:   1100000,
							CurrencyCode: "IDR",
							Class:        "Y",
						},
						SeatsAvailable: 32,
						AircraftModel:  "Airbus A320",
						BaggageInfo:    "7kg cabin, 20kg checked",
						OnboardServices: []string{
							"Snack",
							"Beverage",
						},
					},
					{
						FlightNumber:      "ID6520",
						AirlineName:       "Batik Air",
						AirlineIATA:       "ID",
						Origin:            "CGK",
						Destination:       "DPS",
						DepartureDateTime: "2025-12-15T13:30:00+0700",
						ArrivalDateTime:   "2025-12-15T16:20:00+0800",
						TravelTime:        "1h 50m",
						NumberOfStops:     0,
						Fare: struct {
							BasePrice    int    "json:\"basePrice\""
							Taxes        int    "json:\"taxes\""
							TotalPrice   int    "json:\"totalPrice\""
							CurrencyCode string "json:\"currencyCode\""
							Class        string "json:\"class\""
						}{
							BasePrice:    1050000,
							Taxes:        130000,
							TotalPrice:   1180000,
							CurrencyCode: "IDR",
							Class:        "Y",
						},
						SeatsAvailable: 18,
						AircraftModel:  "Boeing 737-800",
						BaggageInfo:    "7kg cabin, 20kg checked",
						OnboardServices: []string{
							"Meal",
							"Beverage",
							"Entertainment",
						},
					},
					{
						FlightNumber:      "ID7042",
						AirlineName:       "Batik Air",
						AirlineIATA:       "ID",
						Origin:            "CGK",
						Destination:       "DPS",
						DepartureDateTime: "2025-12-15T18:45:00+0700",
						ArrivalDateTime:   "2025-12-15T23:50:00+0800",
						TravelTime:        "3h 5m",
						NumberOfStops:     1,
						Connections: []struct {
							StopAirport  string "json:\"stopAirport\""
							StopDuration string "json:\"stopDuration\""
						}{
							{
								StopAirport:  "UPG",
								StopDuration: "55m",
							},
						},
						Fare: struct {
							BasePrice    int    "json:\"basePrice\""
							Taxes        int    "json:\"taxes\""
							TotalPrice   int    "json:\"totalPrice\""
							CurrencyCode string "json:\"currencyCode\""
							Class        string "json:\"class\""
						}{
							BasePrice:    850000,
							Taxes:        100000,
							TotalPrice:   950000,
							CurrencyCode: "IDR",
							Class:        "Y",
						},
						SeatsAvailable:  41,
						AircraftModel:   "Airbus A320",
						BaggageInfo:     "7kg cabin, 20kg checked",
						OnboardServices: []string{"Snack"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := api.NewAPIRepository()
			got, err := repo.GetBatikAirFlights(
				context.Background(),
				test.origin,
				test.destination,
				test.departureDate,
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				require.True(t, test.wantErr)
			} else {
				require.False(t, test.wantErr)
			}
			require.Equal(t, test.want, got)
		})
	}
}

func TestAPIRepository_GetGarudaIndonesiaFlights(t *testing.T) {
	const timeLayout = `2006-01-02T15:04:05-07:00`
	departurGA332, err := time.Parse(timeLayout, "2025-12-15T17:15:00+07:00")
	require.NoError(t, err)
	arriveGA332, err := time.Parse(timeLayout, "2025-12-15T18:45:00+08:00")
	require.NoError(t, err)
	departurGA315, err := time.Parse(timeLayout, "2025-12-15T14:00:00+07:00")
	require.NoError(t, err)
	arriveGA315, err := time.Parse(timeLayout, "2025-12-15T15:30:00+07:00")
	require.NoError(t, err)
	departurGA410, err := time.Parse(timeLayout, "2025-12-15T09:30:00+07:00")
	require.NoError(t, err)
	arriveGA410, err := time.Parse(timeLayout, "2025-12-15T12:25:00+08:00")
	require.NoError(t, err)
	departurGA400, err := time.Parse(timeLayout, "2025-12-15T06:00:00+07:00")
	require.NoError(t, err)
	arriveGA400, err := time.Parse(timeLayout, "2025-12-15T08:50:00+08:00")
	require.NoError(t, err)

	tests := []struct {
		name          string
		origin        string
		destination   string
		departureDate string
		want          api.GarudaIndonesiaResponse
		wantErr       bool
	}{
		{
			name:          "succeed",
			origin:        "CGK",
			destination:   "DPS",
			departureDate: "2025-12-15",
			want: api.GarudaIndonesiaResponse{
				Status: "success",
				Flights: []struct {
					FlightID    string "json:\"flight_id\""
					Airline     string "json:\"airline\""
					AirlineCode string "json:\"airline_code\""
					Departure   struct {
						Airport  string    "json:\"airport\""
						City     string    "json:\"city\""
						Time     time.Time "json:\"time\""
						Terminal string    "json:\"terminal\""
					} "json:\"departure\""
					Arrival struct {
						Airport  string    "json:\"airport\""
						City     string    "json:\"city\""
						Time     time.Time "json:\"time\""
						Terminal string    "json:\"terminal\""
					} "json:\"arrival\""
					DurationMinutes int    "json:\"duration_minutes\""
					Stops           int    "json:\"stops\""
					Aircraft        string "json:\"aircraft\""
					Price           struct {
						Amount   int    "json:\"amount\""
						Currency string "json:\"currency\""
					} "json:\"price\""
					Segments []struct {
						FlightNumber string "json:\"flight_number\""
						Departure    struct {
							Airport string    "json:\"airport\""
							Time    time.Time "json:\"time\""
						} "json:\"departure\""
						Arrival struct {
							Airport string    "json:\"airport\""
							Time    time.Time "json:\"time\""
						} "json:\"arrival\""
						DurationMinutes int "json:\"duration_minutes\""
						LayoverMinutes  int "json:\"layover_minutes,omitempty\""
					} "json:\"segments,omitempty\""
					AvailableSeats int    "json:\"available_seats\""
					FareClass      string "json:\"fare_class\""
					Baggage        struct {
						CarryOn int "json:\"carry_on\""
						Checked int "json:\"checked\""
					} "json:\"baggage\""
					Amenities []string "json:\"amenities,omitempty\""
				}{
					{
						FlightID:    "GA400",
						Airline:     "Garuda Indonesia",
						AirlineCode: "GA",
						Departure: struct {
							Airport  string    "json:\"airport\""
							City     string    "json:\"city\""
							Time     time.Time "json:\"time\""
							Terminal string    "json:\"terminal\""
						}{
							Airport:  "CGK",
							City:     "Jakarta",
							Time:     departurGA400,
							Terminal: "3",
						},
						Arrival: struct {
							Airport  string    "json:\"airport\""
							City     string    "json:\"city\""
							Time     time.Time "json:\"time\""
							Terminal string    "json:\"terminal\""
						}{
							Airport:  "DPS",
							City:     "Denpasar",
							Time:     arriveGA400,
							Terminal: "I",
						},
						DurationMinutes: 110,
						Stops:           0,
						Aircraft:        "Boeing 737-800",
						Price: struct {
							Amount   int    "json:\"amount\""
							Currency string "json:\"currency\""
						}{
							Amount:   1250000,
							Currency: "IDR",
						},
						AvailableSeats: 28,
						FareClass:      "economy",
						Baggage: struct {
							CarryOn int "json:\"carry_on\""
							Checked int "json:\"checked\""
						}{
							CarryOn: 1,
							Checked: 2,
						},
						Amenities: []string{
							"wifi",
							"meal",
							"entertainment",
						},
					},
					{
						FlightID:    "GA410",
						Airline:     "Garuda Indonesia",
						AirlineCode: "GA",
						Departure: struct {
							Airport  string    "json:\"airport\""
							City     string    "json:\"city\""
							Time     time.Time "json:\"time\""
							Terminal string    "json:\"terminal\""
						}{
							Airport:  "CGK",
							City:     "Jakarta",
							Time:     departurGA410,
							Terminal: "3",
						},
						Arrival: struct {
							Airport  string    "json:\"airport\""
							City     string    "json:\"city\""
							Time     time.Time "json:\"time\""
							Terminal string    "json:\"terminal\""
						}{
							Airport:  "DPS",
							City:     "Denpasar",
							Time:     arriveGA410,
							Terminal: "I",
						},
						DurationMinutes: 115,
						Stops:           0,
						Aircraft:        "Airbus A330-300",
						Price: struct {
							Amount   int    "json:\"amount\""
							Currency string "json:\"currency\""
						}{
							Amount:   1450000,
							Currency: "IDR",
						},
						AvailableSeats: 15,
						FareClass:      "economy",
						Baggage: struct {
							CarryOn int "json:\"carry_on\""
							Checked int "json:\"checked\""
						}{
							CarryOn: 1,
							Checked: 2,
						},
						Amenities: []string{
							"wifi",
							"power_outlet",
							"meal",
							"entertainment",
						},
					},
					{
						FlightID:    "GA315",
						Airline:     "Garuda Indonesia",
						AirlineCode: "GA",
						Departure: struct {
							Airport  string    "json:\"airport\""
							City     string    "json:\"city\""
							Time     time.Time "json:\"time\""
							Terminal string    "json:\"terminal\""
						}{
							Airport:  "CGK",
							City:     "Jakarta",
							Time:     departurGA315,
							Terminal: "3",
						},
						Arrival: struct {
							Airport  string    "json:\"airport\""
							City     string    "json:\"city\""
							Time     time.Time "json:\"time\""
							Terminal string    "json:\"terminal\""
						}{
							Airport:  "SUB",
							City:     "Surabaya",
							Time:     arriveGA315,
							Terminal: "2",
						},
						DurationMinutes: 90,
						Stops:           0,
						Aircraft:        "Boeing 737",
						Price: struct {
							Amount   int    "json:\"amount\""
							Currency string "json:\"currency\""
						}{
							Amount:   1850000,
							Currency: "IDR",
						},
						Segments: []struct {
							FlightNumber string "json:\"flight_number\""
							Departure    struct {
								Airport string    "json:\"airport\""
								Time    time.Time "json:\"time\""
							} "json:\"departure\""
							Arrival struct {
								Airport string    "json:\"airport\""
								Time    time.Time "json:\"time\""
							} "json:\"arrival\""
							DurationMinutes int "json:\"duration_minutes\""
							LayoverMinutes  int "json:\"layover_minutes,omitempty\""
						}{
							{
								FlightNumber: "GA315",
								Departure: struct {
									Airport string    "json:\"airport\""
									Time    time.Time "json:\"time\""
								}{
									Airport: "CGK",
									Time:    departurGA315,
								},
								Arrival: struct {
									Airport string    "json:\"airport\""
									Time    time.Time "json:\"time\""
								}{
									Airport: "SUB",
									Time:    arriveGA315,
								},
								DurationMinutes: 90,
							},
							{
								FlightNumber: "GA332",
								Departure: struct {
									Airport string    "json:\"airport\""
									Time    time.Time "json:\"time\""
								}{
									Airport: "SUB",
									Time:    departurGA332,
								},
								Arrival: struct {
									Airport string    "json:\"airport\""
									Time    time.Time "json:\"time\""
								}{
									Airport: "DPS",
									Time:    arriveGA332,
								},
								DurationMinutes: 90,
								LayoverMinutes:  105,
							},
						},
						AvailableSeats: 22,
						FareClass:      "economy",
						Baggage: struct {
							CarryOn int "json:\"carry_on\""
							Checked int "json:\"checked\""
						}{
							CarryOn: 1,
							Checked: 2,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := api.NewAPIRepository()
			got, err := repo.GetGarudaIndonesiaFlights(
				context.Background(),
				test.origin,
				test.destination,
				test.departureDate,
			)
			if err != nil {
				require.True(t, test.wantErr)
			} else {
				require.False(t, test.wantErr)
			}
			require.Equal(t, test.want, got)
		})
	}
}

func TestAPIRepository_GetLionAirFlights(t *testing.T) {
	tests := []struct {
		name          string
		origin        string
		destination   string
		departureDate string
		want          api.LionAirResponse
		wantErr       bool
	}{
		{
			name:          "succeed",
			origin:        "CGK",
			destination:   "DPS",
			departureDate: "2025-12-15",
			want: api.LionAirResponse{
				Success: true,
				Data: struct {
					AvailableFlights []struct {
						ID      string "json:\"id\""
						Carrier struct {
							Name string "json:\"name\""
							Iata string "json:\"iata\""
						} "json:\"carrier\""
						Route struct {
							From struct {
								Code string "json:\"code\""
								Name string "json:\"name\""
								City string "json:\"city\""
							} "json:\"from\""
							To struct {
								Code string "json:\"code\""
								Name string "json:\"name\""
								City string "json:\"city\""
							} "json:\"to\""
						} "json:\"route\""
						Schedule struct {
							Departure         string "json:\"departure\""
							DepartureTimezone string "json:\"departure_timezone\""
							Arrival           string "json:\"arrival\""
							ArrivalTimezone   string "json:\"arrival_timezone\""
						} "json:\"schedule\""
						FlightTime int  "json:\"flight_time\""
						IsDirect   bool "json:\"is_direct\""
						StopCount  int  "json:\"stop_count,omitempty\""
						Layovers   []struct {
							Airport         string "json:\"airport\""
							DurationMinutes int    "json:\"duration_minutes\""
						} "json:\"layovers,omitempty\""
						Pricing struct {
							Total    int    "json:\"total\""
							Currency string "json:\"currency\""
							FareType string "json:\"fare_type\""
						} "json:\"pricing\""
						SeatsLeft int    "json:\"seats_left\""
						PlaneType string "json:\"plane_type\""
						Services  struct {
							WifiAvailable    bool "json:\"wifi_available\""
							MealsIncluded    bool "json:\"meals_included\""
							BaggageAllowance struct {
								Cabin string "json:\"cabin\""
								Hold  string "json:\"hold\""
							} "json:\"baggage_allowance\""
						} "json:\"services\""
					} "json:\"available_flights\""
				}{
					AvailableFlights: []struct {
						ID      string "json:\"id\""
						Carrier struct {
							Name string "json:\"name\""
							Iata string "json:\"iata\""
						} "json:\"carrier\""
						Route struct {
							From struct {
								Code string "json:\"code\""
								Name string "json:\"name\""
								City string "json:\"city\""
							} "json:\"from\""
							To struct {
								Code string "json:\"code\""
								Name string "json:\"name\""
								City string "json:\"city\""
							} "json:\"to\""
						} "json:\"route\""
						Schedule struct {
							Departure         string "json:\"departure\""
							DepartureTimezone string "json:\"departure_timezone\""
							Arrival           string "json:\"arrival\""
							ArrivalTimezone   string "json:\"arrival_timezone\""
						} "json:\"schedule\""
						FlightTime int  "json:\"flight_time\""
						IsDirect   bool "json:\"is_direct\""
						StopCount  int  "json:\"stop_count,omitempty\""
						Layovers   []struct {
							Airport         string "json:\"airport\""
							DurationMinutes int    "json:\"duration_minutes\""
						} "json:\"layovers,omitempty\""
						Pricing struct {
							Total    int    "json:\"total\""
							Currency string "json:\"currency\""
							FareType string "json:\"fare_type\""
						} "json:\"pricing\""
						SeatsLeft int    "json:\"seats_left\""
						PlaneType string "json:\"plane_type\""
						Services  struct {
							WifiAvailable    bool "json:\"wifi_available\""
							MealsIncluded    bool "json:\"meals_included\""
							BaggageAllowance struct {
								Cabin string "json:\"cabin\""
								Hold  string "json:\"hold\""
							} "json:\"baggage_allowance\""
						} "json:\"services\""
					}{
						{
							ID: "JT740",
							Carrier: struct {
								Name string "json:\"name\""
								Iata string "json:\"iata\""
							}{
								Name: "Lion Air",
								Iata: "JT",
							},
							Route: struct {
								From struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								} "json:\"from\""
								To struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								} "json:\"to\""
							}{
								From: struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								}{
									Code: "CGK",
									Name: "Soekarno-Hatta International",
									City: "Jakarta",
								},
								To: struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								}{
									Code: "DPS",
									Name: "Ngurah Rai International",
									City: "Denpasar",
								},
							},
							Schedule: struct {
								Departure         string "json:\"departure\""
								DepartureTimezone string "json:\"departure_timezone\""
								Arrival           string "json:\"arrival\""
								ArrivalTimezone   string "json:\"arrival_timezone\""
							}{
								Departure:         "2025-12-15T05:30:00",
								DepartureTimezone: "Asia/Jakarta",
								Arrival:           "2025-12-15T08:15:00",
								ArrivalTimezone:   "Asia/Makassar",
							},
							FlightTime: 105,
							IsDirect:   true,
							Pricing: struct {
								Total    int    "json:\"total\""
								Currency string "json:\"currency\""
								FareType string "json:\"fare_type\""
							}{
								Total:    950000,
								Currency: "IDR",
								FareType: "ECONOMY",
							},
							SeatsLeft: 45,
							PlaneType: "Boeing 737-900ER",
							Services: struct {
								WifiAvailable    bool "json:\"wifi_available\""
								MealsIncluded    bool "json:\"meals_included\""
								BaggageAllowance struct {
									Cabin string "json:\"cabin\""
									Hold  string "json:\"hold\""
								} "json:\"baggage_allowance\""
							}{
								WifiAvailable: false,
								MealsIncluded: false,
								BaggageAllowance: struct {
									Cabin string "json:\"cabin\""
									Hold  string "json:\"hold\""
								}{
									Cabin: "7 kg",
									Hold:  "20 kg",
								},
							},
						},
						{
							ID: "JT742",
							Carrier: struct {
								Name string "json:\"name\""
								Iata string "json:\"iata\""
							}{
								Name: "Lion Air",
								Iata: "JT",
							},
							Route: struct {
								From struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								} "json:\"from\""
								To struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								} "json:\"to\""
							}{
								From: struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								}{
									Code: "CGK",
									Name: "Soekarno-Hatta International",
									City: "Jakarta",
								},
								To: struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								}{
									Code: "DPS",
									Name: "Ngurah Rai International",
									City: "Denpasar",
								},
							},
							Schedule: struct {
								Departure         string "json:\"departure\""
								DepartureTimezone string "json:\"departure_timezone\""
								Arrival           string "json:\"arrival\""
								ArrivalTimezone   string "json:\"arrival_timezone\""
							}{
								Departure:         "2025-12-15T11:45:00",
								DepartureTimezone: "Asia/Jakarta",
								Arrival:           "2025-12-15T14:35:00",
								ArrivalTimezone:   "Asia/Makassar",
							},
							FlightTime: 110,
							IsDirect:   true,
							Pricing: struct {
								Total    int    "json:\"total\""
								Currency string "json:\"currency\""
								FareType string "json:\"fare_type\""
							}{
								Total:    890000,
								Currency: "IDR",
								FareType: "ECONOMY",
							},
							SeatsLeft: 38,
							PlaneType: "Boeing 737-800",
							Services: struct {
								WifiAvailable    bool "json:\"wifi_available\""
								MealsIncluded    bool "json:\"meals_included\""
								BaggageAllowance struct {
									Cabin string "json:\"cabin\""
									Hold  string "json:\"hold\""
								} "json:\"baggage_allowance\""
							}{
								WifiAvailable: false,
								MealsIncluded: false,
								BaggageAllowance: struct {
									Cabin string "json:\"cabin\""
									Hold  string "json:\"hold\""
								}{
									Cabin: "7 kg",
									Hold:  "20 kg",
								},
							},
						},
						{
							ID: "JT650",
							Carrier: struct {
								Name string "json:\"name\""
								Iata string "json:\"iata\""
							}{
								Name: "Lion Air",
								Iata: "JT",
							},
							Route: struct {
								From struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								} "json:\"from\""
								To struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								} "json:\"to\""
							}{
								From: struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								}{
									Code: "CGK",
									Name: "Soekarno-Hatta International",
									City: "Jakarta",
								},
								To: struct {
									Code string "json:\"code\""
									Name string "json:\"name\""
									City string "json:\"city\""
								}{
									Code: "DPS",
									Name: "Ngurah Rai International",
									City: "Denpasar",
								},
							},
							Schedule: struct {
								Departure         string "json:\"departure\""
								DepartureTimezone string "json:\"departure_timezone\""
								Arrival           string "json:\"arrival\""
								ArrivalTimezone   string "json:\"arrival_timezone\""
							}{
								Departure:         "2025-12-15T16:20:00",
								DepartureTimezone: "Asia/Jakarta",
								Arrival:           "2025-12-15T21:10:00",
								ArrivalTimezone:   "Asia/Makassar",
							},
							FlightTime: 230,
							IsDirect:   false,
							StopCount:  1,
							Layovers: []struct {
								Airport         string "json:\"airport\""
								DurationMinutes int    "json:\"duration_minutes\""
							}{
								{
									Airport:         "SUB",
									DurationMinutes: 75,
								},
							},
							Pricing: struct {
								Total    int    "json:\"total\""
								Currency string "json:\"currency\""
								FareType string "json:\"fare_type\""
							}{
								Total:    780000,
								Currency: "IDR",
								FareType: "ECONOMY",
							},
							SeatsLeft: 52,
							PlaneType: "Boeing 737-800",
							Services: struct {
								WifiAvailable    bool "json:\"wifi_available\""
								MealsIncluded    bool "json:\"meals_included\""
								BaggageAllowance struct {
									Cabin string "json:\"cabin\""
									Hold  string "json:\"hold\""
								} "json:\"baggage_allowance\""
							}{
								WifiAvailable: false,
								MealsIncluded: false,
								BaggageAllowance: struct {
									Cabin string "json:\"cabin\""
									Hold  string "json:\"hold\""
								}{
									Cabin: "7 kg",
									Hold:  "20 kg",
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := api.NewAPIRepository()
			got, err := repo.GetLionAirFlights(
				context.Background(),
				test.origin,
				test.destination,
				test.departureDate,
			)
			if err != nil {
				require.True(t, test.wantErr)
			} else {
				require.False(t, test.wantErr)
			}
			require.Equal(t, test.want, got)
		})
	}
}
