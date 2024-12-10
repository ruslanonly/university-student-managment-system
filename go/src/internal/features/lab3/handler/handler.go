package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/lab3/service"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/api/res"
	"net/http"
)

type Handler struct {
	s *service.Service
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	groupName := query.Get("group_name")
	groupEnrollmentYear := query.Get("group_enrollment_year")
	DisciplineTag := query.Get("discipline_tag")

	in := &service.In{
		GroupName:           groupName,
		GroupEnrollmentYear: groupEnrollmentYear,
		DisciplineTag:       DisciplineTag,
	}

	out, err := h.s.Execute(r.Context(), in)

	if err != nil {
		res.WriteError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	res.WriteJSON(w, http.StatusOK, out)
}

func New(s *service.Service) *Handler {
	return &Handler{
		s: s,
	}
}

func Init(r *chi.Mux, h *Handler) {
	r.Get("/lab3", h.Handle)
}
