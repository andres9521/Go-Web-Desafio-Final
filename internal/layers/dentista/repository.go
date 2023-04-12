package dentista

import (
	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
	dentistastore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/dentistaStore"
)

type IDentistaRepository interface {
	Create(d domain.Dentista) (*domain.Dentista, error)
	GetById(id int) (*domain.Dentista, error)
	GetAll() ([]domain.Dentista, error)
	UpdateOne(id int, d domain.Dentista) (*domain.Dentista, error)
	UpdateMany(id int, d domain.Dentista) (*domain.Dentista, error)
	Delete(id int) error
}

type DentistaRepository struct {
	Store dentistastore.DentistaStoreInterface
}

func (dr *DentistaRepository) Create(d domain.Dentista) (*domain.Dentista, error) {

	dentista, err := dr.Store.Create(d)

	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (dr *DentistaRepository) GetById(id int) (*domain.Dentista, error) {

	dentista, err := dr.Store.GetById(id)

	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (dr *DentistaRepository) GetAll() ([]domain.Dentista, error) {

	dentistas, err := dr.Store.GetAll()

	if err != nil {
		return nil, err
	}

	return dentistas, nil

}

func (dr *DentistaRepository) UpdateOne(id int, d domain.Dentista) (*domain.Dentista, error) {

	dentista, err := dr.Store.UpdateOne(id, d)

	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (dr *DentistaRepository) UpdateMany(id int, d domain.Dentista) (*domain.Dentista, error) {

	dentista, err := dr.Store.UpdateOne(id, d)

	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (dr *DentistaRepository) Delete(id int) error {

	if err := dr.Store.Delete(id); err != nil {
		return err
	}

	return nil

}
