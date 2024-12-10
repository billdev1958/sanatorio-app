package models

type OfficeResponse struct {
	OfficeID   int    `json:"office_id"`
	OfficeName string `json:"office_name"`
}

type UpdateOfficeRequest struct {
	OfficeID   int    `json:"office_id"`
	OfficeName string `json:"office_name"`
}
