package repository

import (
	"context"
	"database/sql"

	"github.com/rozanlaudzai/go-mysql-restful-api/exception"
	"github.com/rozanlaudzai/go-mysql-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {

	categories := []domain.Category{}

	query := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	var category domain.Category
	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	query := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	if err != nil {
		return category, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return category, err
	}
	category.Id = int(lastId)

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {

	category := domain.Category{}

	query := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	if err != nil {
		return category, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&category.Id, &category.Name)
		return category, err
	}

	return category, exception.NewNotFoundError("category not found")
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	query := "UPDATE category SET name = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	if err != nil {
		return category, err
	}

	// check rows affected, if it's 0 then category not found
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return category, err
	}
	if rowsAffected == 0 {
		return category, exception.NewNotFoundError("category not found")
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, categoryId int) error {
	query := "DELETE FROM category WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, categoryId)
	if err != nil {
		return err
	}

	// check rows affected, if it's 0 then category not found
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return exception.NewNotFoundError("category not found")
	}

	return nil
}
