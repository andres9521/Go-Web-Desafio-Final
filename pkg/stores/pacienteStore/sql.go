package pacienteStore

import (
	"database/sql"
	"fmt"
	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
)

type PacienteSqlStore struct {
	DB *sql.DB
}

func (s *PacienteSqlStore) Create(p domain.Paciente) (*domain.Paciente, error) {

	var paciente domain.Paciente

	query := "INSERT INTO pacientes (name, apellido, domicilio, dni, fechaDeAlta) VALUES (?,?,?,?,?)"
	row := s.DB.QueryRow(query, p.Nombre, p.Apellido, p.Domicilio, p.DNI, p.FechaDeAlta)

	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta)
	if err != nil {
		return nil, err
	}

	return &paciente, nil

}

func (s *PacienteSqlStore) GetById(id int) (*domain.Paciente, error) {

	var paciente domain.Paciente

	query := "SELECT * FROM pacientes WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta)
	if err != nil {
		return nil, err
	}

	return &paciente, nil

}

func (s *PacienteSqlStore) GetAll() ([]domain.Paciente, error) {

	var pacientes []domain.Paciente

	query := "SELECT * FROM pacientes"
	rows, err := s.DB.Query(query)
	if err != nil {
		fmt.Println("error 1")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var paciente domain.Paciente

		err := rows.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta)
		if err != nil {
			return nil, err
		}

		pacientes = append(pacientes, paciente)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pacientes, nil

}

func (s *PacienteSqlStore) UpdateOne(id int, p domain.Paciente) (*domain.Paciente, error) {

	var paciente domain.Paciente

	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fechaDeAlta = ? WHERE id = ?"
	row := s.DB.QueryRow(query, p.Nombre, p.Apellido, p.Domicilio, p.DNI, p.FechaDeAlta, id)

	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta)
	if err != nil {
		return nil, err
	}

	return &paciente, nil

}

func (s *PacienteSqlStore) UpdateMany(id int, p domain.Paciente) (*domain.Paciente, error) {

	var paciente domain.Paciente

	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fechaDeAlta = ? WHERE id = ?"
	row := s.DB.QueryRow(query, p.Nombre, p.Apellido, p.Domicilio, p.DNI, p.FechaDeAlta, id)

	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta)
	if err != nil {
		return nil, err
	}

	return &paciente, nil
}

func (s *PacienteSqlStore) Delete(id int) error {

	query := "DELETE FROM pacientes WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil

}

func (s *PacienteSqlStore) GetIdByDni(dni string) (int, error) {

	var id int

	query := "SELECT id FROM pacientes WHERE dni = ?"
	result := s.DB.QueryRow(query, dni)

	err := result.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil

}
