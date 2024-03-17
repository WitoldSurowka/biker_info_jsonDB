package main

import (
	"time"
)

func main() {

	currentTime := time.Now()

	ticker := time.NewTicker(1 * time.Second)

	ReceiveSMS("123123123", "Rodo")
	ReceiveSMS("123123123", "help")
	ReceiveSMS("123123123", "miasta")
	ReceiveSMS("123123123", "Ropica")
	ReceiveSMS("123123123", "Ropica")
	ReceiveSMS("123123123", "Ropica")
	ReceiveSMS("123123123", "Kraków")
	ReceiveSMS("123123124", "Ropica")

	for range ticker.C {
		currentTime = time.Now()
		//fmt.Println(currentTime)
		//if currentTime.Hour() == 20 && currentTime.Minute() == 22 && currentTime.Second() == 50 {
		//	wc := NewWeatherConditions(WeatherFetcher("Ropica Górna"))
		//	fmt.Println(wc.WeatherConditionMessage())
		//}
		if currentTime.Second() == 50 {
			SendFeed()
		}

	}

}
