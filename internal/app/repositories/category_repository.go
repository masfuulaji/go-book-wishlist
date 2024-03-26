package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-book-wishlist/internal/app/models"
	"github.com/masfuulaji/go-book-wishlist/internal/app/request"
)

type CategoryRepositoryImpl struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db: db}
}

func (cr *CategoryRepositoryImpl) CreateCategory(category *request.CategoryRequest) error {
	query := "INSERT INTO categories (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := cr.db.Exec(query, category.Name, category.Description, createdAt, updatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepositoryImpl) GetCategories() ([]models.Category, error) {
	query := "SELECT * FROM categories where deleted_at is null"

	var categories []models.Category
	err := cr.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepositoryImpl) GetCategory(id string) (models.Category, error) {
	query := "SELECT * FROM categories WHERE id = $1"

	var category models.Category
	err := cr.db.Get(&category, query, id)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) UpdateCategory(id string, category *request.CategoryRequest) error {
	query := "UPDATE categories SET name = $1, description = $2, updated_at = $3 WHERE id = $4"
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := cr.db.Exec(query, category.Name, category.Description, updatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepositoryImpl) DeleteCategory(id string) error {
	query := "UPDATE categories SET deleted_at = $1 WHERE id = $2"
	deletedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := cr.db.Exec(query, deletedAt, id)
	if err != nil {
		return err
	}

	return nil
}
