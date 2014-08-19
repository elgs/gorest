package gorest

import (
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
			if len(restData) == 1 ||
				strings.HasPrefix(restData[1], "?") ||
				len(restData[1]) == 0 {
				//List records.
				t := r.FormValue("total")
				a := r.FormValue("array")
				includeTotal := false
				array := false
				if t == "1" {
					includeTotal = true
				}
				if a == "1" {
					array = true
				}
				var data interface{}
				var total int64 = -1
				dbo := &MySqlDataOperator{Ds: this.Ds, TableId: tableId}
				if array {
					data, total = dbo.ListArray("", "", 0, 25, includeTotal)
				} else {
					data, total = dbo.ListMap("", "", 0, 25, includeTotal)
				}
				m := map[string]interface{}{
					"data":  data,
					"total": total,
				}
				json, err := json.Marshal(m)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				jsonString := string(json)
				fmt.Fprintf(w, jsonString)
			} else {
				// Load record by id.
				dataId := restData[1]

				dbo := &MySqlDataOperator{Ds: this.Ds, TableId: tableId}
				data := dbo.Load(dataId)

				json, err := json.Marshal(data)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				jsonString := string(json)
				fmt.Fprintf(w, jsonString)
			}
		case "POST":
			// Create the record.
			decoder := json.NewDecoder(r.Body)
			var m map[string]interface{}
			err := decoder.Decode(&m)
			if err != nil {
				fmt.Println(err)
			}
			mUpper := make(map[string]interface{})
			for k, v := range m {
				mUpper[strings.ToUpper(k)] = v
			}
			dbo := &MySqlDataOperator{Ds: this.Ds, TableId: tableId}
			data := dbo.Create(mUpper)
			json, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}
			jsonString := string(json)
			fmt.Fprintf(w, jsonString)
		case "COPY":
			// Duplicate a new record.
			dataId := restData[1]

			dbo := &MySqlDataOperator{Ds: this.Ds, TableId: tableId}
			data := dbo.Duplicate(dataId)

			json, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}
			jsonString := string(json)
			fmt.Fprintf(w, jsonString)
		case "PUT":
			// Update an existing record.
			decoder := json.NewDecoder(r.Body)
			var m map[string]interface{}
			err := decoder.Decode(&m)
			if err != nil {
				fmt.Println(err)
			}
			mUpper := make(map[string]interface{})
			for k, v := range m {
				mUpper[strings.ToUpper(k)] = v
			}
			dbo := &MySqlDataOperator{Ds: this.Ds, TableId: tableId}
			data := dbo.Update(mUpper)
			json, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}
			jsonString := string(json)
			fmt.Fprintf(w, jsonString)
		case "DELETE":
			// Remove the record.
			dataId := restData[1]

			dbo := &MySqlDataOperator{Ds: this.Ds, TableId: tableId}
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
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprint(this.Host, ":", this.Port), nil)
}
