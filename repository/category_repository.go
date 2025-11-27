package repository

import (
	"context"
	"database/sql"

	"github.com/rozanlaudzai/go-mysql-restful-api/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	DeleteById(ctx context.Context, tx *sql.Tx, categoryId int) error
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
}
