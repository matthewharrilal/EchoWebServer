package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type LifeInfo struct {
	Happy string `json:"happy"`

	Name string `json:"name"`
	
}

type Dog struct {
	Name string `json:"name"`

	Age int `json:"age"`

	isGoodBoi bool `json:"isGoodBoi"`
}

type Cat struct {
	Name string `json:"name"`
}

func greeting(context echo.Context) error {
	// Context holds relevant information about the request such as data response body ...
	return context.String(http.StatusOK, "Thanks for visiting the server!")
}

func goodbye(context echo.Context) error {
	nameParam := context.QueryParam("name")

	// From my understanding the difference between params and query params is that query params usually have more information about the type of path that the user wants to be directed to
	// /goodbye/:name and to extract it context.Param["name"]

	// That is one way of extracting the parameters you can also add it in the route by denoting the : symbol along with the name of the resource

	// jsonObject := map[string]string{
	// 	"name": nameParam,
	// }
	// name := context.JSON(http.StatusOK, jsonObject)
	// return context.String(http.StatusOK, fmt.Sprintf("Are you sure you want to leave %s?", nameParam))
	return context.JSON(http.StatusOK, map[string]string{
		"message": nameParam,
	})
}

func lifeInfo(context echo.Context) error {
	lifeInfo := LifeInfo{} // So you can pass this object by reference as opposed to value
	
	err := json.NewDecoder(context.Request().Body).Decode(&lifeInfo) // The body of data that the user is trying to send through the request
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unmarshaled data >>>> %s", lifeInfo)
	return context.String(http.StatusOK, "Accepted Request")
}

func dogs(context echo.Context) error {
	dog := Dog{}

	// Now that we have our dog object instantiated
	// We have to intercept the request body and we have to unmarshal it into our dog object
	err := json.NewDecoder(context.Request().Body).Decode(&dog) // Unmarshal data and write it to our dog object in memory
	// Why does this function return an errror ... design decision?

	fmt.Printf("JSON from decoded data %s and our dog object %s", err, dog)
	return context.String(http.StatusOK, "Accepted Request")
}

func cats(context echo.Context) error {
	// This is another way of quite literally binding the json to our instantiated object
	cat := Cat{}

	err := context.Bind(&cat) // Unmarshal data and bind it immediately by writing it to our cat object
	if err != nil {
		fmt.Printf("Error when binding data to our cat object %s", err)
	}

	fmt.Printf("This is the cat object %s", cat)
	return context.String(http.StatusOK, "Accepted Request")
}

func main() {
	fmt.Println("Server is running on port 4000!")

	// Instantiate the client that is going to be trigerring the requests
	client := echo.New() // Creating an instance of our server
	client.Use(middleware.Logger())
	client.GET("/greeting", greeting) // The framework provides us the context?
	client.POST("/lifeInfo", lifeInfo)

	// You can attach middlewares to a route
	

	client.POST("/dogs", dogs)
	client.POST("/cats", cats)
	client.GET("/goodbye", goodbye)

	client.Start(":4000")

}
