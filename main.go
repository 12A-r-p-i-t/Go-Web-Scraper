package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	var allowedDomains []string
	c.AllowedDomains = allowedDomains

	c.OnRequest(func(req *colly.Request) {
		fmt.Println("Heyy There")
	})
	err := c.Visit("https://hackerspaces.org/")
	if err != nil {
		log.Fatal("Error in visiting the given URL :", err)
	}
}
