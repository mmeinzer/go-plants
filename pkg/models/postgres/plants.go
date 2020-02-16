package postgres

import (
	"database/sql"
	"errors"

	"mattmeinzer.com/plants/pkg/models"
)

// PlantModel which wraps a sql.DB connection pool
type PlantModel struct {
	DB *sql.DB
}

// Insert will insert a new snippet into the database
func (m *PlantModel) Insert(ownerID int, name string) (int, error) {
	stmt := `INSERT INTO plants(name, owner) VALUES($1, $2) RETURNING id`

	var id int
	err := m.DB.QueryRow(stmt, name, ownerID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get will return a specific snippet based on its id
func (m *PlantModel) Get(id int) (*models.Plant, error) {
	plant := &models.Plant{}
	err := m.DB.QueryRow(`SELECT id, name FROM plants WHERE id = $1`, id).Scan(&plant.ID, &plant.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return plant, nil
}

// Top will return 10 of the users plants
func (m *PlantModel) Top() ([]*models.Plant, error) {
	stmt := `SELECT id, name FROM plants ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plants := []*models.Plant{}

	for rows.Next() {
		p := &models.Plant{}

		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, err
		}

		plants = append(plants, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return plants, nil
}
