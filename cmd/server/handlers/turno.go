package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
	"github.com/andres9521/Go-Web-Desafio-Final/internal/layers/turno"
	"github.com/andres9521/Go-Web-Desafio-Final/pkg/web"
	"github.com/gin-gonic/gin"
)

type TurnoHandler struct {
	TurnoService turno.ITurnoService
}

func (th *TurnoHandler) Post(c *gin.Context) {

	var turno domain.Turno

	err := c.ShouldBindJSON(&turno)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	t, err := th.TurnoService.Create(turno)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusCreated, t)

}

func (th *TurnoHandler) GetById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	turno, err := th.TurnoService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "turn not found")
		return
	}

	web.Success(c, http.StatusOK, turno)

}

func (th *TurnoHandler) GetAll(c *gin.Context) {

	turnos, err := th.TurnoService.GetAll()
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, turnos)

}

func (th *TurnoHandler) Patch(c *gin.Context) {

	type Request struct {
		Descripcion string
		PacienteId  int
		DentistaId  int
	}

	var r Request
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	update, err := th.TurnoService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "turn not found")
		return
	}

	if err := c.ShouldBindJSON(&r); err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	if r.Descripcion != "" {
		update.Descripcion = r.Descripcion
	}

	if r.PacienteId != 0 {
		update.PacienteId = r.PacienteId
	}

	if r.DentistaId != 0 {
		update.DentistaId = r.DentistaId
	}

	t, err := th.TurnoService.UpdateOne(id, *update)
	if err != nil {
		web.NewApiError(c, http.StatusConflict, "conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, t)

}

func (th *TurnoHandler) Put(c *gin.Context) {

	var turno domain.Turno
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	_, err = th.TurnoService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "turn not found")
		return
	}

	err = c.ShouldBindJSON(&turno)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	t, err := th.TurnoService.UpdateMany(id, turno)
	if err != nil {
		web.NewApiError(c, http.StatusConflict, "conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, t)

}

func (th *TurnoHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid id")
		return
	}

	_, err = th.TurnoService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound, "not_found", "turn not found")
		return
	}

	err = th.TurnoService.Delete(id)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, fmt.Sprintf("Turn id = %d has been deleted.", id))

}

func (th *TurnoHandler) PostByPatientDniAndDentistLicense(c *gin.Context) {

	dni := c.Query("dni")
	matricula := c.Query("matricula")
	var turno domain.Turno
	err := c.ShouldBindJSON(&turno)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", "invalid JSON")
		return
	}

	t, err := th.TurnoService.CreateByPatientDniAndDentistLicense(turno, dni, matricula)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	web.Success(c, http.StatusCreated, t)

}

func (th *TurnoHandler) GetByPatientDni(c *gin.Context) {

	dni := c.Query("dni")

	turno, err := th.TurnoService.GetByPatientDni(dni)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, turno)

}
