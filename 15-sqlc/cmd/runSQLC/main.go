package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/mateus-sousa/fc-goexpert/15-sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
			course.CategoryName, course.ID, course.Name, course.Description, course.Price)
	}
}

func OldCrud(ctx context.Context, queries *db.Queries) {
	id := uuid.New().String()
	err := queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          id,
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})
	if err != nil {
		panic(err)
	}
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
		ID:          id,
	})
	if err != nil {
		panic(err)
	}
	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}
	err = queries.DeleteCategory(ctx, id)
	if err != nil {
		panic(err)
	}
}
