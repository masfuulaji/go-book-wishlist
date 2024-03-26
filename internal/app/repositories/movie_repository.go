package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-book-wishlist/internal/app/models"
	"github.com/masfuulaji/go-book-wishlist/internal/app/request"
)

type MovieRepositoryImpl struct {
	db *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *MovieRepositoryImpl {
	return &MovieRepositoryImpl{db: db}
}

func (mr *MovieRepositoryImpl) CreateMovie(movie *request.MovieRequest) error {
	query := "INSERT INTO movies (title, description, year, director_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := mr.db.Exec(query, movie.Title, movie.Description, movie.Year, movie.DirectorID, createdAt, updatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MovieRepositoryImpl) GetMovies() ([]models.Movie, error) {
	query := "SELECT * FROM movies where deleted_at is null"

	var movies []models.Movie
	err := mr.db.Select(&movies, query)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (mr *MovieRepositoryImpl) GetMovie(id string) (models.Movie, error) {
	query := "SELECT * FROM movies WHERE id = $1"

	var movie models.Movie
	err := mr.db.Get(&movie, query, id)
	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func (mr *MovieRepositoryImpl) UpdateMovie(id string, movie *request.MovieRequest) error {
	query := "UPDATE movies SET title = $1, description = $2, year = $3, director_id = $4, updated_at = $5 WHERE id = $6"
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := mr.db.Exec(query, movie.Title, movie.Description, movie.Year, movie.DirectorID, updatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MovieRepositoryImpl) DeleteMovie(id string) error {
	query := "UPDATE movies SET deleted_at = $1 WHERE id = $2"
	deletedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := mr.db.Exec(query, deletedAt, id)
	if err != nil {
		return err
	}

	return nil
}
