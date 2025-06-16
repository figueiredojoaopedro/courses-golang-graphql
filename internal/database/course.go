package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db *sql.DB
	ID string
	Name string
	Description string
	CategoryID string
} 

func NewCourse(db *sql.DB) *Course{ 
	return &Course{db: db}
}

func (c *Course) Create (name, description, categoryId string) (Course, error){
	id := uuid.New().String()
	query := "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4);"

	_, err := c.db.Exec(query, id, name, description, categoryId)

	if err != nil {
		return Course{}, err
	}

	return Course{
		ID: id,
		Name: name,
		Description: description,
		CategoryID: categoryId,
	}, nil
}

func (c *Course) FindAll () ([]Course, error) {
	query := "SELECT id, name, description, category_id FROM courses;"

	results, err := c.db.Query(query);
	
	if err != nil {
		return nil, err;
	}

	defer results.Close()

	courses := []Course{}

	for results.Next(){ 
		var id, name, description, category_id string;

		if err := results.Scan(&id, &name,	&description, &category_id); err != nil {
			return nil, err; 
		}

		courses = append(courses, Course{ID: id, Name:  name, Description: description, CategoryID: category_id})
	}

	return courses, nil
}

func ( c*Course) FindByCategoryID (categoryId string) ([]Course, error){
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryId)

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next(){
		var id, name, description, category_id string

		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil{
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: category_id})
	}

	return courses, nil
}