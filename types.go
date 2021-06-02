package main

// type Scraper struct {
// 	url     string  `json:"url"`
// 	product Product `json:"product"`
// }

type Response struct {
	URL     		string 	`json:"url"`
	ProductName     string 	`json:"productname"`
	Price        	string 	`json:"price"`
	Reviews  		string 	`json:"reviews"`
	Description		string	`json:"description"`
}
