package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	tableId := "test.t"
	RegisterDataInterceptor(tableId, &MyDataInterceptor{})
	ds := "user:pass@tcp(host:3306)/test"
	db, err := sql.Open("mysql", ds)
	defer db.Close()

	if err != nil {
		fmt.Println("sql.Open:", err)
	}
	dbo := &DbOperator{TableId: tableId, Db: db}
	data := dbo.Load("0")

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	jsonString := string(json)

	fmt.Println(jsonString)
}

func (this *MyDataInterceptor) BeforeCreate(data map[string]interface{}) bool {
	fmt.Println("Here I'm in BeforeCreate")
	return true
}

func (this *MyDataInterceptor) BeforeLoad(id string) bool {
	fmt.Println("Here I'm in BeforeLoad")
	return true
}

func (this *MyDataInterceptor) AfterLoad(data map[string]string) {
	fmt.Println("Here I'm in AfterLoad")
}
