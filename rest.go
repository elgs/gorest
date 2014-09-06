package gorest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

type Gorest struct {
	EnableHttp bool
	PortHttp   uint16
	HostHttp   string

	EnableHttps   bool
	PortHttps     uint16
	HostHttps     string
	CertFileHttps string
	KeyFileHttps  string

	UrlPrefix string
	Dbo       DataOperator
}

func (this *Gorest) Serve() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	handler := func(w http.ResponseWriter, r *http.Request) {
		context := make(map[string]interface{})
		context["api_token_id"] = r.Header.Get("api_token_id")
		context["api_token_key"] = r.Header.Get("api_token_key")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", r.Header.Get("Access-Control-Request-Method"))
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
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
				filter := r.Form["filter"]
				sort := r.FormValue("sort")
				s := r.FormValue("start")
				l := r.FormValue("limit")
				includeTotal := false
				array := false
				if t == "1" {
					includeTotal = true
				}
				if a == "1" {
					array = true
				}
				start, err := strconv.ParseInt(s, 10, 0)
				if err != nil {
					start = 0
					err = nil
				}
				limit, err := strconv.ParseInt(l, 10, 0)
				if err != nil {
					limit = 25
					err = nil
				}
				var data interface{}
				var total int64 = -1
				if array {
					data, total, err = this.Dbo.ListArray(tableId, filter, sort, start, limit, includeTotal, context)
				} else {
					data, total, err = this.Dbo.ListMap(tableId, filter, sort, start, limit, includeTotal, context)
				}
				m := map[string]interface{}{
					"data":  data,
					"total": total,
				}
				if err != nil {
					m["err"] = err.Error()
				}
				jsonData, err := json.Marshal(m)
				jsonString := string(jsonData)
				fmt.Fprint(w, jsonString)
			} else {
				// Load record by id.
				dataId := restData[1]

				data, err := this.Dbo.Load(tableId, dataId, context)

				m := map[string]interface{}{
					"data": data,
				}
				if err != nil {
					m["err"] = err.Error()
				}
				jsonData, err := json.Marshal(m)
				jsonString := string(jsonData)
				fmt.Fprint(w, jsonString)
			}
		case "POST":
			if tableId == "_query" {
				sqlSelect := r.FormValue("sql_select")
				sqlSelectCount := r.FormValue("sql_select_count")
				t := r.FormValue("total")
				a := r.FormValue("array")
				s := r.FormValue("start")
				l := r.FormValue("limit")
				includeTotal := false
				array := false
				if t == "1" {
					includeTotal = true
				}
				if a == "1" {
					array = true
				}
				start, err := strconv.ParseInt(s, 10, 0)
				if err != nil {
					start = 0
				}
				limit, err := strconv.ParseInt(l, 10, 0)
				if err != nil {
					limit = 25
				}
				var data interface{}
				var total int64 = -1
				if array {
					data, total, err = this.Dbo.QueryArray(tableId, sqlSelect, sqlSelectCount, start, limit, includeTotal, context)
				} else {
					data, total, err = this.Dbo.QueryMap(tableId, sqlSelect, sqlSelectCount, start, limit, includeTotal, context)
				}
				m := map[string]interface{}{
					"data":  data,
					"total": total,
				}
				if err != nil {
					m["err"] = err.Error()
				}
				jsonData, err := json.Marshal(m)
				jsonString := string(jsonData)
				fmt.Fprint(w, jsonString)
			} else {
				// Create the record.
				metaValues := r.URL.Query()["meta"]
				meta := false
				if metaValues != nil && metaValues[0] == "1" {
					meta = true
				}
				context["meta"] = meta
				decoder := json.NewDecoder(r.Body)
				m := make(map[string]interface{})
				err := decoder.Decode(&m)
				if err != nil {
					m["err"] = err.Error()
					jsonData, _ := json.Marshal(m)
					jsonString := string(jsonData)
					fmt.Fprint(w, jsonString)
					return
				}
				mUpper := make(map[string]interface{})
				for k, v := range m {
					mUpper[strings.ToUpper(k)] = v
				}
				data, err := this.Dbo.Create(tableId, mUpper, context)
				m = map[string]interface{}{
					"data": data,
				}
				if err != nil {
					m["err"] = err.Error()
				}
				jsonData, err := json.Marshal(m)
				jsonString := string(jsonData)
				fmt.Fprint(w, jsonString)
			}
		case "COPY":
			// Duplicate a new record.
			dataId := restData[1]
			data, err := this.Dbo.Duplicate(tableId, dataId, context)

			m := map[string]interface{}{
				"data": data,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			fmt.Fprint(w, jsonString)
		case "PUT":
			// Update an existing record.
			metaValues := r.URL.Query()["meta"]
			meta := false
			if metaValues != nil && metaValues[0] == "1" {
				meta = true
			}
			context["meta"] = meta
			decoder := json.NewDecoder(r.Body)
			m := make(map[string]interface{})
			err := decoder.Decode(&m)
			if err != nil {
				m["err"] = err.Error()
				jsonData, _ := json.Marshal(m)
				jsonString := string(jsonData)
				fmt.Fprint(w, jsonString)
				return
			}
			mUpper := make(map[string]interface{})
			for k, v := range m {
				mUpper[strings.ToUpper(k)] = v
			}
			data, err := this.Dbo.Update(tableId, mUpper, context)
			m = map[string]interface{}{
				"data": data,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			fmt.Fprint(w, jsonString)
		case "DELETE":
			// Remove the record.
			dataId := restData[1]

			data, err := this.Dbo.Delete(tableId, dataId, context)

			m := map[string]interface{}{
				"data": data,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			fmt.Fprint(w, jsonString)
		default:
			// Give an error message.
		}
	}
	http.HandleFunc("/", handler)

	if this.EnableHttp {
		go func() {
			fmt.Println(fmt.Sprint("Listening on http://", this.HostHttp, ":", this.PortHttp, "/", this.UrlPrefix))
			http.ListenAndServe(fmt.Sprint(this.HostHttp, ":", this.PortHttp), nil)
		}()
	}
	if this.EnableHttps {
		go func() {
			fmt.Println(fmt.Sprint("Listening on https://", this.HostHttps, ":", this.PortHttps, "/", this.UrlPrefix))
			http.ListenAndServeTLS(fmt.Sprint(this.HostHttps, ":", this.PortHttps), this.CertFileHttps, this.KeyFileHttps, nil)
		}()
	}
	if this.EnableHttp || this.EnableHttps {
		select {}
	} else {
		fmt.Println("Neither http nor https is listening, therefore I am quiting.")
	}
}
