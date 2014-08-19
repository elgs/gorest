package gorest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Gorest struct {
	Port      uint16
	Host      string
	UrlPrefix string
	Ds        string
}

func (this *Gorest) Serve() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		urlPath := r.URL.Path
		urlPrefix := fmt.Sprint("/", this.UrlPrefix, "/")
		if !strings.HasPrefix(urlPath, urlPrefix) {
			return
		}
		restUrl := urlPath[len(urlPrefix):]
		restData := strings.Split(restUrl, "/")
		tableId := restData[0]
		switch r.Method {
		case "GET":
			// Serve the resource.
			dataId := restData[1]

			dbo, err := getDbo(this.Ds, tableId)
			if err != nil {
				fmt.Println(err)
			}
			data := dbo.Load(dataId)

			json, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}
			jsonString := string(json)
			fmt.Fprintf(w, jsonString)
		case "POST":
			// Create a new record.
			fmt.Println(r.Method, ": ", urlPath)
		case "PUT":
			// Update an existing record.
			fmt.Println(r.Method, ": ", urlPath)
		case "DELETE":
			// Remove the record.
			dataId := restData[1]

			dbo, err := getDbo(this.Ds, tableId)
			if err != nil {
				fmt.Println(err)
			}
			data := dbo.Delete(dataId)

			json, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}
			jsonString := string(json)
			fmt.Fprintf(w, jsonString)
		default:
			// Give an error message.
		}
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprint(this.Host, ":", this.Port), nil)
}

func getConn(ds string) (*sql.DB, error) {
	db, err := sql.Open("mysql", ds)
	db.SetMaxIdleConns(10)
	return db, err
}

func getDbo(ds string, tableId string) (DataOperator, error) {
	db, err := getConn(ds)
	return &MySqlDataOperator{TableId: tableId, Db: db}, err
}
