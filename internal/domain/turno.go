package domain

type Turno struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	PacienteId  int    `json:"pacienteId"`
	DentistaId  int    `json:"dentistaId"`
}
