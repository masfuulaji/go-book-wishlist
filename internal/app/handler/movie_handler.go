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
	MovieRepository    *repositories.MovieRepositoryImpl
	CategoryRepository *repositories.CategoryRepositoryImpl
	DirectorRepository *repositories.DirectorRepositoryImpl
}

func NewMovieHandler(db *sqlx.DB) *movieHandlerImpl {
	return &movieHandlerImpl{
		MovieRepository:    repositories.NewMovieRepository(db),
		CategoryRepository: repositories.NewCategoryRepository(db),
		DirectorRepository: repositories.NewDirectorRepository(db),
	}
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
	res.DirectorID = data.DirectorID

	categories, err := h.CategoryRepository.GetCategories()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var resCategory []response.CategoryResponse

	for _, category := range categories {
		resCategory = append(resCategory, response.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	directors, err := h.DirectorRepository.GetDirectors()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var resDirector []response.DirectorResponse

	for _, director := range directors {
		resDirector = append(resDirector, response.DirectorResponse{
			ID:           director.ID,
			Name:         director.Name,
			PlaceOfBirth: director.PlaceOfBirth,
		})
	}
	return movie.Form(res, resCategory, resDirector).Render(c.Request().Context(), c.Response())
}

func (h *movieHandlerImpl) NewMovie(c echo.Context) error {
	categories, err := h.CategoryRepository.GetCategories()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var resCategory []response.CategoryResponse

	for _, category := range categories {
		resCategory = append(resCategory, response.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	directors, err := h.DirectorRepository.GetDirectors()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var resDirector []response.DirectorResponse

	for _, director := range directors {
		resDirector = append(resDirector, response.DirectorResponse{
			ID:           director.ID,
			Name:         director.Name,
			PlaceOfBirth: director.PlaceOfBirth,
		})
	}
	return movie.Form(response.MovieResponse{}, resCategory, resDirector).Render(c.Request().Context(), c.Response())
}

func (h *movieHandlerImpl) StoreMovie(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	description := c.FormValue("description")
	year := c.FormValue("year")
	director := c.FormValue("director")

	intID := 0
	if id != "" {
		var err error
		intID, err = strconv.Atoi(id)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}

	intDirector, err := strconv.Atoi(director)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	movie := &request.MovieRequest{
		ID:          intID,
		Title:       title,
		Description: description,
		Year:        year,
		DirectorID:  intDirector,
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
