package repository

import (
	"context"
	"fmt"
	go_database "go-database"
	"go-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCategoryInsert(t *testing.T) {
	categoryRepository := NewCategoryRepository(go_database.GetConnection())

	ctx := context.Background()
	category := entity.Category{
		Name:        "test repository",
		Description: "test repository description",
	}

	result, err := categoryRepository.Insert(ctx, category)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCategoryFindById(t *testing.T) {
	categoryRepository := NewCategoryRepository(go_database.GetConnection())

	ctx := context.Background()
	category, err := categoryRepository.FindById(ctx, 35)

	if err != nil {
		panic(err)
	}

	fmt.Println(category)
}

func TestCategoryFindAll(t *testing.T) {
	categoryRepository := NewCategoryRepository(go_database.GetConnection())

	ctx := context.Background()
	categories, err := categoryRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category)
	}
}
