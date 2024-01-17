package route

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/internal/app/handler"
	"github.com/masfuulaji/go-book-wishlist/internal/database"
	"github.com/masfuulaji/go-book-wishlist/views"
)

func SetupRoutes(e *echo.Echo) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))))
	e.GET("/", func(c echo.Context) error {
		return views.Home().Render(c.Request().Context(), c.Response())
	})

	categoryHandler := handler.NewCategoryHandler(db.DB)
	e.GET("/category/delete/:id", categoryHandler.DeleteCategory)
	e.GET("/category", categoryHandler.PageCategory)
	e.GET("/category/edit/:id", categoryHandler.EditCategory)
	e.GET("/category/new", categoryHandler.NewCategory)
	e.POST("/category", categoryHandler.StoreCategory)

	directorHandler := handler.NewDirectorHandler(db.DB)
	e.GET("/director/delete/:id", directorHandler.DeleteDirector)
	e.GET("/director", directorHandler.PageDirector)
	e.GET("/director/edit/:id", directorHandler.EditDirector)
	e.GET("/director/new", directorHandler.NewDirector)
	e.POST("/director", directorHandler.StoreDirector)
}
