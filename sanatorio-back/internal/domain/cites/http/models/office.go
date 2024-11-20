package models

type OfficeResponse struct {
	OfficeID       int    `json:"office_id"`
	OfficeStatusID int    `json:"office_status"`
	OfficeName     string `json:"office_name"`
	StatusName     string `json:"status_name"`
}

type UpdateOfficeRequest struct {
	OfficeID   int    `json:"office_id"`
	OfficeName string `json:"office_name"`
}
