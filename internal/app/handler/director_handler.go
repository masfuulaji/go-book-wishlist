package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/internal/app/repositories"
	"github.com/masfuulaji/go-book-wishlist/internal/app/request"
	"github.com/masfuulaji/go-book-wishlist/internal/app/response"
	"github.com/masfuulaji/go-book-wishlist/views/director"
)

type DirectorHandler interface {
	PageDirector(c echo.Context) error
	EditDirector(c echo.Context) error
	NewDirector(c echo.Context) error
	StoreDirector(c echo.Context) error
	DeleteDirector(c echo.Context) error
}

type directorHandlerImpl struct {
	DirectorRepository *repositories.DirectorRepositoryImpl
}

func NewDirectorHandler(db *sqlx.DB) *directorHandlerImpl {
	return &directorHandlerImpl{DirectorRepository: repositories.NewDirectorRepository(db)}
}

func (h *directorHandlerImpl) PageDirector(c echo.Context) error {

	directors, err := h.DirectorRepository.GetDirectors()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var res []response.DirectorResponse

	for _, director := range directors {
		res = append(res, response.DirectorResponse{
			ID:           director.ID,
			Name:         director.Name,
			PlaceOfBirth: director.PlaceOfBirth,
			Birthday:     director.Birthday,
		})
	}
	return director.Director(res).Render(c.Request().Context(), c.Response())
}

func (h *directorHandlerImpl) EditDirector(c echo.Context) error {
	id := c.Param("id")
	data, err := h.DirectorRepository.GetDirector(id)
	if err != nil {
		return err
	}

	var res response.DirectorResponse
	res.ID = data.ID
	res.Name = data.Name
	res.Birthday = data.Birthday
	res.PlaceOfBirth = data.PlaceOfBirth
	return director.Form(res).Render(c.Request().Context(), c.Response())
}

func (h *directorHandlerImpl) NewDirector(c echo.Context) error {
	return director.Form(response.DirectorResponse{}).Render(c.Request().Context(), c.Response())
}

func (h *directorHandlerImpl) StoreDirector(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	birthday := c.FormValue("birthday")
	placeOfBirth := c.FormValue("place_of_birth")

	intID := 0
	if id != "" {
		var err error
		intID, err = strconv.Atoi(id)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}

	parseBirday, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	director := &request.Director{
		ID:           intID,
		Name:         name,
		Birthday:     parseBirday,
		PlaceOfBirth: placeOfBirth,
	}

	if id == "" {
		err := h.DirectorRepository.CreateDirector(director)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	} else {
		err := h.DirectorRepository.UpdateDirector(id, director)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}

	return c.Redirect(http.StatusFound, "/director")
}

func (h *directorHandlerImpl) DeleteDirector(c echo.Context) error {
	id := c.Param("id")

	err := h.DirectorRepository.DeleteDirector(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusFound, "/director")
}
