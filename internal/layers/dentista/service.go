package dentista

import "github.com/andres9521/Go-Web-Desafio-Final/internal/domain"

type IDentistaService interface {
	Create(d domain.Dentista) (*domain.Dentista, error)
	GetById(id int) (*domain.Dentista, error)
	GetAll() ([]domain.Dentista, error)
	UpdateOne(id int, d domain.Dentista) (*domain.Dentista, error)
	UpdateMany(id int, d domain.Dentista) (*domain.Dentista, error)
	Delete(id int) error
}

type DentistaService struct {
	Repository IDentistaRepository
}

func (ds *DentistaService) Create(d domain.Dentista) (*domain.Dentista, error) {

	dentista, err := ds.Repository.Create(d)
	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (ds *DentistaService) GetById(id int) (*domain.Dentista, error) {

	dentista, err := ds.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (ds *DentistaService) GetAll() ([]domain.Dentista, error) {

	dentistas, err := ds.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return dentistas, nil

}

func (ds *DentistaService) UpdateOne(id int, d domain.Dentista) (*domain.Dentista, error) {

	dentista, err := ds.Repository.UpdateOne(id, d)
	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (ds *DentistaService) UpdateMany(id int, d domain.Dentista) (*domain.Dentista, error) {

	dentista, err := ds.Repository.UpdateMany(id, d)
	if err != nil {
		return nil, err
	}

	return dentista, nil

}

func (ds *DentistaService) Delete(id int) error {

	if err := ds.Repository.Delete(id); err != nil {
		return err
	}

	return nil

}
