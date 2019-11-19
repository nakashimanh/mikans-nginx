package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API-/")
	})

	e.GET("/api/ping", func(c echo.Context) error {
		layout := "2006-01-02 15:04:05"
		loc, _ := time.LoadLocation("Asia/Tokyo")
		now := time.Now().In(loc).Format(layout)
		return c.String(http.StatusOK, "API:Ping - "+now)
	})

	e.Logger.Fatal(e.Start(":1188"))
}
