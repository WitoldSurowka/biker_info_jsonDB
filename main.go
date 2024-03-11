package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)

	ticker := time.NewTicker(1 * time.Second)
	//defer ticker.Stop()

	for range ticker.C {
		currentTime = time.Now()
		fmt.Println(currentTime)
		// Check if current time is 20:00:00
		if currentTime.Hour() == 23 && currentTime.Minute() == 10 && currentTime.Second() == 10 {
			wc := NewWeatherConditions(WeatherFetcher())
			fmt.Println(wc.WeatherConditionMessage())
		}

	}

}
