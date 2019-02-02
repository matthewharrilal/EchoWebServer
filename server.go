package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func greeting(context echo.Context) error {
	// Context holds relevant information about the request such as data response body ...
	return context.String(http.StatusOK, "Thanks for visiting the server!")
}

func goodbye(context echo.Context) error {
	nameParam := context.QueryParam("name")
	// jsonObject := map[string]string{
	// 	"name": nameParam,
	// }
	// name := context.JSON(http.StatusOK, jsonObject)
	return context.String(http.StatusOK, fmt.Sprintf("Are you sure you want to leave %s?", nameParam))
}

func main() {
	fmt.Println("Server is running on port 4000!")

	// Instantiate the client that is going to be trigerring the requests
	client := echo.New() // Creating an instance of our server

	client.GET("/greeting", greeting) // The framework provides us the context?
	client.GET("/goodbye", goodbye)

	client.Start(":4000")

}
