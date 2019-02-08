package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type Region struct {
	gorm.Model

	Descriptions []Description `json:"descriptions"` // Slice of description objects

	Identification int `json:"id"`

	IsMainSeries bool `json:"bool"`

	Name string `json:"name"`

	Names []Language `json:"names"`

	PokemonEntries []PokemonSpecies `json:"pokemon_entries"`
}

type Description struct {
	Description string `json:"description"`

	Language Language `json:"language"`
}

type PokemonSpecies struct {
	EntryNumber int `json:"entry_number"`

	PokemonSpecies Language `json:"pokemon_species"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func ConfigureDatabase() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	defer db.Close() // Close database connection after surrounding function executes
}

func ObtainRegion(channel chan int) Region {
	// Pass the region that is being returned through the channel that the main go routine is going to receive

	url := "http://pokeapi.co/api/v2/pokedex/kanto/"

	// Allows you to add added configurations to the http request and execute you can construct the request and execute with the client which add specified configurations
	client := http.Client{
		// If for some reason the request to the pokedex lags it will take awhile timeout allowed for request ... no more than a minute
		Timeout: time.Second * 60,
	}

	// To construct a new request HTTP Method, desired url, and nil?
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		// Checking if error present
		log.Fatal(err)
	}

	// Execute request and assign callback status
	res, getErr := client.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	region := Region{}

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(res.Body).Decode(&region)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("REGION >>> ", region)
	return region
}

func main() {
	fmt.Println("Server is running on port 4000!")

	// Instantiate the client that is going to be trigerring the requests
	client := echo.New() // Creating an instance of our server
	// client.GET("/region", ObtainRegion())
	go ObtainRegion()

	// Now that you have the region
	client.Start(":4000")

}
