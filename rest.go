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
				p := r.URL.Query()
				t := p["total"]
				includeTotal := false
				if t != nil && len(t) > 0 && t[0] == "1" {
					includeTotal = true
				}
				dbo := getDbo(this.Ds, tableId)
				data, total := dbo.List("", "", 0, 25, includeTotal)
				m := map[string]interface{}{
					"data":  data,
					"total": total,
				}
				json, err := json.Marshal(m)
				if err != nil {
					fmt.Println(err)
				}
				jsonString := string(json)
				fmt.Fprintf(w, jsonString)
			} else {
				// Load record by id.
				dataId := restData[1]

				dbo := getDbo(this.Ds, tableId)
				data := dbo.Load(dataId)

				json, err := json.Marshal(data)
				if err != nil {
					fmt.Println(err)
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
			dbo := getDbo(this.Ds, tableId)
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

			dbo := getDbo(this.Ds, tableId)
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
			dbo := getDbo(this.Ds, tableId)
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

			dbo := getDbo(this.Ds, tableId)
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

func getDbo(ds string, tableId string) DataOperator {
	return &MySqlDataOperator{Ds: ds, TableId: tableId}
}
