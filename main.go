package main

import (
	"fmt"
	"net/http"
	"log"
	hand "github.com/MajorNick/picnic-task/handlers"
)


func main(){
	
	fileserver := http.FileServer(http.Dir("./htmls"))
	http.Handle("/htmls/",http.StripPrefix("/htmls",fileserver))
	http.HandleFunc("/",hand.StartHandler)
	http.HandleFunc("/rawdata.html",hand.Rawdata)
	http.HandleFunc("/piechart.html",hand.PieChart)
	http.HandleFunc("/rawdata/filter.html",hand.GetValues)
	
	fmt.Println("Starting Server:")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error in Server Handling: %v", err)
	}
	
}