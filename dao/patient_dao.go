package dao

import (
	"log"
	"errors"
	"github.com/reactmed/goneurdicom/utils"
	"github.com/reactmed/goneurdicom/domain"
	"database/sql"
	"reflect"
)

var patientDaoInstance PatientDao

func NewPatientDao(ctx utils.AppContext) PatientDao {
	if patientDaoInstance == nil {
		patientDaoInstance = &patientDao{
			ctx: ctx,
		}
	}
	return patientDaoInstance
}

type PatientDao interface {
	Save(user domain.Patient) (domain.Patient, error)
	Update(user domain.Patient) (domain.Patient, error)
	Delete(id int) error
	FindOne(id int) (*domain.Patient, error)
	FindAll() ([]domain.Patient, error)
	Exists(id int) (bool, error)
	Count() (int, error)
}

type patientDao struct {
	ctx utils.AppContext
}

func (dao *patientDao) getDb() (*sql.DB, error) {
	db, err := dao.ctx.Get(reflect.TypeOf((*sql.DB)(nil)).Elem())
	return (db).(*sql.DB), err
}

func (dao *patientDao) Save(patient domain.Patient) (domain.Patient, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return patient, err
	}
	stmt, err := db.Prepare(`
	INSERT INTO patients (patient_id, patient_name, patient_sex, patient_birthdate, patient_age) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`)
	if err != nil {
		log.Fatal(err)
		return patient, err
	}
	err = stmt.QueryRow(patient.PatientId, patient.PatientName, patient.PatientSex,
		patient.PatientBirthdate, patient.PatientAge).Scan(&patient.Id)
	if err != nil {
		log.Fatal(err)
		return patient, err
	}
	return patient, nil
}

func (dao *patientDao) Update(patient domain.Patient) (domain.Patient, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return patient, err
	}
	stmt, err := db.Prepare(`UPDATE patients SET patient_id = $1, patient_name = $2, patient_sex = $3, 
	patient_birthdate = $4, patient_age = $5 WHERE id = $6
	`)
	if err != nil {
		log.Fatal(err)
		return patient, err
	}
	_, err = stmt.Exec(patient.PatientId, patient.PatientName, patient.PatientSex,
		patient.PatientBirthdate, patient.PatientAge, patient.Id)
	if err != nil {
		log.Fatal(err)
		return patient, err
	}
	return patient, nil
}

func (dao *patientDao) Delete(id int) error {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := db.Prepare("DELETE FROM patients WHERE id = $1")
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (dao *patientDao) FindOne(id int) (*domain.Patient, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stmt, err := db.Prepare(`
	SELECT patient_id, patient_name, patient_sex, patient_birthdate, patient_age FROM patients
	WHERE id = $1
	`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	patient := &domain.Patient{}
	err = stmt.QueryRow(id).Scan(
		&patient.PatientId, &patient.PatientName, &patient.PatientSex,
		&patient.PatientBirthdate, &patient.PatientAge,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return patient, nil
}

func (dao *patientDao) FindAll() ([]domain.Patient, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stmt, err := db.Prepare(`
	SELECT id, patient_id, patient_name, patient_sex, patient_birthdate, patient_age FROM patients
	`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	patients := make(domain.Patients, 0)
	for rows.Next() {
		patient := domain.Patient{}
		rows.Scan(
			&patient.Id, &patient.PatientId, &patient.PatientName, &patient.PatientSex,
			&patient.PatientBirthdate, &patient.PatientAge,
		)
		patients = append(patients, patient)
	}
	return patients, nil
}

func (dao *patientDao) Exists(id int) (bool, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	stmt, err := db.Prepare("SELECT EXISTS(SELECT * FROM patients WHERE id = $1)")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	var res bool
	err = stmt.QueryRow(id).Scan(&res)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return res, nil
}

func (dao *patientDao) Count() (int, error) {
	db, err := dao.getDb()
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	rows, err := db.Query("SELECT COUNT(*) FROM patients")
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var count int
		rows.Scan(&count)
		return count, nil
	}
	return -1, errors.New("empty rows")
}
