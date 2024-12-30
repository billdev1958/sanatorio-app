package user

type Permissions = int

const (
	_ = iota
	CreateSuperAdmin
	CreateAdmin
	CreateDoctor
	CreateReceptionist
	CreatePatient
	CreateBeneficiary
	ViewUsers
	EditSuperAdmin
	EditAdmin
	EditDoctor
	EditReceptionist
	EditPatient
	DeleteSuperAdmin
	DeleteAdmin
	DeleteDoctor
	DeleteReceptionist
	DeletePatient
	CreateSchedule
	ViewSchedule
	EditSchedule
	DeleteSchedule
	CreateAppointment
	ViewAppointment
	EditAppointment
	DeleteAppointment
	CancelAppointment
	CreateLaboratory
	EditLaboratory
	ViewLaboratory
	DeleteLaboratory
	CreateMedicalHistory
	ViewMedicalHistory
	EditMedicalHistory
	CreateEvolutionNote
	ViewEvolutionNote
	EditEvolutionNote
	CreatePrescription
	CreateIncapacity
	ViewIncapacity
	EditIncapacity
)

type RolePermission int

const ()
