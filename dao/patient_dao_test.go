package dao

import (
	"testing"
	"github.com/reactmed/goneurdicom/utils"
	"github.com/reactmed/goneurdicom/domain"
	"sort"
	"database/sql"
	_ "github.com/lib/pq"
	"reflect"
)

func CreatePatientDao() PatientDao {
	connStr := "user=neurdicom password=neurdicom dbname=neurdicom_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Can not connect to db")
	}
	ctx := utils.GetAppContext().Bind(reflect.TypeOf((*sql.DB)(nil)).Elem(), db, utils.Singleton)
	db.Exec(
		`
	TRUNCATE patients CASCADE;
	INSERT INTO patients (id, patient_name, patient_sex, patient_birthdate, patient_age, patient_id) 
	VALUES (1, 'Patient Name 1', 'M', '1979-01-01', '30Y', '1'),
	(2, 'Patient Name 2', 'F', '1983-01-01', '31Y', '2');
	SELECT setval('patients_id_seq'::REGCLASS, 3);
	`)
	dao := NewPatientDao(ctx)
	return dao
}

func TestPatientDao_Save(t *testing.T) {
	dao := CreatePatientDao()
	patient := domain.Patient{
		PatientId:        "3",
		PatientName:      "Patient Name 3",
		PatientSex:       "M",
		PatientBirthdate: "2000-01-01",
		PatientAge:       "18Y",
	}
	patient, err := dao.Save(patient)
	utils.AssertNotErr(err, t)
	if patient, err := dao.FindOne(patient.Id); err != nil || patient == nil {
		t.Error("User not found")
	}

	utils.AssertEqual(patient.PatientId, "3", t)
	utils.AssertEqual(patient.PatientName, "Patient Name 3", t)
	utils.AssertEqual(patient.PatientSex, "M", t)
	utils.AssertEqual(patient.PatientBirthdate, "2000-01-01", t)
	utils.AssertEqual(patient.PatientAge, "18Y", t)
}

func TestPatientDao_Update(t *testing.T) {
	patient := domain.Patient{
		PatientId:        "4",
		PatientName:      "Patient Name 4",
		PatientSex:       "M",
		PatientBirthdate: "1990-01-01",
		PatientAge:       "28Y",
	}
	dao := CreatePatientDao()
	patient, err := dao.Update(patient)
	utils.AssertNotErr(err, t)
	patient2, err := dao.FindOne(1)
	utils.AssertNotNil(patient2, t)

	utils.AssertEqual(patient.PatientId, "4", t)
	utils.AssertEqual(patient.PatientName, "Patient Name 4", t)
	utils.AssertEqual(patient.PatientSex, "M", t)
	utils.AssertEqual(patient.PatientBirthdate, "1990-01-01", t)
	utils.AssertEqual(patient.PatientAge, "28Y", t)
}

func TestPatientDao_Delete(t *testing.T) {
	dao := CreatePatientDao()
	dao.Delete(1)
	res, _ := dao.Exists(1)
	utils.AssertFalsef(res, t, "Patient should be removed")
}

func TestPatientDao_FindOne(t *testing.T) {
	dao := CreatePatientDao()
	patient, err := dao.FindOne(1)
	utils.AssertNilf(err, t, "Can not find patient")
	utils.AssertNotNilf(patient, t, "Patient not found")

	utils.AssertEqual(patient.PatientId, "1", t)
	utils.AssertEqual(patient.PatientName, "Patient Name 1", t)
	utils.AssertEqual(patient.PatientSex, "M", t)
	utils.AssertEqual(patient.PatientBirthdate, "1979-01-01", t)
	utils.AssertEqual(patient.PatientAge, "30Y", t)
}

func TestPatientDao_FindAll(t *testing.T) {
	dao := CreatePatientDao()
	patients, _ := dao.FindAll()
	patients2 := make([]interface{}, len(patients))
	for i, v := range patients {
		patients2[i] = v
	}
	utils.AssertHasLen(patients2, 2, t)
	sort.Slice(patients, func(i, j int) bool {
		return patients[i].Id < patients[j].Id
	})

	patient := patients[0]
	utils.AssertEqual(patient.PatientId, "1", t)
	utils.AssertEqual(patient.PatientName, "Patient Name 1", t)
	utils.AssertEqual(patient.PatientSex, "M", t)
	utils.AssertEqual(patient.PatientBirthdate, "1979-01-01", t)
	utils.AssertEqual(patient.PatientAge, "30Y", t)

	patient = patients[1]
	utils.AssertEqual(patient.PatientId, "2", t)
	utils.AssertEqual(patient.PatientName, "Patient Name 2", t)
	utils.AssertEqual(patient.PatientSex, "F", t)
	utils.AssertEqual(patient.PatientBirthdate, "1983-01-01", t)
	utils.AssertEqual(patient.PatientAge, "31Y", t)
}

func TestPatientDao_Exists(t *testing.T) {
	dao := CreatePatientDao()
	res, _ := dao.Exists(1)
	utils.AssertTruef(res, t, "Patient should exist")
}

func TestPatientDao_Count(t *testing.T) {
	dao := CreatePatientDao()
	c, _ := dao.Count()
	utils.AssertEqual(c, 2, t)
}
