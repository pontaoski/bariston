package web

import (
	"baritone/bot/templating"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Host() {
	app := echo.New()
	app.Renderer = &templating.Template{}
	app.Static("/static", "assets")
	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})
	err := app.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
