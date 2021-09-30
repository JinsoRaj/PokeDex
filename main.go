// Jinso Raj
//Test 43 - PokeDex Full stack 
// main.go and handlers.go must run  together
// run command:   go run .


package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)


//struct json out (full data)
type AllPokemons struct {
	Pokemon []struct {
	  ID            int       `json:"id"`
	  Num           string    `json:"num"`
	  Name          string    `json:"name"`
	  Img           string    `json:"img"`
	  Type          []string  `json:"type"`
	  Height        string    `json:"height"`
	  Weight        string    `json:"weight"`
	  Candy         string    `json:"candy"`
	  CandyCount    int       `json:"candy_count,omitempty"`
	  Egg           string    `json:"egg"`
	  SpawnChance   float64   `json:"spawn_chance"`
	  AvgSpawns     int       `json:"avg_spawns"`
	  SpawnTime     string    `json:"spawn_time"`
	  Multipliers   []float64 `json:"multipliers"`
	  Weaknesses    []string  `json:"weaknesses"`
	  NextEvolution []struct {
		Num  string `json:"num"`
		Name string `json:"name"`
	  } `json:"next_evolution,omitempty"`
	  PrevEvolution []struct {
		Num  string `json:"num"`
		Name string `json:"name"`
	  } `json:"prev_evolution,omitempty"`
	} `json:"pokemon"`
  }

  // json struct (for search by name page. image not mentioned. but needed for UI)
  type AllPokemonsInfo struct {
	PokemonInfo []struct {
	  Num           string    `json:"num"`
	  Name          string    `json:"name"`
	  Img           string    `json:"img"`
	  Type          []string  `json:"type"`
	  Height        string    `json:"height"`
	  Weight        string    `json:"weight"`
	  Weaknesses    []string  `json:"weaknesses"`
	  NextEvolution []struct {
		Num  string `json:"num"`
		Name string `json:"name"`
	  } `json:"next_evolution,omitempty"`
	  PrevEvolution []struct {
		Num  string `json:"num"`
		Name string `json:"name"`
	  } `json:"prev_evolution,omitempty"`
	} `json:"pokemon"`
  }


// error checking fn
func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// full data return - unmarshaling process
func getAllPokemons() (ap *AllPokemons) {
	file, err := os.OpenFile("pokedex.json", os.O_RDONLY, 0444)
	errorCheck(err)
	b, err := ioutil.ReadAll(file)
	var allPkms AllPokemons
	json.Unmarshal(b, &allPkms)
	errorCheck(err)
	return &allPkms 
}
// get pokemon info for specific  name match
func getinfoAllPokemons() (ap *AllPokemonsInfo) {
	file, err := os.OpenFile("pokedex.json", os.O_RDONLY, 0444)
	errorCheck(err)
	b, err := ioutil.ReadAll(file)
	var allPkmsIn AllPokemonsInfo
	json.Unmarshal(b, &allPkmsIn)
	errorCheck(err)
	return &allPkmsIn 
}

//handling all endpoints
func handleRequests() {
	http.HandleFunc("/", HomePage) // home UI
	//fix style paths in html
	http.Handle("/styles/", http.FileServer(http.Dir("./templates")))
	http.HandleFunc("/getTableData", GetTableData) // Level A
	http.HandleFunc("/test", ApiTest) // for test purpose
	http.HandleFunc("/info", InfoPage) // Level B
	log.Fatal(http.ListenAndServe(":3000", nil))
}


//main - calling all handlers
func main() {
	handleRequests()
}


//app run command:   go run .