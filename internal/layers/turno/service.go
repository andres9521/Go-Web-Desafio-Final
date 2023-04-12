package turno

import (
	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
)

type ITurnoService interface {
	Create(t domain.Turno) (*domain.Turno, error)
	GetById(id int) (*domain.Turno, error)
	GetAll() ([]domain.Turno, error)
	UpdateOne(id int, t domain.Turno) (*domain.Turno, error)
	UpdateMany(id int, t domain.Turno) (*domain.Turno, error)
	Delete(id int) error
	CreateByPatientDniAndDentistLicense(t domain.Turno, dni string, matricula string) (*domain.Turno, error)
	GetByPatientDni(dni string) (*domain.Turno, error)
}

type TurnoService struct {
	Repository ITurnoRepository
}

func (ts *TurnoService) Create(t domain.Turno) (*domain.Turno, error) {

	turno, err := ts.Repository.Create(t)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (ts *TurnoService) GetById(id int) (*domain.Turno, error) {

	turno, err := ts.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (ts *TurnoService) GetAll() ([]domain.Turno, error) {

	turnos, err := ts.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return turnos, nil

}

func (ts *TurnoService) UpdateOne(id int, t domain.Turno) (*domain.Turno, error) {

	turno, err := ts.Repository.UpdateOne(id, t)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (ts *TurnoService) UpdateMany(id int, t domain.Turno) (*domain.Turno, error) {

	turno, err := ts.Repository.UpdateMany(id, t)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (ts *TurnoService) Delete(id int) error {

	if err := ts.Repository.Delete(id); err != nil {
		return err
	}

	return nil

}

func (ts *TurnoService) CreateByPatientDniAndDentistLicense(t domain.Turno, dni string, matricula string) (*domain.Turno, error) {

	turno, err := ts.Repository.CreateByPatientDniAndDentistLicense(t, dni, matricula)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (ts *TurnoService) GetByPatientDni(dni string) (*domain.Turno, error) {

	turno, err := ts.Repository.GetByPatientDni(dni)
	if err != nil {
		return nil, err
	}

	return turno, nil

}
