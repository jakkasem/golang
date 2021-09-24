package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Info main log")
	handleRequests()
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Info home")
	fmt.Fprint(w, "<h1>Home</h1>")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/books", getAllBooks)

	log.Println(http.ListenAndServe(":8080", nil))

}

type book struct {
	Isbn       string `json:"isbn"`
	Name       string
	AuthorName string
	Year       string
	TotalPage  int
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Info getallbooks")

	allBooks := []book{
		book{
			Isbn:       "978-616-08-2539-4",
			Name:       "พัฒนาเว็บแอปพลิเคชั่นด้วย JavaScript",
			AuthorName: "จตุรพัชร์ พัฒนทรงศิวิไล",
			Year:       "2529",
			TotalPage:  460,
		},
		book{
			Isbn:       "978-616-93108-0-8",
			Name:       "BIG DATA SERIES I, Introduction to a Big Data Project",
			AuthorName: "ดร.อสมา กุลวานิชไชยนันท์",
			Year:       "2561",
			TotalPage:  247,
		},
	}

	json.NewEncoder(w).Encode(allBooks)

}
