package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	start := time.Now()

	e := echo.New()
	fmt.Println("------")
	t := time.Now()

	elapsed := t.Sub(start)

	fmt.Println("time = " + elapsed.String())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!.....1234")
	})

	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1324"))

}
