package v1

import "net/http"

func (h *handler) UserRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/v1/users", h.RegisterUser)

}
