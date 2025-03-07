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

func (c *Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()
	stmt, err := c.db.Prepare("INSERT INTO courses (id, name, description, categoryId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return Course{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description, categoryId)
	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description}, nil

}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *Course) FindAllByCategoryID(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses WHERE categoryId = ?", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
