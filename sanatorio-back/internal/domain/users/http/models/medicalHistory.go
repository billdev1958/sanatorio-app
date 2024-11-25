package models

import (
	"time"
)

type MedicalHistoryRequest struct {
	MedicalHistoryID string `json:"medical_history_id"` // Unique identifier
}

type CompleteMedicalHistoryRequest struct {
	DateOfRecord             time.Time `json:"date_of_record"`            // Fecha de registro
	TimeOfRecord             string    `json:"time_of_record"`            // Hora de registro
	PlaceOfOrigin            string    `json:"place_of_origin"`           // Lugar de procedencia
	EthnicGroup              string    `json:"ethnic_group"`              // Grupo étnico
	PhoneNumber              string    `json:"phone_number"`              // Teléfono
	Address                  string    `json:"address"`                   // Domicilio
	Occupation               string    `json:"occupation"`                // Ocupación
	GuardianName             string    `json:"guardian_name"`             // Nombre del tutor
	FamilyMedicalHistory     string    `json:"family_medical_history"`    // Antecedentes médicos familiares
	NonPathologicalHistory   string    `json:"non_pathological_history"`  // Antecedentes no patológicos
	PathologicalHistory      string    `json:"pathological_history"`      // Antecedentes patológicos
	GynecObstetricHistory    string    `json:"gynec_obstetric_history"`   // Antecedentes gineco-obstétricos
	CurrentCondition         string    `json:"current_condition"`         // Condición actual
	Cardiovascular           string    `json:"cardiovascular"`            // Sistema cardiovascular
	Respiratory              string    `json:"respiratory"`               // Sistema respiratorio
	Gastrointestinal         string    `json:"gastrointestinal"`          // Sistema gastrointestinal
	Genitourinary            string    `json:"genitourinary"`             // Sistema genitourinario
	HematicLymphatic         string    `json:"hematic_lymphatic"`         // Sistema hemático y linfático
	Endocrine                string    `json:"endocrine"`                 // Sistema endocrino
	NervousSystem            string    `json:"nervous_system"`            // Sistema nervioso
	Musculoskeletal          string    `json:"musculoskeletal"`           // Sistema musculoesquelético
	Skin                     string    `json:"skin"`                      // Piel
	BodyTemperature          string    `json:"body_temperature"`          // Temperatura corporal
	Weight                   string    `json:"weight"`                    // Peso
	Height                   string    `json:"height"`                    // Altura
	BMI                      string    `json:"bmi"`                       // Índice de masa corporal (IMC)
	HeartRate                string    `json:"heart_rate"`                // Frecuencia cardíaca
	RespiratoryRate          string    `json:"respiratory_rate"`          // Frecuencia respiratoria
	BloodPressure            string    `json:"blood_pressure"`            // Presión arterial
	Physical                 string    `json:"physical"`                  // Examen físico
	Head                     string    `json:"head"`                      // Cabeza
	NeckAndChest             string    `json:"neck_and_chest"`            // Cuello y tórax
	Abdomen                  string    `json:"abdomen"`                   // Abdomen
	Genital                  string    `json:"genital"`                   // Genitales
	Extremities              string    `json:"extremities"`               // Extremidades
	PreviousResults          string    `json:"previous_results"`          // Resultados anteriores
	Diagnoses                string    `json:"diagnoses"`                 // Diagnósticos
	PharmacologicalTreatment string    `json:"pharmacological_treatment"` // Tratamiento farmacológico
	Prognosis                string    `json:"prognosis"`                 // Pronóstico
	DoctorName               string    `json:"doctor_name"`               // Nombre del médico
	MedicalLicense           string    `json:"medical_license"`           // Cédula médica
	SpecialtyLicense         string    `json:"specialty_license"`         // Cédula de especialidad
	Status                   bool      `json:"status"`                    // Estado completo o incompleto
	Updated_At               time.Time `json:"updated_at"`                // Última actualización
	MedicalHistoryID         string    `json:"medical_history_id"`        // Identificador único de historial médico
}

type MedicalHistoryResponse struct {
	MedicalHistoryID         string    `json:"medical_history_id"`        // Unique identifier
	DateOfRecord             time.Time `json:"date_of_record"`            // Fecha de registro
	TimeOfRecord             string    `json:"time_of_record"`            // Hora de registro (formato HH:MM:SS)
	PatientName              string    `json:"patient_name"`              // Nombre del paciente
	Curp                     string    `json:"curp"`                      // CURP
	BirthDate                time.Time `json:"birth_date"`                // Fecha de nacimiento
	Age                      string    `json:"age"`                       // Edad
	Gender                   string    `json:"gender"`                    // Género
	PlaceOfOrigin            string    `json:"place_of_origin"`           // Lugar de procedencia
	EthnicGroup              string    `json:"ethnic_group"`              // Grupo étnico
	PhoneNumber              string    `json:"phone_number"`              // Teléfono
	Address                  string    `json:"address"`                   // Domicilio
	Occupation               string    `json:"occupation"`                // Ocupación
	GuardianName             string    `json:"guardian_name"`             // Nombre del tutor
	FamilyMedicalHistory     string    `json:"family_medical_history"`    // Antecedentes médicos familiares
	NonPathologicalHistory   string    `json:"non_pathological_history"`  // Antecedentes no patológicos
	PathologicalHistory      string    `json:"pathological_history"`      // Antecedentes patológicos
	GynecObstetricHistory    string    `json:"gynec_obstetric_history"`   // Antecedentes gineco-obstétricos
	CurrentCondition         string    `json:"current_condition"`         // Condición actual
	Cardiovascular           string    `json:"cardiovascular"`            // Sistema cardiovascular
	Respiratory              string    `json:"respiratory"`               // Sistema respiratorio
	Gastrointestinal         string    `json:"gastrointestinal"`          // Sistema gastrointestinal
	Genitourinary            string    `json:"genitourinary"`             // Sistema genitourinario
	HematicLymphatic         string    `json:"hematic_lymphatic"`         // Sistema hemático y linfático
	Endocrine                string    `json:"endocrine"`                 // Sistema endocrino
	NervousSystem            string    `json:"nervous_system"`            // Sistema nervioso
	Musculoskeletal          string    `json:"musculoskeletal"`           // Sistema musculoesquelético
	Skin                     string    `json:"skin"`                      // Piel
	BodyTemperature          string    `json:"body_temperature"`          // Temperatura corporal
	Weight                   string    `json:"weight"`                    // Peso
	Height                   string    `json:"height"`                    // Altura
	BMI                      string    `json:"bmi"`                       // Índice de masa corporal (IMC)
	HeartRate                string    `json:"heart_rate"`                // Frecuencia cardíaca
	RespiratoryRate          string    `json:"respiratory_rate"`          // Frecuencia respiratoria
	BloodPressure            string    `json:"blood_pressure"`            // Presión arterial
	Physical                 string    `json:"physical"`                  // Examen físico
	Head                     string    `json:"head"`                      // Cabeza
	NeckAndChest             string    `json:"neck_and_chest"`            // Cuello y tórax
	Abdomen                  string    `json:"abdomen"`                   // Abdomen
	Genital                  string    `json:"genital"`                   // Genitales
	Extremities              string    `json:"extremities"`               // Extremidades
	PreviousResults          string    `json:"previous_results"`          // Resultados anteriores
	Diagnoses                string    `json:"diagnoses"`                 // Diagnósticos
	PharmacologicalTreatment string    `json:"pharmacological_treatment"` // Tratamiento farmacológico
	Prognosis                string    `json:"prognosis"`                 // Pronóstico
	DoctorName               string    `json:"doctor_name"`               // Nombre del médico
	MedicalLicense           string    `json:"medical_license"`           // Cédula médica
	SpecialtyLicense         string    `json:"specialty_license"`         // Cédula de especialidad
	Status                   bool      `json:"status"`                    // Estado activo o inactivo
}
