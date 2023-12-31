package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("insert into categories (id, name, description) values ($1, $2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("select id, name, description from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []Category
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (Category, error) {
	rows, err := c.db.Query("select ca.id, ca.name, ca.description from categories ca join courses co on ca.id = co.category_id where co.id = $1", courseId)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()
	var category Category
	if rows.Next() {
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return Category{}, err
		}
	}
	return category, nil
}

func (c *Category) Find(id string) (Category, error) {
	rows, err := c.db.Query("select id, name, description from categories where id = $1", id)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()
	var category Category
	if rows.Next() {
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return Category{}, err
		}
	}
	return category, nil
}
