package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryId string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("insert into courses (id, name, description, category_id) values($1, $2, $3, $4)",
		id, name, description, categoryId)
	if err != nil {
		return Course{}, err
	}
	return Course{ID: id, Name: name, Description: description, CategoryId: categoryId}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("select id, name, description, category_id from courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("select id, name, description, category_id from courses where category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
