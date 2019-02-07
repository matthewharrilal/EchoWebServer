package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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


// func lifeInfo(context echo.Context) error {
// 	lifeInfo := LifeInfo{} // So you can pass this object by reference as opposed to value

// 	err := json.NewDecoder(context.Request().Body).Decode(&lifeInfo) // The body of data that the user is trying to send through the request
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Unmarshaled data >>>> %s", lifeInfo)
// 	return context.String(http.StatusOK, "Accepted Request")
}

// func dogs(context echo.Context) error {
// 	dog := Dog{}

// 	// Now that we have our dog object instantiated
// 	// We have to intercept the request body and we have to unmarshal it into our dog object
// 	err := json.NewDecoder(context.Request().Body).Decode(&dog) // Unmarshal data and write it to our dog object in memory
// 	// Why does this function return an errror ... design decision?

// 	fmt.Printf("JSON from decoded data %s and our dog object %s", err, dog)
// 	return context.String(http.StatusOK, "Accepted Request")
}

// func cats(context echo.Context) error {
// 	// This is another way of quite literally binding the json to our instantiated object
// 	cat := Cat{}

// 	err := context.Bind(&cat) // Unmarshal data and bind it immediately by writing it to our cat object
// 	if err != nil {
// 		fmt.Printf("Error when binding data to our cat object %s", err)
// 	}

// 	fmt.Printf("This is the cat object %s", cat)
// 	return context.String(http.StatusOK, "Accepted Request")
// }

func main() {
	fmt.Println("Server is running on port 4000!")

	// Instantiate the client that is going to be trigerring the requests
	client := echo.New() // Creating an instance of our server
	// client.Use(middleware.Logger())
	// client.GET("/greeting", greeting) // The framework provides us the context?
	// client.POST("/lifeInfo", lifeInfo)

	// // You can attach middlewares to a route

	// client.POST("/dogs", dogs)
	// client.POST("/cats", cats)
	// client.GET("/goodbye", goodbye)

	client.Start(":4000")

}
