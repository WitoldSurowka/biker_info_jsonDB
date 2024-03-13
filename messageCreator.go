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
	City    string
}

// NewWeatherConditions constructor for WeatherConditions structure.
func NewWeatherConditions(precip float64, tempMin int, wind float64, city string) *WeatherConditions {
	return &WeatherConditions{
		Precip:  precip,
		TempMin: tempMin,
		Wind:    wind,
		City:    city,
	}
}

// CheckConditions method checks for exceeded values and composes a message.
func (wc *WeatherConditions) WeatherConditionMessage() string {
	var message, cityString string
	cityString = "[" + wc.City + "]"

	if wc.Precip > 1.9 && wc.Precip < 3 {
		message += fmt.Sprintf("Mrzawka: %.2f mm.\n", wc.Precip)
	}
	if wc.Precip >= 3 && wc.Precip < 10 {
		message += fmt.Sprintf("Deszcz: %.2f mm.\n", wc.Precip)
	}
	if wc.Precip >= 10 {
		message += fmt.Sprintf("Ulewa: %.2f mm.\n", wc.Precip)
	}
	if wc.TempMin < 10 && wc.TempMin > 2 {
		message += fmt.Sprintf("Zimno - temperatura minimalna: %v °C.\n", wc.TempMin)
	}
	if wc.TempMin <= 2 {
		message += fmt.Sprintf("Zimno i możliwe oblodzenie - temperatura minimalna: %v °C.\n", wc.TempMin)
	}
	if wc.Wind >= 11 {
		message += fmt.Sprintf("Wyjątkowo silny wiatr: %.2f km/h.\n", wc.Wind*3.6)
	}
	if len(message) > 0 {
		message = fmt.Sprintln("Uwaga"+cityString+", jutro (", m[fmt.Sprint(time.Now().Weekday()+1)], ") niekorzystne warunki dla jednośladów:") + message
	}
	if wc.Wind <= 8 && wc.TempMin >= 18 && wc.Precip == 0 {
		message = fmt.Sprintf("Jutro "+cityString+" idealne warunki pogodowe dla jednośladów: \n"+
			"wiatr:%.2f km/h,\n"+
			"temperatura minimalna:%v °C,\n"+
			"brak opadów.\n", wc.Wind*3.6, wc.TempMin)
	}
	if len(message) == 0 {
		message = "Jutro odpowiednie warunki do jazdy jednośladem \\m/" + cityString
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
