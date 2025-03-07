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
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	stmt, err := c.db.Prepare("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)")

	if err != nil {
		return Category{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) Find(id string) (Category, error) {
	var category Category
	err := c.db.QueryRow("SELECT * FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (c *Category) FindCategoryByCourseId(courseId string) (Category, error) {
	var category Category

	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c INNER JOIN courses co ON c.id = co.categoryId WHERE co.id = ?", courseId).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}
