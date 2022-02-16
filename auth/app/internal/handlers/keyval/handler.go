package keyval

import (
	"bitbucket.org/idomteam/idom-api/auth/internal/apperror"
	"bitbucket.org/idomteam/idom-api/auth/internal/client/keyval_service"
	"bitbucket.org/idomteam/idom-api/auth/pkg/jwt"
	"bitbucket.org/idomteam/idom-api/auth/pkg/logging"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


type Handler struct {
	Logger      logging.Logger
	Service keyval_service.Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/api/test", jwt.Middleware(apperror.Middleware(h.Test)))
	router.HandlerFunc(http.MethodGet, "/api/get-key", jwt.Middleware(apperror.Middleware(h.GetKey)))
	router.HandlerFunc(http.MethodPost, "/api/set-key", jwt.Middleware(apperror.Middleware(h.SetKey)))
}



func (h *Handler) Test(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("GET KEY")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

	return nil
}


func (h *Handler) GetKey(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("GET KEY")
	w.Header().Set("Content-Type", "application/json")

	h.Logger.Debug("get uuid from context")
	key := r.URL.Query().Get("key")
	if key == "" {
		return apperror.BadRequestError("uuid query parameter is required and must be a comma separated integers")
	}

	note, err := h.Service.GetKey(r.Context(), key)
	if err != nil {
		return err
	}
	noteBytes, err := json.Marshal(note)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(noteBytes)

	return nil
}

func (h *Handler) SetKey(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("GET KEY")
	w.Header().Set("Content-Type", "application/json")

	h.Logger.Debug("get uuid from context")
	key := r.URL.Query().Get("key")
	val := r.URL.Query().Get("val")
	if key == "" {
		return apperror.BadRequestError("key query parameter is required and must be a comma separated integers")
	}
	if val == "" {
		return apperror.BadRequestError("val query parameter is required and must be a comma separated integers")
	}

	err := h.Service.SetKey(r.Context(), key, val)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("added"))

	return nil
}
