package paciente

import (
	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
)

type IPacienteService interface {
	Create(p domain.Paciente) (*domain.Paciente, error)
	GetById(id int) (*domain.Paciente, error)
	GetAll() ([]domain.Paciente, error)
	UpdateOne(id int, p domain.Paciente) (*domain.Paciente, error)
	UpdateMany(id int, p domain.Paciente) (*domain.Paciente, error)
	Delete(id int) error
}

type PacienteService struct {
	Repository IPacienteRepository
}

func (ps *PacienteService) Create(p domain.Paciente) (*domain.Paciente, error) {

	paciente, err := ps.Repository.Create(p)

	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (ps *PacienteService) GetById(id int) (*domain.Paciente, error) {

	paciente, err := ps.Repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (ps *PacienteService) GetAll() ([]domain.Paciente, error) {

	pacientes, err := ps.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return pacientes, nil

}

func (ps *PacienteService) UpdateOne(id int, p domain.Paciente) (*domain.Paciente, error) {

	paciente, err := ps.Repository.UpdateOne(id, p)

	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (ps *PacienteService) UpdateMany(id int, p domain.Paciente) (*domain.Paciente, error) {

	paciente, err := ps.Repository.UpdateMany(id, p)

	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (ps *PacienteService) Delete(id int) error {

	if err := ps.Repository.Delete(id); err != nil {
		return err
	}

	return nil

}
