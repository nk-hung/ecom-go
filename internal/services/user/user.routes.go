package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nk-hung/ecom-go/pkg/utils"
	"github.com/nk-hung/ecom-go/types"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("login", h.handleLogin).Methods("POST")
	router.HandleFunc("register", h.handleRegister).Methods("POST")
	router.HandleFunc("hello", h.handleRegister).Methods("GET")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if user exist
	// if it doesn't we create the user
}

func (h *Handler) test(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }
