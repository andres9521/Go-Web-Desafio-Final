package turnoStore

import (
	"database/sql"

	"github.com/andres9521/Go-Web-Desafio-Final/internal/domain"
	dentistastore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/dentistaStore"
	pacientestore "github.com/andres9521/Go-Web-Desafio-Final/pkg/stores/pacienteStore"
)

type TurnoSqlStore struct {
	DB            *sql.DB
	PacienteStore pacientestore.PacienteStoreInterface
	DentistaStore dentistastore.DentistaStoreInterface
}

func (s *TurnoSqlStore) Create(t domain.Turno) (*domain.Turno, error) {

	query := "INSERT INTO turnos (descripcion, pacienteId, dentistaId) VALUES (?,?,?)"
	result, err := s.DB.Exec(query, t.Descripcion, t.PacienteId, t.DentistaId)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	turno := &domain.Turno{
		Id:          int(id),
		Descripcion: t.Descripcion,
		PacienteId:  t.PacienteId,
		DentistaId:  t.DentistaId,
	}

	return turno, nil

}

func (s *TurnoSqlStore) GetById(id int) (*domain.Turno, error) {

	var turno domain.Turno

	query := "SELECT * FROM turnos WHERE id = ?;"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&turno.Id, &turno.Descripcion, &turno.PacienteId, &turno.DentistaId)
	if err != nil {
		return nil, err
	}

	return &turno, nil

}

func (s *TurnoSqlStore) GetAll() ([]domain.Turno, error) {

	var turnos []domain.Turno

	query := "SELECT * FROM turnos;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var turno domain.Turno

		err := rows.Scan(&turno.Id, &turno.Descripcion, &turno.PacienteId, &turno.DentistaId)
		if err != nil {
			return nil, err
		}

		turnos = append(turnos, turno)

	}

	return turnos, nil

}

func (s *TurnoSqlStore) UpdateOne(id int, t domain.Turno) (*domain.Turno, error) {

	query := "UPDATE turnos SET descripcion = ?, pacienteId = ?, dentistaId = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, t.Descripcion, t.PacienteId, t.DentistaId, id)
	if err != nil {
		return nil, err
	}

	turno := &domain.Turno{
		Id:          id,
		Descripcion: t.Descripcion,
		PacienteId:  t.PacienteId,
		DentistaId:  t.DentistaId,
	}

	return turno, nil

}

func (s *TurnoSqlStore) UpdateMany(id int, t domain.Turno) (*domain.Turno, error) {

	query := "UPDATE turnos SET descripcion = ?, pacienteId = ?, dentistaId = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, t.Descripcion, t.PacienteId, t.DentistaId, id)
	if err != nil {
		return nil, err
	}

	turno := &domain.Turno{
		Id:          id,
		Descripcion: t.Descripcion,
		PacienteId:  t.PacienteId,
		DentistaId:  t.DentistaId,
	}

	return turno, nil

}

func (s *TurnoSqlStore) Delete(id int) error {

	query := "DELETE FROM turnos WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *TurnoSqlStore) CreateByPatientDniAndDentistLicense(t domain.Turno, dni string, matricula string) (*domain.Turno, error) {

	idPaciente, err := s.PacienteStore.GetIdByDni(dni)
	if err != nil {
		return nil, err
	}

	idDentista, err := s.DentistaStore.GetIdByLicense(matricula)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO turnos (descripcion, pacienteId, dentistaId) VALUES (?,?,?);"
	result, err := s.DB.Exec(query, t.Descripcion, idPaciente, idDentista)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	turno := &domain.Turno{
		Id:          int(id),
		Descripcion: t.Descripcion,
		PacienteId:  t.PacienteId,
		DentistaId:  t.DentistaId,
	}

	return turno, nil

}

func (s *TurnoSqlStore) GetByPatientDni(dni string) (*domain.Turno, error) {

	var turno domain.Turno

	idPaciente, err := s.PacienteStore.GetIdByDni(dni)
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM turnos WHERE pacienteId = ?"
	row := s.DB.QueryRow(query, idPaciente)

	err = row.Scan(&turno.Id, &turno.Descripcion, &turno.PacienteId, &turno.DentistaId)
	if err != nil {
		return nil, err
	}

	return &turno, nil

}
