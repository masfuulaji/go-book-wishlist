package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-book-wishlist/internal/app/models"
	"github.com/masfuulaji/go-book-wishlist/internal/app/request"
)

type DirectorRepositoryImpl struct {
	db *sqlx.DB
}

func NewDirectorRepository(db *sqlx.DB) *DirectorRepositoryImpl {
	return &DirectorRepositoryImpl{db: db}
}

func (dr *DirectorRepositoryImpl) CreateDirector(director *request.Director) error {
	query := "INSERT INTO directors (name, birthday, place_of_birth, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	formatBirday := director.Birthday.Format("2006-01-02")

	_, err := dr.db.Exec(query, director.Name, formatBirday, director.PlaceOfBirth, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (dr *DirectorRepositoryImpl) GetDirectors() ([]models.Director, error) {
	query := "SELECT * FROM directors where deleted_at is null"

	var directors []models.Director
	err := dr.db.Select(&directors, query)
	if err != nil {
		return nil, err
	}

	return directors, nil
}

func (dr *DirectorRepositoryImpl) GetDirector(id string) (models.Director, error) {
	query := "SELECT * FROM directors WHERE id = $1"

	var director models.Director
	err := dr.db.Get(&director, query, id)
	if err != nil {
		return models.Director{}, err
	}

	return director, nil
}

func (dr *DirectorRepositoryImpl) UpdateDirector(id string, director *request.Director) error {
	query := "UPDATE directors SET name = $1, birthday = $2, place_of_birth = $3, updated_at = $4 WHERE id = $5"
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := dr.db.Exec(query, director.Name, director.Birthday, director.PlaceOfBirth, updatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (dr *DirectorRepositoryImpl) DeleteDirector(id string) error {
	query := "UPDATE directors SET deleted_at = $1 WHERE id = $2"
	deletedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := dr.db.Exec(query, deletedAt, id)
	if err != nil {
		return err
	}

	return nil
}
