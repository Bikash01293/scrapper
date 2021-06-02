package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("ping"))
}

func main() {
	addr := ":7171"

	http.HandleFunc("/ping", ping)

	r := mux.NewRouter()

	r.HandleFunc("/search", getData).Methods("POST")

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//Verify the param "URL" exists
	// URL := r.URL.Query().Get("url")

	// var scraper Scraper
	URL := r.Header.Get("url")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", URL)

	//Create a new collector which will be in charge of collect the data from HTML
	c := colly.NewCollector()

	//Slices to store the data
	var response []string

	//onHTML function allows the collector to use a callback function when the specific HTML tag is reached
	//in this case whenever our collector finds an
	//anchor tag with href it will call the anonymous function
	// specified below which will get the info from the href and append it to our slice
	c.OnHTML("div#dp", func(e *colly.HTMLElement) {

		url := e.Request.URL.String()
		productName := e.ChildText("span.a-size-large.product-title-word-break")
		price := e.ChildText(".a-spacing-none.a-text-left.a-size-mini.twisterSwatchPrice")
		description := e.ChildText("div#productDescription")
		stars := e.ChildText("span#acrCustomerReviewText")

		if url == "" || productName == "" && price == "" || description == "" || stars == "" {
			// If we can't get any url, product name, price, reviews we return and go directly to the next element
			return
		}

		response = append(response, url, productName, price, description, stars)
	})

	//Command to visit the website
	c.Visit(URL)

	// response.Header().Set("Content-Type", "application/json")
	// response.WriteHeader(http.StatusCreated)
	res := make(map[string]interface{})
	res["url"] = response[0]
	res["productName"] = response[1]
	res["price"] = response[2]
	res["description"] = response[3]
	res["stars"] = response[4]

	// des, _ := json.Marshal(res)
	// response.Write([]byte(des))

	// parse our response slice into JSON format
	b, err := json.Marshal(res)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}
	Caller()
	// Add some header and write the body for our endpoint
	// w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
