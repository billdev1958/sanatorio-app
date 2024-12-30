package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/entities"
	users "sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/pkg"
	"strings"

	"sanatorioApp/internal/domain/cites/http/models"
	"time"
)

type usecase struct {
	repo cites.CitesRepository
}

func NewUsecase(repo cites.CitesRepository) cites.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterOfficeSchedule(ctx context.Context, request models.RegisterOfficeSchedule) (string, error) {
	if len(request.SelectedDays) == 0 || len(request.TimeSlots) == 0 {
		return "", fmt.Errorf("selectedDays and timeSlots cannot be empty")
	}

	layout := "15:04"
	var officeSchedules []entities.OfficeSchedule

	for _, day := range request.SelectedDays {
		for _, slot := range request.TimeSlots {
			times := strings.Split(slot, " - ")
			if len(times) != 2 {
				return "", fmt.Errorf("invalid time slot format: %s", slot)
			}

			timeStart, err := time.Parse(layout, times[0])
			if err != nil {
				return "", fmt.Errorf("invalid time format for TimeStart: %w", err)
			}
			timeEnd, err := time.Parse(layout, times[1])
			if err != nil {
				return "", fmt.Errorf("invalid time format for TimeEnd: %w", err)
			}

			duration := timeEnd.Sub(timeStart)

			officeSchedules = append(officeSchedules, entities.OfficeSchedule{
				Office: entities.Office{
					ID: request.OfficeID,
				},
				ShiftID: request.ShiftID,
				Services: entities.Services{
					ID: request.ServiceID,
				},
				DoctorUser: users.DoctorUser{
					AccountID: request.DoctorID,
				},
				OfficeStatus: entities.OfficeStatus{
					ID: int(entities.OfficeStatusAvailable),
				},
				Schedule: entities.Schedule{
					DayOfWeek:    day,
					TimeStart:    timeStart,
					TimeEnd:      timeEnd,
					TimeDuration: duration,
				},
			})
		}
	}

	message, err := u.repo.RegisterOfficeSchedule(ctx, officeSchedules)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (u *usecase) RegisterAppointment(ctx context.Context, request models.RegisterAppointmentRequest) (string, error) {

	appointment := entities.Appointment{
		PatientAccountID: request.PatientAccountID,
		OfficeID:         request.OfficeID,
		Date:             request.Date,
		ScheduleID:       request.ScheduleID,
		StatusID:         int(entities.AppointmentStatusPendiente),
	}

	message, err := u.repo.RegisterAppointment(ctx, appointment)
	if err != nil {
		return "", fmt.Errorf("failed to register appointment: %w", err)
	}
	return message, nil
}

func (u *usecase) GetSchedules(ctx context.Context, filtersRequest models.OfficeSCheduleFiltersRequest) ([]models.OfficeScheduleResponse, error) {
	filters, err := pkg.Filter(filtersRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to process filters: %w", err)
	}

	schedules, err := u.repo.GetSchedules(ctx, filters)
	if err != nil {
		return nil, err
	}

	var responses []models.OfficeScheduleResponse

	for _, schedule := range schedules {
		response := models.OfficeScheduleResponse{
			OfficeScheduleID: schedule.ID,
			ServiceID:        schedule.Services.ID,
			ScheduleID:       schedule.Schedule.ID,
			OfficeID:         schedule.Office.ID,
			OfficeStatusID:   schedule.Services.ID,
			ShiftID:          schedule.ShiftID,
			DoctorID:         schedule.DoctorUser.AccountID,
			ServiceName:      schedule.Services.Name,
			DaySchedule:      schedule.Schedule.DayOfWeek,
			TimeStart:        schedule.Schedule.TimeStart.Format("15:04"),
			TimeEnd:          schedule.Schedule.TimeEnd.Format("15:04"),
			TimeDuration:     schedule.Schedule.TimeDuration.String(),
			OfficeName:       schedule.Office.Name,
			OfficeStatusName: schedule.OfficeStatus.Name,
			ShiftName:        schedule.ShiftName,
			DoctorName:       schedule.DoctorUser.FirstName,
			DoctorLastName1:  schedule.DoctorUser.LastName1,
			DoctorLastName2:  schedule.DoctorUser.LastName2,
			MedicalLicense:   schedule.DoctorUser.MedicalLicense,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
