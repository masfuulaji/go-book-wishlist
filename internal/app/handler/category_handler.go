package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/views/category"
)

func PageCategory(c echo.Context) error {
	return category.Category().Render(c.Request().Context(), c.Response())
}
