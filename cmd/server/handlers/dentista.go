package handlers

import (
	"net/http"
	"strconv"

	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
	"github.com/andres9521/Go-Web-Desafio-Final/internal/layers/dentista"
	"github.com/andres9521/Go-Web-Desafio-Final/pkg/web"
	"github.com/gin-gonic/gin"
)

type DentistaHandler struct {
	DentistaService dentista.IDentistaService
}

func (dh *DentistaHandler) Post(c *gin.Context) {

	var dentista domain.Dentista

	err := c.ShouldBindJSON(&dentista)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
	}

	d, err := dh.DentistaService.Create(dentista)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
	}

	web.Success(c, http.StatusCreated, d)

}

func (dh *DentistaHandler) GetById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	dentista, err := dh.DentistaService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "dentist not found")
		return
	}

	web.Success(c, http.StatusOK, dentista)

}

func (dh *DentistaHandler) GetAll(c *gin.Context) {

	dentistas, err := dh.DentistaService.GetAll()
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
	}

	web.Success(c, http.StatusOK, dentistas)

}

func (dh *DentistaHandler) Patch(c *gin.Context) {

	type Request struct {
		Nombre    string `json:"nombre"`
		Apellido  string `json:"apellido"`
		Matricula string `json:"matricula"`
	}

	var r Request
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	targetDentista, err := dh.DentistaService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "dentist not found")
		return
	}

	if err := c.ShouldBindJSON(&r); err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	var update domain.Dentista

	if r.Nombre != "" {
		update = domain.Dentista{
			Nombre:    r.Nombre,
			Apellido:  targetDentista.Apellido,
			Matricula: targetDentista.Matricula,
		}
	}

	if r.Apellido != "" {
		update = domain.Dentista{
			Nombre:    targetDentista.Nombre,
			Apellido:  r.Apellido,
			Matricula: targetDentista.Matricula,
		}
	}

	if r.Matricula != "" {
		update = domain.Dentista{
			Nombre:    targetDentista.Nombre,
			Apellido:  targetDentista.Apellido,
			Matricula: r.Matricula,
		}
	}

	d, err := dh.DentistaService.UpdateOne(id, update)
	if err != nil {
		web.NewApiError(c, http.StatusConflict, "conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, d)

}

func (dh *DentistaHandler) Put(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	_, err = dh.DentistaService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "dentist not found")
		return
	}

	var dentista domain.Dentista
	err = c.ShouldBindJSON(&dentista)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	d, err := dh.DentistaService.UpdateMany(id, dentista)
	if err != nil {
		web.NewApiError(c, http.StatusConflict, "conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, d)

}

func (dh *DentistaHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	err = dh.DentistaService.Delete(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "dentist not found")
	}

	web.Success(c, http.StatusOK, nil)

}
