package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie1 struct {
	title  string
	year   int  `json:"released"`
	color  bool `json:"color,omitempty"`
	actors []string
}

type Movie2 struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies1 := []Movie1{
		{title: "Casablanca", year: 1942, color: false, actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{title: "Cool Hand Luke", year: 1967, color: true, actors: []string{"Paul Newman"}},
		{title: "Bullitt", year: 1968, color: true, actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	movies2 := []Movie2{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data1, err := json.Marshal(movies1)
	if err != nil {
		log.Fatalf("Json marshaling failed : %s", err)
	}
	fmt.Printf("%s \n", data1) //	[{},{},{}] -> we should use upperCase word in field off struct

	data2, err := json.Marshal(movies2)
	if err != nil {
		log.Fatalf("Json marshaling failed : %s", err)
	}
	fmt.Printf("%s \n", data2)

	//[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]

	data3, err := json.MarshalIndent(movies2, "", "    ")
	if err != nil {
		log.Fatalf("Json marshaling failed : %s", err)
	}
	fmt.Printf("%s \n", data3)
	//[
	//	{
	//	"Title": "Casablanca",
	//	"released": 1942,
	//	"Actors": [
	//	"Humphrey Bogart",
	//	"Ingrid Bergman"
	//	]
	//	},
	//	{
	//	"Title": "Cool Hand Luke",
	//	"released": 1967,
	//	"color": true,
	//	"Actors": [
	//	"Paul Newman"
	//	]
	//	},
	//	{
	//	"Title": "Bullitt",
	//	"released": 1968,
	//	"color": true,
	//	"Actors": [
	//	"Steve McQueen",
	//	"Jacqueline Bisset"
	//	]
	//	}
	//]

}
