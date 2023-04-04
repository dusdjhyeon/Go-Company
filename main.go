package main

import (
	"fmt"
	"strings"

	scrapper "go-price-comparison/scrapper"

	"github.com/labstack/echo"
)

var fileName = "jobs.csv"

// Handler function
// 기본 home.html로 연결 func
func Handler(c echo.Context) error {
	return c.File("home.html")
}

// HandleFunc function
// 검색 쿼리에 대해 스크래핑 요청 func
func HandleFunc(c echo.Context) error {
	query := strings.ToLower(scrapper.CleanString(c.FormValue("query")))
	fmt.Println(query)
	scrapper.Scrapper(query)

	return c.Attachment(fileName, query+".csv")
}

func main() {
	e := echo.New()
	e.GET("/", Handler)
	e.POST("/scrape", HandleFunc)
	e.Logger.Fatal(e.Start(":1323"))

}
