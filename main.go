package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

// initializing a data structure to keep the scraped data
type WeatherStatus struct {
	url, precip string
}

func main() {
	var status WeatherStatus
	c := colly.NewCollector()
	shouldStop := false

	c.OnHTML(".daily-weather-list-item", func(e *colly.HTMLElement) {
		if shouldStop {
			return
		}

		status.url = e.ChildAttr("a", "href")
		status.precip = e.ChildText(".Precipitation-module__main-sU6qN[data-color=true]")

		fmt.Println(status)
		//c.OnHTML scrape in a loop, so after the desired data is fetched, we do not process data no more
		if strings.Contains(status.url, "i=1") {
			shouldStop = true
		}
	})

	c.Visit("https://www.yr.no/nb/v%C3%A6rvarsel/daglig-tabell/2-3094802/Polen/Ma%C5%82opolskie/Krak%C3%B3w/Krak%C3%B3w")
	c.Wait() // Wait until scraping is complete

	precipStringLong := status.precip[:len(status.precip)-2]
	precipStringShort := precipStringLong[8:]
	precipStringShort = strings.Replace(precipStringShort, ",", ".", 1)
	precip, err := strconv.ParseFloat(precipStringShort, 32)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}

	fmt.Println(precip)

}
