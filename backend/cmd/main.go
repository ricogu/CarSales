package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ricogu/CarSales/pkg/database"
	"github.com/ricogu/CarSales/pkg/handler"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	log.Printf("Starting service\n")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_schema := os.Getenv("DB_DATABASE")
	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_host, db_schema)

	//leave this statement for local db testing
	//connString := "root:root@tcp(localhost:3306)/config"

	log.Println(connString)

	sqlManager, err := database.NewSqlManager(connString)
	for i := 0; i < 10 && nil != err; i++ {
		time.Sleep(2 * time.Second)
		log.Printf("Reconnecting DB service\n")
		sqlManager, err = database.NewSqlManager(connString)
	}

	if err != nil {
		panic(err)
	}

	RESTHandler := handler.NewRESTHandler(*sqlManager)

	router := mux.NewRouter()

	router.HandleFunc("/Batteries", RESTHandler.GetBatteries).Methods("GET")
	router.HandleFunc("/Wheels", RESTHandler.GetWheels).Methods("GET")
	router.HandleFunc("/Tires", RESTHandler.GetTires).Methods("GET")
	router.HandleFunc("/Orders", RESTHandler.SubmitOrder).Methods("POST")
	router.HandleFunc("/Orders", RESTHandler.GetOrders).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}
