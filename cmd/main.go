package main

import (
	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/internal/route"
)

func main() {
	e := echo.New()
	route.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
