package paciente

import (
	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
	pacientestore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/pacienteStore"
)

type IPacienteRepository interface {
	Create(p domain.Paciente) (*domain.Paciente, error)
	GetById(id int) (*domain.Paciente, error)
	GetAll() ([]domain.Paciente, error)
	UpdateOne(id int, p domain.Paciente) (*domain.Paciente, error)
	UpdateMany(id int, p domain.Paciente) (*domain.Paciente, error)
	Delete(id int) error
}

type PacienteRepository struct {
	Store pacientestore.PacienteStoreInterface
}

func (pr *PacienteRepository) Create(p domain.Paciente) (*domain.Paciente, error) {

	paciente, err := pr.Store.Create(p)
	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (pr *PacienteRepository) GetById(id int) (*domain.Paciente, error) {

	paciente, err := pr.Store.GetById(id)
	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (pr *PacienteRepository) GetAll() ([]domain.Paciente, error) {

	pacientes, err := pr.Store.GetAll()
	if err != nil {
		return nil, err
	}

	return pacientes, nil

}

func (pr *PacienteRepository) UpdateOne(id int, p domain.Paciente) (*domain.Paciente, error) {

	paciente, err := pr.Store.UpdateOne(id, p)
	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (pr *PacienteRepository) UpdateMany(id int, p domain.Paciente) (*domain.Paciente, error) {

	paciente, err := pr.Store.UpdateMany(id, p)
	if err != nil {
		return nil, err
	}

	return paciente, nil

}

func (pr *PacienteRepository) Delete(id int) error {

	if err := pr.Store.Delete(id); err != nil {
		return err
	}

	return nil

}
