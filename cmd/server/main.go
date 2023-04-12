package main

import (
	"database/sql"

	"github.com/andres9521/Go-Web-Desafio-Final/cmd/server/handlers"
	"github.com/andres9521/Go-Web-Desafio-Final/internal/layers/dentista"
	"github.com/andres9521/Go-Web-Desafio-Final/internal/layers/paciente"
	"github.com/andres9521/Go-Web-Desafio-Final/internal/layers/turno"
	dentistastore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/dentistaStore"
	pacientestore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/pacienteStore"
	turnostore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/turnoStore"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:966457Saske+@tcp(localhost:3306)/my_db")
	if err != nil {
		panic(err)
	}

	errPing := db.Ping()
	if err != nil {
		panic(errPing)
	}

	dentistaStorage := dentistastore.DentistaSqlStore{DB: db}
	dentistaRepo := dentista.DentistaRepository{Store: &dentistaStorage}
	dentistaService := dentista.DentistaService{Repository: &dentistaRepo}
	dentistaHandler := handlers.DentistaHandler{DentistaService: &dentistaService}

	pacienteStorage := pacientestore.PacienteSqlStore{DB: db}
	pacienteRepo := paciente.PacienteRepository{Store: &pacienteStorage}
	pacienteService := paciente.PacienteService{Repository: &pacienteRepo}
	pacienteHandler := handlers.PacienteHandler{PacienteService: &pacienteService}

	turnoStorage := turnostore.TurnoSqlStore{DB: db, PacienteStore: &pacienteStorage, DentistaStore: &dentistaStorage}
	turnoRepo := turno.TurnoRepository{Store: &turnoStorage}
	turnoService := turno.TurnoService{Repository: &turnoRepo}
	turnoHandler := handlers.TurnoHandler{TurnoService: &turnoService}

	r := gin.Default()

	dentistas := r.Group("dentistas")
	{
		dentistas.POST("", dentistaHandler.Post)
		dentistas.GET("", dentistaHandler.GetAll)
		dentistas.GET(":id", dentistaHandler.GetById)
		dentistas.PATCH(":id", dentistaHandler.Patch)
		dentistas.PUT(":id", dentistaHandler.Put)
		dentistas.DELETE(":id", dentistaHandler.Delete)

	}

	pacientes := r.Group("pacientes")
	{
		pacientes.POST("", pacienteHandler.Post)
		pacientes.GET("", pacienteHandler.GetAll)
		pacientes.GET(":id", pacienteHandler.GetById)
		pacientes.PATCH(":id", pacienteHandler.Patch)
		pacientes.PUT(":id", pacienteHandler.Put)
		pacientes.DELETE(":id", pacienteHandler.Delete)
	}

	turnos := r.Group("turnos")
	{
		turnos.POST("", turnoHandler.Post)
		turnos.GET("", turnoHandler.GetAll)
		turnos.GET(":id", turnoHandler.GetById)
		turnos.PATCH(":id", turnoHandler.Patch)
		turnos.PUT(":id", turnoHandler.Put)
		turnos.DELETE(":id", turnoHandler.Delete)
		turnos.POST("/byDniAndLicense", turnoHandler.PostByPatientDniAndDentistLicense)
		turnos.GET("/byDniPatient", turnoHandler.GetByPatientDni)
	}

	r.Run(":8080")

}
