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

// Handle godoc
// @Summary Lab3
// @Description Lab3
// @Tags Lab3
// @Param group_name query string true "Название группы"
// @Accept json
// @Produce json
// @Router /lab3 [get]
// @Success 200 {object} []service.Lab3ReportItem
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	groupName := query.Get("group_name")

	in := &service.In{
		GroupName: groupName,
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

func Init(r *chi.Mux, h *Handler, middlewares ...func(handler http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/lab3", h.Handle)
	})
}
