package handlers

import (
	"net/http"
	"strconv"

	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
	"github.com/andres9521/Go-Web-Desafio-Final/internal/layers/paciente"
	"github.com/andres9521/Go-Web-Desafio-Final/pkg/web"
	"github.com/gin-gonic/gin"
)

type PacienteHandler struct {
	PacienteService paciente.IPacienteService
}

func (ph *PacienteHandler) Post(c *gin.Context) {

	var paciente domain.Paciente

	err := c.ShouldBindJSON(&paciente)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
	}

	p, err := ph.PacienteService.Create(paciente)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
	}

	web.Success(c, http.StatusCreated, p)

}

func (ph *PacienteHandler) GetById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	paciente, err := ph.PacienteService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "patient not found")
		return
	}

	web.Success(c, http.StatusOK, paciente)

}

func (ph *PacienteHandler) GetAll(c *gin.Context) {

	pacientes, err := ph.PacienteService.GetAll()
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
	}

	web.Success(c, http.StatusOK, pacientes)

}

func (ph *PacienteHandler) Patch(c *gin.Context) {

	type Request struct {
		Nombre      string `json:"nombre"`
		Apellido    string `json:"apellido"`
		Domicilio   string `json:"domicilio"`
		DNI         string `json:"dni"`
		FechaDeAlta string `json:"fechaDeAlta"`
	}

	var r Request
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	update, err := ph.PacienteService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "patient not found")
		return
	}

	err = c.ShouldBindJSON(&r)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	if r.Nombre != "" {
		update.Nombre = r.Nombre
	}

	if r.Apellido != "" {
		update.Apellido = r.Apellido
	}

	if r.Domicilio != "" {
		update.Domicilio = r.Domicilio
	}

	if r.DNI != "" {
		update.DNI = r.DNI
	}

	if r.FechaDeAlta != "" {
		update.FechaDeAlta = r.FechaDeAlta
	}

	result, err := ph.PacienteService.UpdateOne(id, *update)
	if err != nil {
		web.NewApiError(c, http.StatusConflict, "conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, result)

}

func (ph *PacienteHandler) Put(c *gin.Context) {

	var paciente domain.Paciente
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	_, err = ph.PacienteService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "patient not found")
		return
	}

	err = c.ShouldBindJSON(&paciente)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	p, err := ph.PacienteService.UpdateMany(id, paciente)
	if err != nil {
		web.NewApiError(c, http.StatusConflict, "conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, p)

}

func (ph *PacienteHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	_, err = ph.PacienteService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "patient not found")
		return
	}

	err = ph.PacienteService.Delete(id)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, nil)

}
