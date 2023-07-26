package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/artificial-lua/example-webserver-with-echo-go/scraper"
	"github.com/labstack/echo"
)

const savedFileName string = "pages.csv"
const provideFileName string = "page.csv"

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func hamdleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("pages.csv")
	term := strings.ToLower(scraper.CleanString(c.FormValue(("term"))))
	scraper.Scraper("ff14", 4337, term)
	return c.Attachment(savedFileName, provideFileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.GET("/home", hamdleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
