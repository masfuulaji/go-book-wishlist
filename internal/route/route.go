package route

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/internal/app/handler"
	"github.com/masfuulaji/go-book-wishlist/internal/database"
)

func SetupRoutes(e *echo.Echo) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	categoryHandler := handler.NewCategoryHandler(db.DB)
	e.GET("/category/delete/:id", categoryHandler.DeleteCategory)
	e.GET("/category", categoryHandler.PageCategory)
	e.GET("/category/edit/:id", categoryHandler.EditCategory)
	e.GET("/category/new", categoryHandler.NewCategory)
	e.POST("/category", categoryHandler.StoreCategory)
}
