package main

import (
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"
	"encoding/json"
)


// All Pages 

func HomePage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "templates/home.html")
	fmt.Println("/root  homePage used")
}

//for json load (test purpose)
func ApiTest(w http.ResponseWriter, r *http.Request){
	d,_ := ioutil.ReadFile("pokedex.json");
	rawMsg := json.RawMessage(string(d))
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(rawMsg, &objmap)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(objmap["pokemon"])
	fmt.Println("Full api loaded @ /test")
}

//Level A - get data and show in HTML table
func GetTableData(w http.ResponseWriter, r *http.Request) {
	resp := getAllPokemons()
	t, err := template.ParseFiles("templates/table.html")
	errorCheck(err)
	t.Execute(w, resp)
	fmt.Println("Table data loaded @ /getTableData")

}

//Level B - get data in new struct and display each pokemon info
func InfoPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inf := getinfoAllPokemons()
	inp := r.Form["name"][0]
	for _, v := range inf.PokemonInfo{
		if v.Name == inp{
			data := v
			t, err := template.ParseFiles("templates/pokeinfo.html")
			errorCheck(err)
			t.Execute(w, data)
		}
	}    
}