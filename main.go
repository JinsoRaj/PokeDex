package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"html/template"
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
	fmt.Println(res.Pokemon[0])
}

func apiTest(w http.ResponseWriter, r *http.Request){
	
	d,_ := ioutil.ReadFile("pok2.json");

    rawMsg := json.RawMessage(string(d))
    var objmap map[string]*json.RawMessage
    err := json.Unmarshal(rawMsg, &objmap)
    if err != nil {
      fmt.Println(err)
    }
    //fmt.Println(objmap)
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(objmap["pokemon"])

    //fmt.Fprintf(w, "show api")
    //fmt.Println("apiPage used")
    //fmt.Fprintf(w, "/test api")
    //fmt.Println("testApiPage used")
}

func ShowHome(w http.ResponseWriter, r *http.Request) {
	resp := getAllPokemons()
	t, err := template.ParseFiles("templates/index.html")
	errorCheck(err)
	t.Execute(w, resp)
	//http.ServeFile(w, r, "templates/index.html")
}

//handling all endpoints
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/home", ShowHome)
	http.HandleFunc("/api", apiPage)
	http.HandleFunc("/getTableData", apiTest)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

//main
func main() {
    handleRequests()
}