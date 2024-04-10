package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()


	// belajar coli eh colly
	// c.OnRequest(func(r *colly.Request){
	// 	fmt.Println("Berkunjung ke: ", r.URL)
	// })

	// c.OnError(func(_ *colly.Response, err error){
	// 	log.Println("Ada error cuy: ", err)
	// })

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Tervisit kan: ", r.Request.URL)
	// })

	// c.OnHTML("a", func(e *colly.HTMLElement){
	// 	fmt.Printf("%v\n", e.Attr("href"))
	// })

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println(r.Request.URL, "Ter Scrap kan")
	// })

	// c.Visit("https://tomba-hopkins.github.io/portfolio-tailwindcss/")



	// Praktek

	c.OnHTML(".box", func(e *colly.HTMLElement){
		title := e.ChildText(".title")

		lyrics := []string{}

		e.ForEach(".lyric", func(_ int,el *colly.HTMLElement){
			lyric := strings.TrimSpace(el.Text)
			lyrics = append(lyrics, lyric)
		})	

		fmt.Printf("Title: %v\n\n", title)

		for _, r := range lyrics{
			fmt.Printf("%v\n\n", r)
		}
	})
	
	c.Visit("https://tomba-hopkins.github.io/belajar-scraping-web-golang/views/")




}