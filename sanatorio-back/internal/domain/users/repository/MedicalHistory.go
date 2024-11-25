package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) GetMedicalHistoryByID(ctx context.Context, MedicalHistoryID string) (entities.MedicalHistory, error) {
	query := "SELECT * FROM medical_history WHERE medical_history_id = ?"

	var medicalHistory entities.MedicalHistory

	err := ur.storage.DbPool.QueryRow(ctx, query, MedicalHistoryID).Scan(
		&medicalHistory.ID,
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

	if err != nil {
		if err == pgx.ErrNoRows {
			return medicalHistory, fmt.Errorf("doctor with ID %v not found", MedicalHistoryID)
		}
		return medicalHistory, fmt.Errorf("failed to get doctor by ID: %w", err)
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
            status = $43,
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
