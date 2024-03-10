package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

// initializing a data structure to keep the scraped data
type WeatherStatus struct {
	url, precip string
}

func main() {
	var status WeatherStatus
	c := colly.NewCollector()

	c.OnHTML(".daily-list-item", func(e *colly.HTMLElement) {

		status.url = e.ChildAttr("a", "href")
		status.precip = e.ChildText(".precip")

		fmt.Println(status)
		if strings.Contains(status.url, "tomorrow") {
			return
		}
	})

	c.Visit("https://www.accuweather.com/pl/pl/krakow/274455/weather-forecast/274455")
	c.Wait() // Wait until scraping is complete

	file, err := os.Create("weather.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{"url", "precipitation"}
	if err := writer.Write(headers); err != nil {
		log.Fatalln("Failed to write headers to CSV file", err)
	}

	record := []string{status.url, status.precip}

	if err := writer.Write(record); err != nil {
		log.Fatalln("Failed to write record to CSV file", err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatalln("Error occurred during flushing to CSV file", err)
	}
}
