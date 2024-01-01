package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/internal/app/handler"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/category", handler.PageCategory)
	e.GET("/category/edit/:id", handler.EditCategory)
	e.GET("/category/new", handler.NewCategory)
}
