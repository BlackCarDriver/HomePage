package main


type HomePageNews struct{
	Img_url	string 			`json:"imgurl"`
	Link_url string 		`json:"linkurl"`
	Title	string 			`json:"title"`
}

type ArticlaBox struct{
	Img_url	string 			`json:"imgurl"`
	Link_url string 		`json:"linkurl"`
	Brif	string 			`json:"brief"`
	Date	string 			`json:"date"`
}
