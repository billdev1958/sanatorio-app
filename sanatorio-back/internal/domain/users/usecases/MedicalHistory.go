package usecase

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	"time"
)

func (u *usecase) GetMedicalHistoryByID(ctx context.Context, md models.MedicalHistoryRequest) (models.MedicalHistoryResponse, error) {
	medicalHistoryEntity, err := u.repo.GetMedicalHistoryByID(ctx, md.MedicalHistoryID)
	if err != nil {
		return models.MedicalHistoryResponse{}, fmt.Errorf("failed to fetch medical history: %w", err)
	}

	medicalHistoryData := models.MedicalHistoryResponse{
		MedicalHistoryID:         medicalHistoryEntity.MedicalHistoryID,
		DateOfRecord:             medicalHistoryEntity.DateOfRecord,
		TimeOfRecord:             medicalHistoryEntity.TimeOfRecord,
		PatientName:              medicalHistoryEntity.PatientName,
		Curp:                     medicalHistoryEntity.Curp,
		BirthDate:                medicalHistoryEntity.BirthDate,
		Age:                      medicalHistoryEntity.Age,
		Gender:                   medicalHistoryEntity.Gender,
		PlaceOfOrigin:            medicalHistoryEntity.PlaceOfOrigin,
		EthnicGroup:              medicalHistoryEntity.EthnicGroup,
		PhoneNumber:              medicalHistoryEntity.PhoneNumber,
		Address:                  medicalHistoryEntity.Address,
		Occupation:               medicalHistoryEntity.Occupation,
		GuardianName:             medicalHistoryEntity.GuardianName,
		FamilyMedicalHistory:     medicalHistoryEntity.FamilyMedicalHistory,
		NonPathologicalHistory:   medicalHistoryEntity.NonPathologicalHistory,
		PathologicalHistory:      medicalHistoryEntity.PathologicalHistory,
		GynecObstetricHistory:    medicalHistoryEntity.GynecObstetricHistory,
		CurrentCondition:         medicalHistoryEntity.CurrentCondition,
		Cardiovascular:           medicalHistoryEntity.Cardiovascular,
		Respiratory:              medicalHistoryEntity.Respiratory,
		Gastrointestinal:         medicalHistoryEntity.Gastrointestinal,
		Genitourinary:            medicalHistoryEntity.Genitourinary,
		HematicLymphatic:         medicalHistoryEntity.HematicLymphatic,
		Endocrine:                medicalHistoryEntity.Endocrine,
		NervousSystem:            medicalHistoryEntity.NervousSystem,
		Musculoskeletal:          medicalHistoryEntity.Musculoskeletal,
		Skin:                     medicalHistoryEntity.Skin,
		BodyTemperature:          medicalHistoryEntity.BodyTemperature,
		Weight:                   medicalHistoryEntity.Weight,
		Height:                   medicalHistoryEntity.Height,
		BMI:                      medicalHistoryEntity.BMI,
		HeartRate:                medicalHistoryEntity.HeartRate,
		RespiratoryRate:          medicalHistoryEntity.RespiratoryRate,
		BloodPressure:            medicalHistoryEntity.BloodPressure,
		Physical:                 medicalHistoryEntity.Physical,
		Head:                     medicalHistoryEntity.Head,
		NeckAndChest:             medicalHistoryEntity.NeckAndChest,
		Abdomen:                  medicalHistoryEntity.Abdomen,
		Genital:                  medicalHistoryEntity.Genital,
		Extremities:              medicalHistoryEntity.Extremities,
		PreviousResults:          medicalHistoryEntity.PreviousResults,
		Diagnoses:                medicalHistoryEntity.Diagnoses,
		PharmacologicalTreatment: medicalHistoryEntity.PharmacologicalTreatment,
		Prognosis:                medicalHistoryEntity.Prognosis,
		DoctorName:               medicalHistoryEntity.DoctorName,
		MedicalLicense:           medicalHistoryEntity.MedicalLicense,
		SpecialtyLicense:         medicalHistoryEntity.SpecialtyLicense,
		Status:                   medicalHistoryEntity.Status,
	}

	return medicalHistoryData, nil

}

func (u *usecase) CompleteMedicalHistory(ctx context.Context, request models.CompleteMedicalHistoryRequest) (string, error) {

	// Obtener la fecha y hora actuales
	updated_at := time.Now()
	dateRecord := time.Now()

	// Crear un puntero de tipo time.Time para la hora actual
	timeOnly := time.Date(0, 1, 1, dateRecord.Hour(), dateRecord.Minute(), dateRecord.Second(), 0, time.UTC)

	// Crear la instancia de MedicalHistory con las correcciones necesarias
	medicalHistory := entities.MedicalHistory{
		DateOfRecord:             &dateRecord,
		TimeOfRecord:             &timeOnly,
		PlaceOfOrigin:            request.PlaceOfOrigin,
		EthnicGroup:              request.EthnicGroup,
		PhoneNumber:              request.PhoneNumber,
		Address:                  request.Address,
		Occupation:               request.Occupation,
		GuardianName:             request.GuardianName,
		FamilyMedicalHistory:     request.FamilyMedicalHistory,
		NonPathologicalHistory:   request.NonPathologicalHistory,
		PathologicalHistory:      request.PathologicalHistory,
		GynecObstetricHistory:    request.GynecObstetricHistory,
		CurrentCondition:         request.CurrentCondition,
		Cardiovascular:           request.Cardiovascular,
		Respiratory:              request.Respiratory,
		Gastrointestinal:         request.Gastrointestinal,
		Genitourinary:            request.Genitourinary,
		HematicLymphatic:         request.HematicLymphatic,
		Endocrine:                request.Endocrine,
		NervousSystem:            request.NervousSystem,
		Musculoskeletal:          request.Musculoskeletal,
		Skin:                     request.Skin,
		BodyTemperature:          request.BodyTemperature,
		Weight:                   request.Weight,
		Height:                   request.Height,
		BMI:                      request.BMI,
		HeartRate:                request.HeartRate,
		RespiratoryRate:          request.RespiratoryRate,
		BloodPressure:            request.BloodPressure,
		Physical:                 request.Physical,
		Head:                     request.Head,
		NeckAndChest:             request.NeckAndChest,
		Abdomen:                  request.Abdomen,
		Genital:                  request.Genital,
		Extremities:              request.Extremities,
		PreviousResults:          request.PreviousResults,
		Diagnoses:                request.Diagnoses,
		PharmacologicalTreatment: request.PharmacologicalTreatment,
		Prognosis:                request.Prognosis,
		DoctorName:               request.DoctorName,
		MedicalLicense:           request.MedicalLicense,
		SpecialtyLicense:         request.SpecialtyLicense,
		Status:                   true,
		Updated_At:               updated_at,
		MedicalHistoryID:         request.MedicalHistoryID,
	}

	// Llamar al repositorio para actualizar el registro
	success, err := u.repo.CompleteMedicalHistory(ctx, medicalHistory)
	if err != nil {
		log.Printf("Error updating medical history: %v", err)
		return "", err
	}

	// Validar el éxito de la operación
	if !success {
		log.Println("Failed to update medical history: operation returned false")
		return "", fmt.Errorf("failed to update medical history")
	}

	// Retornar un mensaje de éxito
	return "Medical history updated successfully", nil
}
