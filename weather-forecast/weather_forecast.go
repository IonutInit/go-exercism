// Package weather contains a function returning a string with the current location and current condition in a friendly format.
package weather

// CurrentCondition is a string taking in the current condition.
var CurrentCondition string
// CurrentLocation is a string taking in the current location.
var CurrentLocation string

// Forecast is a function that returns the weather condition in the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
