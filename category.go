package main

import (
	"budgenius/repository"
	"fmt"
	"log"
)

type Category struct {
	ID   int64
	Name string
	Type string
}

func (a *App) CreateCategory(category Category) error {
	log.Printf("Create category %q with type %q\n", category.Name, category.Type)

	categoryType := repository.CategoryType(category.Type)
	if !categoryType.IsValid() {
		err := fmt.Errorf("unknown category type %q", category.Type)
		log.Println(err)
		return err
	}

	err := a.repository.CreateCategory(a.ctx, category.Name, repository.CategoryType(category.Type))
	if err != nil {
		log.Println("Error creating category:", err)
		return err
	}

	log.Printf("Category %q with type %s created successfully\n", category.Name, category.Type)

	return nil
}

func (a *App) Categories() ([]Category, error) {
	log.Println("Request categories")

	categories, err := a.repository.Categories(a.ctx)
	if err != nil {
		log.Println("Error getting categories:", err)
		return nil, err
	}

	c := make([]Category, len(categories))
	for i, category := range categories {
		c[i] = Category{
			ID:   category.ID,
			Name: category.Name,
			Type: string(category.Type),
		}
	}

	log.Printf("Successfully got %d categories\n", len(c))

	return c, nil
}

func (a *App) DeleteCategory(id int64) error {
	log.Printf("Delete category with id %d\n", id)

	err := a.repository.DeleteCategory(a.ctx, id)
	if err != nil {
		log.Println("Error deleting category:", err)
		return err
	}

	log.Printf("Category with id %d deleted successfully\n", id)

	return nil
}
