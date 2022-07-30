package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"mrLate/internal/period"
	"net/http"
	"time"
)

type Request struct {
	Departure         time.Time
	ReservePeriod     period.Period
	AdditionalReserve period.Period
	// address
}

//func (r Request) convertDeparture() (time.Time, error) {
//	departureTime, err := time.Parse(time.RFC3339, r.Departure)
//	if err != nil {
//		return time.Time{}, err
//	}
//	return departureTime, err
//}

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	h := Handler{
		Router: mux.NewRouter(),
	}
	h.Router.HandleFunc("/", h.Handler).Methods(http.MethodPost)
	return &h
}

func getData(r *http.Request) (Request, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return Request{}, err
	}
	req := Request{}
	if err = json.Unmarshal(body, &req); err != nil {
		return Request{}, err
	}
	fmt.Printf("req: %v\n", req)
	return req, nil
}

func calcExitTime(req Request, roadPeriod period.Period) time.Time {
	// дата и время отправления
	//departureTime := time.Date(2022, time.August, 18, 14, 50, 0, 0, time.UTC)
	departureTime := req.Departure
	fmt.Printf("departureTime: %v\n", departureTime)

	// за сколько времени мы хотим быть на месте
	reservePeriod := req.ReservePeriod
	reserveDuration := reservePeriod.GetDuration()
	fmt.Printf("reserveDuration: %v\n", reserveDuration)

	// во сколько мы должны прибыть на место
	timeDstArrival := departureTime.Add(-reserveDuration)
	fmt.Printf("timeDstArrival: %v\n", timeDstArrival)

	roadDuration := roadPeriod.GetDuration()
	fmt.Printf("roadDuration: %v\n", roadDuration)

	// во сколько мы должны выйти
	exitTime := timeDstArrival.Add(-roadDuration)
	fmt.Printf("exitTime: %v\n", exitTime)
	return exitTime
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	req, err := getData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// тут должен быть запрос к API 2gis с пунктом отправления, прибытия и видом транспорта

	// сколько времени уйдет на дорогу
	// это время будем получать из API
	roadPeriod := period.Period{
		Hours:   1,
		Minutes: 25,
	}
	exitTime := calcExitTime(req, roadPeriod)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(exitTime.Format("Mon 02 Jan 2006 15:04:05 MST")))
}
