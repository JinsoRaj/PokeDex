package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

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

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getAllPokemons() (ap *AllPokemons) {
	file, err := os.OpenFile("pok2.json", os.O_RDWR|os.O_APPEND, 0666)
	errorCheck(err)
	b, err := ioutil.ReadAll(file)
	var allPkms AllPokemons
	json.Unmarshal(b, &allPkms)
	errorCheck(err)
	return &allPkms 
}


// Pages (home, api , apitest)
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("homePage used")
}

func apiPage(w http.ResponseWriter, r *http.Request){
	res := getAllPokemons()
	//fmt.Fprintf(w, res)
	fmt.Println(res.Pokemon[0].Height)
}

func apiTest(w http.ResponseWriter, r *http.Request){
	
    fmt.Fprintf(w, "/test api")
    //fmt.Println("testApiPage used")
}

//handling all endpoints
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api", apiPage)
	http.HandleFunc("/test", apiTest)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

//main
func main() {
    handleRequests()
}