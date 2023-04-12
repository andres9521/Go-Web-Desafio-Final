package dentistaStore

import (
	"database/sql"
	"fmt"

	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
)

type DentistaSqlStore struct {
	DB *sql.DB
}

func (s *DentistaSqlStore) Create(d domain.Dentista) (*domain.Dentista, error) {

	var dentista domain.Dentista

	query := "INSERT INTO dentistas (nombre, apellido, matricula) VALUES (?,?,?);"
	row := s.DB.QueryRow(query, d.Nombre, d.Apellido, d.Matricula)

	err := row.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return nil, err
	}

	return &dentista, nil

}

func (s *DentistaSqlStore) GetById(id int) (*domain.Dentista, error) {

	var dentista domain.Dentista

	query := "SELECT * FROM dentistas WHERE id = ?;"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return nil, err
	}

	return &dentista, nil

}
func (s *DentistaSqlStore) GetAll() ([]domain.Dentista, error) {

	var dentistas []domain.Dentista

	query := "SELECT * FROM dentistas;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dentista domain.Dentista

		err := rows.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
		if err != nil {
			return nil, err
		}

		dentistas = append(dentistas, dentista)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dentistas, nil

}

func (s *DentistaSqlStore) UpdateOne(id int, d domain.Dentista) (*domain.Dentista, error) {

	var dentista domain.Dentista

	query := "UPDATE dentistas SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;"
	row := s.DB.QueryRow(query, d.Nombre, d.Apellido, d.Matricula, id)

	err := row.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return nil, err
	}

	return &dentista, nil

}

func (s *DentistaSqlStore) UpdateMany(id int, d domain.Dentista) (*domain.Dentista, error) {

	var dentista domain.Dentista

	query := "UPDATE dentistas SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;"
	row := s.DB.QueryRow(query, d.Nombre, d.Apellido, d.Matricula, id)

	err := row.Scan(&dentista.Id, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return nil, err
	}

	return &dentista, nil

}

func (s *DentistaSqlStore) Delete(id int) error {

	query := "DELETE FROM dentistas WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil

}

func (s *DentistaSqlStore) GetIdByLicense(matricula string) (int, error) {

	var id int

	query := "SELECT id FROM dentistas WHERE matricula = ?"
	row := s.DB.QueryRow(query, matricula)

	err := row.Scan(&id)
	fmt.Println(err)
	if err != nil {
		return -1, err
	}

	return id, nil

}
