package v1

import (
	"encoding/json"
	"net/http"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/http/models"
)

type handler struct {
	uc user.Usecase
}

func NewHandler(uc user.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud en el modelo LoginUser
	var request models.LoginUser
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para autenticar al usuario
	response, err := h.uc.LoginUser(r.Context(), request)
	if err != nil {
		// Manejar el error de autenticación
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa con el token JWT
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud en el modelo RegisterUserRequest
	request := models.RegisterUserByAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Respuesta de error si el payload no es válido
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para registrar el usuario
	response, err := h.uc.RegisterUser(r.Context(), request)
	if err != nil {
		// Manejar el error y enviar una respuesta adecuada usando response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa usando response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extraer el accountID de la ruta del URL
	accountID := r.PathValue("accountID")
	if accountID == "" {
		http.Error(w, "accountID es obligatorio", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para obtener el usuario por ID
	response, err := h.uc.GetUserByID(r.Context(), accountID)
	if err != nil {
		// Manejar el error y enviar una respuesta adecuada
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido y el estado de la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta con los datos del usuario
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	// Extraer el accountID de la ruta del URL
	accountID := r.PathValue("accountID")
	if accountID == "" {
		http.Error(w, "accountID es obligatorio", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para obtener el doctor por ID
	response, err := h.uc.GetDoctorByID(r.Context(), accountID)
	if err != nil {
		// Manejar el error y enviar una respuesta adecuada
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido y el estado de la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta con los datos del doctor
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Llamar al caso de uso para obtener la lista de usuarios
	users, err := h.uc.GetUsers(r.Context())
	if err != nil {
		// Manejar el error y enviar una respuesta adecuada
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve users",
			Errors:  err.Error(),
		})
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la lista de usuarios en la respuesta
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func (h *handler) GetDoctors(w http.ResponseWriter, r *http.Request) {
	// Llamar al caso de uso para obtener la lista de doctores
	doctors, err := h.uc.GetDoctors(r.Context())
	if err != nil {
		// Manejar el error y enviar una respuesta adecuada
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve doctors",
			Errors:  err.Error(),
		})
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la lista de doctores en la respuesta
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "Doctors retrieved successfully",
		Data:    doctors,
	})
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud en el modelo UpdateUser
	var request models.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para actualizar el usuario
	response, err := h.uc.UpdateUser(r.Context(), request)
	if err != nil {
		// Manejar el error de actualización
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud en el modelo UpdateDoctor
	var request models.UpdateDoctor
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para actualizar el doctor
	response, err := h.uc.UpdateDoctor(r.Context(), request)
	if err != nil {
		// Manejar el error de actualización
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud para obtener el account_id
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para eliminar el usuario
	response, err := h.uc.DeleteUser(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud para obtener el account_id
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para eliminar el doctor
	response, err := h.uc.DeleteDoctor(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) SoftDeleteUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud para obtener el account_id
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para realizar el soft delete del usuario
	response, err := h.uc.SoftDeleteUser(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) SoftDeleteDoctor(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud para obtener el account_id
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para realizar el soft delete del doctor
	response, err := h.uc.SoftDeleteDoctor(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
