package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) GetMedicalHistoryByID(ctx context.Context, medicalHistoryID string) (entities.MedicalHistory, error) {
	query := `
SELECT 
    medical_history_id,
    date_of_record,
    time_of_record,
    patient_name,
    curp,
    birth_date,
    COALESCE(age, '') AS age,
    gender,
    COALESCE(place_of_origin, '') AS place_of_origin,
    COALESCE(ethnic_group, '') AS ethnic_group,
    COALESCE(phone_number, '') AS phone_number,
    COALESCE(address, '') AS address,
    COALESCE(occupation, '') AS occupation,
    COALESCE(guardian_name, '') AS guardian_name,
    COALESCE(family_medical_history, '') AS family_medical_history,
    COALESCE(non_pathological_history, '') AS non_pathological_history,
    COALESCE(pathological_history, '') AS pathological_history,
    COALESCE(gynec_obstetric_history, '') AS gynec_obstetric_history,
    COALESCE(current_condition, '') AS current_condition,
    COALESCE(cardiovascular, '') AS cardiovascular,
    COALESCE(respiratory, '') AS respiratory,
    COALESCE(gastrointestinal, '') AS gastrointestinal,
    COALESCE(genitourinary, '') AS genitourinary,
    COALESCE(hematic_lymphatic, '') AS hematic_lymphatic,
    COALESCE(endocrine, '') AS endocrine,
    COALESCE(nervous_system, '') AS nervous_system,
    COALESCE(musculoskeletal, '') AS musculoskeletal,
    COALESCE(skin, '') AS skin,
    COALESCE(body_temperature, '') AS body_temperature,
    COALESCE(weight, '') AS weight,
    COALESCE(height, '') AS height,
    COALESCE(bmi, '') AS bmi,
    COALESCE(heart_rate, '') AS heart_rate,
    COALESCE(respiratory_rate, '') AS respiratory_rate,
    COALESCE(blood_pressure, '') AS blood_pressure,
    COALESCE(physical, '') AS physical,
    COALESCE(head, '') AS head,
    COALESCE(neck_and_chest, '') AS neck_and_chest,
    COALESCE(abdomen, '') AS abdomen,
    COALESCE(genital, '') AS genital,
    COALESCE(extremities, '') AS extremities,
    COALESCE(previous_results, '') AS previous_results,
    COALESCE(diagnoses, '') AS diagnoses,
    COALESCE(pharmacological_treatment, '') AS pharmacological_treatment,
    COALESCE(prognosis, '') AS prognosis,
    COALESCE(doctor_name, '') AS doctor_name,
    COALESCE(medical_license, '') AS medical_license,
    COALESCE(specialty_license, '') AS specialty_license,
    status_md
FROM medical_history
WHERE medical_history_id = $1


`

	var medicalHistory entities.MedicalHistory

	// Ejecutar la consulta y escanear los resultados
	err := ur.storage.DbPool.QueryRow(ctx, query, medicalHistoryID).Scan(
		&medicalHistory.MedicalHistoryID,
		&medicalHistory.DateOfRecord,
		&medicalHistory.TimeOfRecord,
		&medicalHistory.PatientName,
		&medicalHistory.Curp,
		&medicalHistory.BirthDate,
		&medicalHistory.Age,
		&medicalHistory.Gender,
		&medicalHistory.PlaceOfOrigin,
		&medicalHistory.EthnicGroup,
		&medicalHistory.PhoneNumber,
		&medicalHistory.Address,
		&medicalHistory.Occupation,
		&medicalHistory.GuardianName,
		&medicalHistory.FamilyMedicalHistory,
		&medicalHistory.NonPathologicalHistory,
		&medicalHistory.PathologicalHistory,
		&medicalHistory.GynecObstetricHistory,
		&medicalHistory.CurrentCondition,
		&medicalHistory.Cardiovascular,
		&medicalHistory.Respiratory,
		&medicalHistory.Gastrointestinal,
		&medicalHistory.Genitourinary,
		&medicalHistory.HematicLymphatic,
		&medicalHistory.Endocrine,
		&medicalHistory.NervousSystem,
		&medicalHistory.Musculoskeletal,
		&medicalHistory.Skin,
		&medicalHistory.BodyTemperature,
		&medicalHistory.Weight,
		&medicalHistory.Height,
		&medicalHistory.BMI,
		&medicalHistory.HeartRate,
		&medicalHistory.RespiratoryRate,
		&medicalHistory.BloodPressure,
		&medicalHistory.Physical,
		&medicalHistory.Head,
		&medicalHistory.NeckAndChest,
		&medicalHistory.Abdomen,
		&medicalHistory.Genital,
		&medicalHistory.Extremities,
		&medicalHistory.PreviousResults,
		&medicalHistory.Diagnoses,
		&medicalHistory.PharmacologicalTreatment,
		&medicalHistory.Prognosis,
		&medicalHistory.DoctorName,
		&medicalHistory.MedicalLicense,
		&medicalHistory.SpecialtyLicense,
		&medicalHistory.Status,
	)

	// Manejo de errores
	if err != nil {
		if err == pgx.ErrNoRows {
			return medicalHistory, fmt.Errorf("medical history with ID %v not found", medicalHistoryID)
		}
		return medicalHistory, fmt.Errorf("failed to retrieve medical history by ID: %w", err)
	}

	return medicalHistory, nil
}

func (ur *userRepository) CompleteMedicalHistory(ctx context.Context, md entities.MedicalHistory) (bool, error) {
	query := `
        UPDATE medical_history
        SET 
            date_of_record = $1,
            time_of_record = $2,
            place_of_origin = $3,
            ethnic_group = $4,
            phone_number = $5,
            address = $6,
            occupation = $7,
            guardian_name = $8,
            family_medical_history = $9,
            non_pathological_history = $10,
            pathological_history = $11,
            gynec_obstetric_history = $12,
            current_condition = $13,
            cardiovascular = $14,
            respiratory = $15,
            gastrointestinal = $16,
            genitourinary = $17,
            hematic_lymphatic = $18,
            endocrine = $19,
            nervous_system = $20,
            musculoskeletal = $21,
            skin = $22,
            body_temperature = $23,
            weight = $24,
            height = $25,
            bmi = $26,
            heart_rate = $27,
            respiratory_rate = $28,
            blood_pressure = $29,
            physical = $30,
            head = $31,
            neck_and_chest = $32,
            abdomen = $33,
            genital = $34,
            extremities = $35,
            previous_results = $36,
            diagnoses = $37,
            pharmacological_treatment = $38,
            prognosis = $39,
            doctor_name = $40,
            medical_license = $41,
            specialty_license = $42,
            status_md = $43,
            updated_at = $44
        WHERE medical_history_id = $45
    `

	_, err := ur.storage.DbPool.Exec(
		ctx,
		query,
		md.DateOfRecord,
		md.TimeOfRecord,
		md.PlaceOfOrigin,
		md.EthnicGroup,
		md.PhoneNumber,
		md.Address,
		md.Occupation,
		md.GuardianName,
		md.FamilyMedicalHistory,
		md.NonPathologicalHistory,
		md.PathologicalHistory,
		md.GynecObstetricHistory,
		md.CurrentCondition,
		md.Cardiovascular,
		md.Respiratory,
		md.Gastrointestinal,
		md.Genitourinary,
		md.HematicLymphatic,
		md.Endocrine,
		md.NervousSystem,
		md.Musculoskeletal,
		md.Skin,
		md.BodyTemperature,
		md.Weight,
		md.Height,
		md.BMI,
		md.HeartRate,
		md.RespiratoryRate,
		md.BloodPressure,
		md.Physical,
		md.Head,
		md.NeckAndChest,
		md.Abdomen,
		md.Genital,
		md.Extremities,
		md.PreviousResults,
		md.Diagnoses,
		md.PharmacologicalTreatment,
		md.Prognosis,
		md.DoctorName,
		md.MedicalLicense,
		md.SpecialtyLicense,
		md.Status,
		md.Updated_At,
		md.MedicalHistoryID,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}
