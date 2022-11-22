package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
)

// Bussiness logic here

type categoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepositoryImpl{DB: db}
}

func (repository *categoryRepositoryImpl) Insert(ctx context.Context, category entity.Category) (entity.Category, error) {
	script := "INSERT INTO categories (name, description) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, category.Name, category.Description)

	if err != nil {
		return category, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return category, err
	}
	category.Id = int32(id)
	return category, nil
}

func (repository *categoryRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Category, error) {
	script := "SELECT id, name, description FROM categories WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	category := entity.Category{}

	if err != nil {
		return category, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name, &category.Description)
		return category, err
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *categoryRepositoryImpl) FindAll(ctx context.Context) ([]entity.Category, error) {
	script := "SELECT id, name, description FROM categories"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		rows.Scan(&category.Id, &category.Name, &category.Description)
		categories = append(categories, category)
	}

	return categories, nil
}
