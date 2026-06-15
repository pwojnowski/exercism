// Package weather provides tools to forecast weather
// for given city and condition.
package weather

var (
    // CurrentCondition provides current condition.
	CurrentCondition string
    // CurrentLocation provides current location.
	CurrentLocation  string
)

// Forecast calculates forecast for given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
