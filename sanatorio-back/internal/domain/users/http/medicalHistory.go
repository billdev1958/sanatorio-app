package v1

import (
	"encoding/json"
	"net/http"
	"sanatorioApp/internal/domain/users/http/models"
)

func (h *handler) GetMedicalHistoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var request models.MedicalHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"status":"error","message":"invalid request payload","data":null}`, http.StatusBadRequest)
		return
	}

	if request.MedicalHistoryID == "" {
		http.Error(w, `{"status":"error","message":"missing medical_history_id","data":null}`, http.StatusBadRequest)
		return
	}

	responseData, err := h.uc.GetMedicalHistoryByID(r.Context(), request)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "Medical history retrieved successfully",
		Data:    responseData,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"status":"error","message":"failed to encode response","data":null}`, http.StatusInternalServerError)
		return
	}
}

func (h *handler) CompleteMedicalHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var request models.CompleteMedicalHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"status":"error","message":"invalid request payload","data":null}`, http.StatusBadRequest)
		return
	}

	medicalHistory, err := h.uc.CompleteMedicalHistory(r.Context(), request)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "Medical history updated successfully",
		Data:    medicalHistory,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"status":"error","message":"failed to encode response","data":null}`, http.StatusInternalServerError)
		return
	}
}
