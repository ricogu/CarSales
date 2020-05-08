package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ricogu/CarSales/pkg/database"
)

func main() {

	sqlManager, err := database.NewSqlManager("root:root@tcp(localhost:3306)/config")
	if nil != err {
		panic(err)
	}

	result, _ := sqlManager.ListTiresByWheel("Model 3")
	fmt.Printf("result: %v\n", result)
}
