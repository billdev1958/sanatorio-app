package v1

import "net/http"

func (h *handler) EmailRoutes(mux *http.ServeMux) {
	mux.Handle("POST /v1/patient/confirm", http.HandlerFunc(h.ConfirmAccount))

}
