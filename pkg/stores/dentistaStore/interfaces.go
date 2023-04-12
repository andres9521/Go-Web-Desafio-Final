package dentistaStore

import "github.com/andres9521/Go-Web-Desafio-Final/internal/domain"

type DentistaStoreInterface interface {
	Create(d domain.Dentista) (*domain.Dentista, error)
	GetById(id int) (*domain.Dentista, error)
	GetAll() ([]domain.Dentista, error)
	UpdateOne(id int, d domain.Dentista) (*domain.Dentista, error)
	UpdateMany(id int, d domain.Dentista) (*domain.Dentista, error)
	Delete(id int) error
	GetIdByLicense(matricula string) (int, error)
}
