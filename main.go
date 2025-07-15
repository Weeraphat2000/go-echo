package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func getHello(c echo.Context) error {
	// This function handles the GET request to the /hello endpoint.
	fmt.Println("Received a request at /hello")
	return c.String(200, "Hello, World!")
}

func main() {
	// This is the entry point of the application.
	// You can add your application logic here.
	e := echo.New()

	e.GET("/hello", getHello)

	e.Logger.Fatal(e.Start(":8080"))
}
