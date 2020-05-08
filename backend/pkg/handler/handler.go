package handler

import (
	"encoding/json"
	"errors"
	"github.com/ricogu/CarSales/pkg/data"
	"github.com/ricogu/CarSales/pkg/database"
	"log"
	"net/http"
	"strconv"
)

type RESTHandler struct {
	sql database.SqlManager
}

func NewRESTHandler(conn database.SqlManager) *RESTHandler {
	return &RESTHandler{sql: conn}
}

func (handler *RESTHandler) GetBatteries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("#getBatteries")

	batteries, err := handler.sql.ListAllBatteries()
	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(batteries)
}

func (handler *RESTHandler) GetWheels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("#getWheelsByBattery")

	batteryIds, ok := r.URL.Query()["batteryId"]

	if !ok || len(batteryIds[0]) < 1 {
		log.Println("Url Param 'batteryId' is missing")
		handleError(w, errors.New("Url Param 'batteryId' is missing"))
		return
	}

	batteryId, err := strconv.Atoi(batteryIds[0])

	if err != nil {
		handleError(w, err)
		return
	}

	wheels, err := handler.sql.ListWheelsByBattery(batteryId)

	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(wheels)
}

func (handler *RESTHandler) GetTires(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("#getTiresByWheel")

	batteryIds, ok := r.URL.Query()["wheelId"]

	if !ok || len(batteryIds[0]) < 1 {
		log.Println("Url Param 'wheelId' is missing")
		handleError(w, errors.New("Url Param 'wheelId' is missing"))
		return
	}

	wheelId, err := strconv.Atoi(batteryIds[0])

	if err != nil {
		handleError(w, err)
		return
	}

	tires, err := handler.sql.ListTiresByWheel(wheelId)

	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(tires)
}

func (handler *RESTHandler) SubmitOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("#Received Order Submission")

	var orderReceived data.OrderSubmission
	err := json.NewDecoder(r.Body).Decode(&orderReceived)

	if err != nil {
		handleError(w, err)
		return
	}

	orderSubmitted, err := handler.sql.SubmitOrder(orderReceived.BatteryId, orderReceived.TireId, orderReceived.WheelId)

	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(orderSubmitted)
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("Error: %s\n", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte([]byte(err.Error())))
}
