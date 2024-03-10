package main

import (
	"fmt"
	"time"
)

//Make weekdays polish

var m = map[string]string{
	"Monday":    "Poniedziałek",
	"Tuesday":   "Wtorek",
	"Wednesday": "Środa",
	"Thursday":  "Czwartek",
	"Friday":    "Piątek",
	"Saturday":  "Sobota",
	"Sunday":    "Niedziela",
}

// WeatherConditions structure holds weather variables and their threshold values.
type WeatherConditions struct {
	Precip  float64 // Precipitation in mm
	TempMin int     // Minimum temperature in °C
	Wind    float64 // Wind speed in km/h
}

// NewWeatherConditions constructor for WeatherConditions structure.
func NewWeatherConditions(precip float64, tempMin int, wind float64) *WeatherConditions {
	return &WeatherConditions{
		Precip:  precip,
		TempMin: tempMin,
		Wind:    wind,
	}
}

// CheckConditions method checks for exceeded values and composes a message.
func (wc *WeatherConditions) WeatherConditionMessage() string {
	var message string

	if wc.Precip > 2.2 && wc.Precip < 5 {
		message += fmt.Sprintf("Lekki deszcz: %.2f mm.\n", wc.Precip)
	}
	if wc.Precip >= 5 {
		message += fmt.Sprintf("Znaczny deszcz: %.2f mm.\n", wc.Precip)
	}
	if wc.TempMin < 10 && wc.TempMin > 2 {
		message += fmt.Sprintf("Zimno - temperatura minimalna: %v °C.\n", wc.TempMin)
	}
	if wc.TempMin <= 2 {
		message += fmt.Sprintf("Możliwe oblodzenie - temperatura minimalna: %v °C.\n", wc.TempMin)
	}
	if wc.Wind >= 14 {
		message += fmt.Sprintf("Wyjątkowo silny wiatr: %.2f km/h.\n", wc.Wind*3.6)
	}
	if len(message) == 0 {
		message = "Jutro odpowiednie warunki do jazdy jednośladem \\m/"
	} else {
		message = fmt.Sprintln("Uwaga, jutro (", m[fmt.Sprint(time.Now().Weekday()+1)], ") niekorzystne warunki dla jednośladów:") + message
	}

	return message
}

//func main() {
//	//Example values and thresholds
//	wc := NewWeatherConditions(3, 3, 20.0)
//
//	//Check conditions and print message
//	fmt.Println(wc.CheckConditions())
//}
