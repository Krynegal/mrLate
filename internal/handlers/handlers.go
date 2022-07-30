package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"mrLate/internal/period"
	"net/http"
	"time"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	h := Handler{
		Router: mux.NewRouter(),
	}
	h.Router.HandleFunc("/main", h.Handler).Methods(http.MethodPost)
	return &h
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {

	// дата и время отправления
	departureTime := time.Date(2022, time.August, 18, 14, 50, 0, 0, time.UTC)
	fmt.Println(departureTime)

	// за сколько времени мы хотим быть на месте
	reservePeriod := period.Period{
		Hours:   0,
		Minutes: 40,
	}
	reserveDuration := reservePeriod.GetDuration()
	fmt.Println(reserveDuration)

	// во сколько мы должны прибыть на место
	timeDstArrival := departureTime.Add(-reserveDuration)
	fmt.Println(timeDstArrival)

	// сколько времени уйдет на дорогу
	// это время будем получать из API
	roadPeriod := period.Period{
		Hours:   1,
		Minutes: 25,
	}
	roadDuration := roadPeriod.GetDuration()
	fmt.Println(roadDuration)

	// во сколько мы должны выйти
	exitTime := timeDstArrival.Add(-roadDuration)
	fmt.Println(exitTime)

	w.Write([]byte(exitTime.Format("Mon 02 Jan 2006 15:04:05 MST")))
}
