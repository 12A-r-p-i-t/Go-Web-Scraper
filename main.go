package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Courses struct {
	Img     string
	College string
	Degree  string
}

var CourseraCourses []Courses

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/6.0"
	c.SetRequestTimeout(120 * time.Second)

	c.OnHTML("#rendered-content > div > div > div > div.browse-children-wrapper > section > div.product-offerings-wrapper > section > div > div.css-191v487 > section:nth-child(1) > div > div.slick-slider.slick-initialized > div > div", func(e *colly.HTMLElement) {
		e.ForEach(".slick-slide", func(i int, h *colly.HTMLElement) {
			course := Courses{}
			course.Img = h.ChildAttrs("img", "src")[0]
			course.College = h.ChildText("span")
			course.Degree = h.ChildText("h2")
			CourseraCourses = append(CourseraCourses, course)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		bytes, err := json.Marshal(CourseraCourses)
		if err != nil {
			log.Fatal("Error in marshalling data")
		}
		err = os.WriteFile("output.json", bytes, 0664)
		if err != nil {
			log.Fatal("Error in writing to file")
		}
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.coursera.org/browse/data-science/machine-learning")
}
