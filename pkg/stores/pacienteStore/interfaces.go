package pacienteStore

import "github.com/andres9521/Go-Web-Desafio-Final/internal/domain"

type PacienteStoreInterface interface {
	Create(p domain.Paciente) (*domain.Paciente, error)
	GetById(id int) (*domain.Paciente, error)
	GetAll() ([]domain.Paciente, error)
	UpdateOne(id int, p domain.Paciente) (*domain.Paciente, error)
	UpdateMany(id int, p domain.Paciente) (*domain.Paciente, error)
	Delete(id int) error
	GetIdByDni(dni string) (int, error)
}
