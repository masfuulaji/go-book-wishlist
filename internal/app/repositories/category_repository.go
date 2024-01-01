package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-book-wishlist/internal/app/models"
)

var category models.Category

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategories() ([]models.Category, error)
	GetCategory(id string) (models.Category, error)
	UpdateCategory(id string, category *models.Category) error
	DeleteCategory(id string) error
}

type categoryRepositoryImpl struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *categoryRepositoryImpl {
	return &categoryRepositoryImpl{db: db}
}

func (cr *categoryRepositoryImpl) CreateCategory(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2)"

	_, err := cr.db.Exec(query, category.Name, category.Description)
	if err != nil {
		return err
	}

	return nil
}

func (cr *categoryRepositoryImpl) GetCategories() ([]models.Category, error) {
	query := "SELECT * FROM categories"

	var categories []models.Category
	err := cr.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *categoryRepositoryImpl) GetCategory(id string) (models.Category, error) {
	query := "SELECT * FROM categories WHERE id = $1"

	var category models.Category
	err := cr.db.Get(&category, query, id)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (cr *categoryRepositoryImpl) UpdateCategory(id string, category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"

	_, err := cr.db.Exec(query, category.Name, category.Description, id)
	if err != nil {
		return err
	}

	return nil
}

func (cr *categoryRepositoryImpl) DeleteCategory(id string) error {
	query := "DELETE FROM categories WHERE id = $1"

	_, err := cr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
