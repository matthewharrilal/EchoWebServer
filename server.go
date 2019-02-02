package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type LifeInfo struct {
	Happy string `json:"happy"`

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

	defer context.Request().Body.Close()
	err := json.NewDecoder(context.Request().Body).Decode(&lifeInfo) // The body of data that the user is trying to send through the request

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Here is the life info object %s", lifeInfo)
	return context.String(http.StatusOK, "Accepted request")
}

func main() {
	fmt.Println("Server is running on port 4000!")

	// Instantiate the client that is going to be trigerring the requests
	client := echo.New() // Creating an instance of our server

	client.GET("/greeting", greeting) // The framework provides us the context?
	client.POST("/lifeInfo", lifeInfo)
	client.GET("/goodbye", goodbye)

	client.Start(":4000")

}
