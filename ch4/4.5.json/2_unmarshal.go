package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"release"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("marshaling of movies failed : %s \n", err)
	}

	fmt.Printf("%s \n", data)
	//[{"Title":"Casablanca","release":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","release":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","release":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]

	var moviesStruct []struct {
		Title  string
		Year   int  `json:"release"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	if err := json.Unmarshal(data, &moviesStruct); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s \n ", err)
	}
	fmt.Printf("%+v\n", moviesStruct)
	//[{Title:Casablanca Year:1942 Color:false Actors:[Humphrey Bogart Ingrid Bergman]} {Title:Cool Hand Luke Year:1967 Color:true Actors:[Paul Newman]} {Title:Bullitt Year:1968 Color:true Actors:[Steve McQueen Jacqueline Bisset]}]

}
