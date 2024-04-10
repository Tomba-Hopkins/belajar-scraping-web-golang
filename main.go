package main

import (
	"fmt"

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

	c.OnHTML(".LyricsHeader__Container-ejidji-1", func(e *colly.HTMLElement){

		title := e.ChildText("h2.TextLabel-sc-8kw9oj-0.LyricsHeader__Title-ejidji-0")
		lyric := e.ChildText("div.Lyrics__Container-sc-1ynbvzw-1.kUgSbL")
		
		fmt.Printf("Title: %v\nLyrics: %v\n", title, lyric)
	})
	
	c.Visit("https://genius.com/Slipknot-psychosocial-2012-remaster-lyrics")




}