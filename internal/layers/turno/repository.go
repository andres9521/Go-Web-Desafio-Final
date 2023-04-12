package turno

import (
	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
	turnostore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/turnoStore"
)

type ITurnoRepository interface {
	Create(t domain.Turno) (*domain.Turno, error)
	GetById(id int) (*domain.Turno, error)
	GetAll() ([]domain.Turno, error)
	UpdateOne(id int, t domain.Turno) (*domain.Turno, error)
	UpdateMany(id int, t domain.Turno) (*domain.Turno, error)
	Delete(id int) error
	CreateByPatientDniAndDentistLicense(t domain.Turno, dni string, matricula string) (*domain.Turno, error)
	GetByPatientDni(dni string) (*domain.Turno, error)
}

type TurnoRepository struct {
	Store turnostore.TurnoStoreInterface
}

func (tr *TurnoRepository) Create(t domain.Turno) (*domain.Turno, error) {

	turno, err := tr.Store.Create(t)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (tr *TurnoRepository) GetById(id int) (*domain.Turno, error) {

	turno, err := tr.Store.GetById(id)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (tr *TurnoRepository) GetAll() ([]domain.Turno, error) {

	turnos, err := tr.Store.GetAll()
	if err != nil {
		return nil, err
	}

	return turnos, nil

}

func (tr *TurnoRepository) UpdateOne(id int, t domain.Turno) (*domain.Turno, error) {

	turno, err := tr.Store.UpdateOne(id, t)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (tr *TurnoRepository) UpdateMany(id int, t domain.Turno) (*domain.Turno, error) {

	turno, err := tr.Store.UpdateMany(id, t)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (tr *TurnoRepository) Delete(id int) error {

	if err := tr.Store.Delete(id); err != nil {
		return err
	}

	return nil

}

func (tr *TurnoRepository) CreateByPatientDniAndDentistLicense(t domain.Turno, dni string, matricula string) (*domain.Turno, error) {

	turno, err := tr.Store.CreateByPatientDniAndDentistLicense(t, dni, matricula)
	if err != nil {
		return nil, err
	}

	return turno, nil

}

func (tr *TurnoRepository) GetByPatientDni(dni string) (*domain.Turno, error) {

	turno, err := tr.Store.GetByPatientDni(dni)
	if err != nil {
		return nil, err
	}

	return turno, nil

}
