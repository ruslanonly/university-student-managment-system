package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/lab2/service"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/api/res"
)

type Handler struct {
	s *service.Service
}

// Handle godoc
// @Summary Lab2
// @Description Lab2
// @Tags Lab2
// @Param course_name query string true "Название курса (дисциплины)"
// @Param year query string true "Год"
// @Param semester query string true "Семестр"
// @Accept json
// @Produce json
// @Router /lab2 [get]
// @Success 200 {object} []service.Lab2ReportItem
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	courseName := query.Get("course_name")
	year := query.Get("year")
	semester := query.Get("semester")

	parsedYear, err := strconv.Atoi(year)
	if err != nil {
		http.Error(w, "Invalid year: must be an integer", http.StatusBadRequest)
		return
	}

	parsedSemester, err := strconv.Atoi(semester)
	if err != nil {
		http.Error(w, "Invalid semester: must be an integer", http.StatusBadRequest)
		return
	}

	in := &service.In{
		CourseName:	courseName,
		Year: 		parsedYear,
		Semester:   parsedSemester,
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
	r.Get("/lab2", h.Handle)
}
