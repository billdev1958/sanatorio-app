package entities

type Roles int

const (
	_ = iota
	SuperAdmin
	Admin
	Doctor
	Patient
)

type Permissions int

const (
	_ = iota
	CreateUsers
	ViewUsers
	EditUsers
	DeleteUsers

	CreateSchedule
	ViewSchedule
	EditSchedule
	DeleteSchedule

	CreateAppointment
	ViewAppointment
	EditAppointment
	DeleteAppointment

	CreateMedicalHistory
	ViewMedicalHistory
	EditMedicalHistory
	DeleteMedicalHistory
)

type RolePermissions int

const (
	SuperAdminPermissions = CreateUsers | ViewUsers | EditUsers | DeleteUsers | CreateSchedule | ViewSchedule | EditSchedule | DeleteSchedule | CreateAppointment | ViewAppointment | EditAppointment | DeleteAppointment | CreateMedicalHistory | ViewMedicalHistory | EditMedicalHistory | DeleteMedicalHistory

	AdminPermissions = CreateUsers | ViewUsers | EditUsers | CreateSchedule | ViewSchedule | EditSchedule | DeleteSchedule | CreateAppointment | ViewAppointment | EditAppointment | DeleteAppointment | CreateMedicalHistory | ViewMedicalHistory | EditMedicalHistory | DeleteMedicalHistory

	DoctorPermissions = ViewAppointment | EditAppointment | DeleteAppointment | CreateMedicalHistory | ViewMedicalHistory | EditMedicalHistory

	PatientPermissions = CreateAppointment | ViewAppointment | EditAppointment | DeleteAppointment | ViewMedicalHistory
)
