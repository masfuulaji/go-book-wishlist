package handler

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/masfuulaji/go-book-wishlist/internal/app/repositories"
	"github.com/masfuulaji/go-book-wishlist/internal/app/request"
	"github.com/masfuulaji/go-book-wishlist/internal/app/response"
	"github.com/masfuulaji/go-book-wishlist/views/movie"
)

type MovieHandler interface {
	PageMovie(c echo.Context) error
	EditMovie(c echo.Context) error
	NewMovie(c echo.Context) error
	StoreMovie(c echo.Context) error
	DeleteMovie(c echo.Context) error
}

type movieHandlerImpl struct {
	MovieRepository *repositories.MovieRepositoryImpl
}

func NewMovieHandler(db *sqlx.DB) *movieHandlerImpl {
	return &movieHandlerImpl{MovieRepository: repositories.NewMovieRepository(db)}
}

func (h *movieHandlerImpl) PageMovie(c echo.Context) error {

	movies, err := h.MovieRepository.GetMovies()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var res []response.MovieResponse

	for _, movie := range movies {
		res = append(res, response.MovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Year:        movie.Year,
		})
	}

	return movie.Movie(res).Render(c.Request().Context(), c.Response())
}

func (h *movieHandlerImpl) EditMovie(c echo.Context) error {
	id := c.Param("id")
	data, err := h.MovieRepository.GetMovie(id)
	if err != nil {
		return err
	}

	var res response.MovieResponse

	res.ID = data.ID
	res.Title = data.Title
	res.Description = data.Description
	res.Year = data.Year
	return movie.Form(res).Render(c.Request().Context(), c.Response())
}

func (h *movieHandlerImpl) NewMovie(c echo.Context) error {
	return movie.Form(response.MovieResponse{}).Render(c.Request().Context(), c.Response())
}

func (h *movieHandlerImpl) StoreMovie(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	description := c.FormValue("description")
	year := c.FormValue("year")

	intID := 0
	if id != "" {
		var err error
		intID, err = strconv.Atoi(id)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}

	movie := &request.MovieRequest{
		ID:          intID,
		Title:       title,
		Description: description,
		Year:        year,
	}

	if id == "" {
		err := h.MovieRepository.CreateMovie(movie)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	} else {
		err := h.MovieRepository.UpdateMovie(id, movie)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}

	return c.Redirect(http.StatusFound, "/movie")
}

func (h *movieHandlerImpl) DeleteMovie(c echo.Context) error {
	id := c.Param("id")

	err := h.MovieRepository.DeleteMovie(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusFound, "/movie")
}
