package turnoStore

import "github.com/andres9521/Go-Web-Desafio-Final/internal/domain"

type TurnoStoreInterface interface {
	Create(t domain.Turno) (*domain.Turno, error)
	GetById(id int) (*domain.Turno, error)
	GetAll() ([]domain.Turno, error)
	UpdateOne(id int, t domain.Turno) (*domain.Turno, error)
	UpdateMany(id int, t domain.Turno) (*domain.Turno, error)
	Delete(id int) error
	CreateByPatientDniAndDentistLicense(t domain.Turno, dni string, matricula string) (*domain.Turno, error)
	GetByPatientDni(dni string) (*domain.Turno, error)
}
