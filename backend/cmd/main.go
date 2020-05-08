package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ricogu/CarSales/pkg/database"
	"github.com/ricogu/CarSales/pkg/handler"
	"log"
	"net/http"
)

func main() {

	log.Printf("Starting service\n")

	sqlManager, err := database.NewSqlManager("root:root@tcp(localhost:3306)/config")
	if nil != err {
		panic(err)
	}

	RESTHandler := handler.NewRESTHandler(*sqlManager)

	router := mux.NewRouter()

	router.HandleFunc("/Batteries", RESTHandler.GetBatteries).Methods("GET")
	router.HandleFunc("/Wheels", RESTHandler.GetWheels).Methods("GET")
	router.HandleFunc("/Tires", RESTHandler.GetTires).Methods("GET")
	router.HandleFunc("/Orders", RESTHandler.SubmitOrder).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}
