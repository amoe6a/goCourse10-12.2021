package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func calculateDays(date string) int {
	dates := strings.Split(date, ".")
	year, _ := strconv.Atoi(dates[2])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[0])

	locAlmaty, _ := time.LoadLocation("Asia/Almaty")

	dateDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, locAlmaty)

	currentDate := time.Now()

	diff := currentDate.Sub(dateDate)
	return int(diff.Seconds()/86400)
}

func getLifeSec(c echo.Context) error {
	birthDate := c.FormValue("date_of_birth")
	return c.String(http.StatusOK, "Currently, it is " + time.Now().String() + "\nYou have lived " + strconv.Itoa(calculateDays(birthDate)) + " days\n")
}

func main() {
	e := echo.New()

	// curl -F "date_of_birth=DD.MM.YYYY" http://localhost:1323/knowDays
	e.POST("/knowDays", getLifeSec)

	// http://localhost:1323/greet/yourname
	e.GET("/greet/:id", func(c echo.Context) error {
		name := c.Param("id")
		return c.String(http.StatusOK, "Hello " + name)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
