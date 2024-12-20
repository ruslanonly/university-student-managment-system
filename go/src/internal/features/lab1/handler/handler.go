package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/lab1/service"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/api/res"
)

type Handler struct {
	s *service.Service
}

// Handle godoc
// @Summary Lab1
// @Description Lab1
// @Tags Lab1
// @Param phrase query string true "Фраза для поиска"
// @Param period_from query string true "Дата начала периода"
// @Param period_to query string true "Дата окончания периода"
// @Accept json
// @Produce json
// @Router /lab1 [get]
// @Success 200 {object} []service.ReportItem
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	phrase := query.Get("phrase")
	periodFromSrt := query.Get("period_from")

	periodFrom, err := time.Parse(time.DateOnly, periodFromSrt)

	if err != nil {
		res.WriteError(w, http.StatusBadRequest, err, err.Error())
		return
	}

	periodToSrt := query.Get("period_to")

	periodTo, err := time.Parse(time.DateOnly, periodToSrt)

	if err != nil {
		res.WriteError(w, http.StatusBadRequest, err, err.Error())
		return
	}

	result, err := h.s.Execute(r.Context(), service.In{
		Phrase:      phrase,
		PeriodStart: periodFrom,
		PeriodEnd:   periodTo,
	})

	if err != nil {
		res.WriteError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	res.WriteJSON(w, http.StatusOK, result)
}

func New(s *service.Service) *Handler {
	return &Handler{
		s: s,
	}
}

func Init(r *chi.Mux, h *Handler) {
	r.Get("/lab1", h.Handle)
}
