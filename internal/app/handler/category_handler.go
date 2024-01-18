package handler

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/internal/app/repositories"
	"github.com/masfuulaji/go-book-wishlist/internal/app/request"
	"github.com/masfuulaji/go-book-wishlist/internal/app/response"
	"github.com/masfuulaji/go-book-wishlist/views/category"
)

type CategoryHandler interface {
	PageCategory(c echo.Context) error
	EditCategory(c echo.Context) error
	NewCategory(c echo.Context) error
	StoreCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}

type categoryHandlerImpl struct {
	CategoryRepository *repositories.CategoryRepositoryImpl
}

func NewCategoryHandler(db *sqlx.DB) *categoryHandlerImpl {
	return &categoryHandlerImpl{CategoryRepository: repositories.NewCategoryRepository(db)}
}

func (h *categoryHandlerImpl) PageCategory(c echo.Context) error {

	categories, err := h.CategoryRepository.GetCategories()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var res []response.CategoryResponse

	for _, category := range categories {
		res = append(res, response.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return category.Category(res).Render(c.Request().Context(), c.Response())
}

func (h *categoryHandlerImpl) EditCategory(c echo.Context) error {
	id := c.Param("id")

	data, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return err
	}

	var res response.CategoryResponse

	res.ID = data.ID
	res.Name = data.Name
	res.Description = data.Description

	return category.Form(res).Render(c.Request().Context(), c.Response())
}

func (h *categoryHandlerImpl) NewCategory(c echo.Context) error {
	return category.Form(response.CategoryResponse{}).Render(c.Request().Context(), c.Response())
}

func (h *categoryHandlerImpl) StoreCategory(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	description := c.FormValue("description")

	intID := 0
	if id != "" {
		var err error
		intID, err = strconv.Atoi(id)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}
	category := &request.CategoryRequest{
		ID:          intID,
		Name:        name,
		Description: description,
	}

	if id == "" {
		err := h.CategoryRepository.CreateCategory(category)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	} else {
		err := h.CategoryRepository.UpdateCategory(id, category)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}

	return c.Redirect(http.StatusFound, "/category")
}

func (h *categoryHandlerImpl) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	err := h.CategoryRepository.DeleteCategory(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusFound, "/category")
}
