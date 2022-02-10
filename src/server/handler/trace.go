package handler

import (
	"log"
	"net/http"

	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func trace(router chi.Router) {
	router.Post("/", contactTrace)
}

func contactTrace(w http.ResponseWriter, r *http.Request) {
	traceReq := &models.TraceRequest{}
	if err := render.Bind(r, traceReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	_, err := dbInstance.UpdateCovidPositive(traceReq.UserID, true)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	users, err := dbInstance.GetContactUsers(traceReq.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
