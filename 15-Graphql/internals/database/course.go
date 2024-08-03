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
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(name, description, categoryID string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO course (id, name, description, category_id) VALUES($1, $2, $3, $4)", id, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}
	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: category_id})
	}
	return courses, nil
}

// func (c *Category) FindAll() ([]Category, error) {
// 	rows, err := c.db.Query("SELECT * FROM categories")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	categories := []Category{}

// 	for rows.Next() {
// 		var id, name, description string
// 		if err := rows.Scan(&id, &name, &description); err != nil {
// 			return nil, err
// 		}
// 		categories = append(categories, Category{ID: id, Name: name, Description: description})
// 	}
// 	return categories, nil
// }
