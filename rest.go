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

	SessionKey   string
	FileBasePath string
	Mode         string
}

func MakeCookie(key string, m map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonData)
	return EncryptStr(key, jsonString)
}

func ReadCookie(key string, s string) (map[string]interface{}, error) {
	text, err := DecryptStr(key, s)
	var m map[string]interface{}
	err = json.Unmarshal([]byte(text), &m)
	return m, err
}

func (this *Gorest) Serve() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	handler := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", r.Header.Get("Access-Control-Request-Method"))
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		urlPath := r.URL.Path
		//fmt.Println(urlPath)
		var dbo DataOperator = nil
		var urlPrefix string
		context := make(map[string]interface{})
		for kUrlPrefix, dataOperator := range dataOperatorRegistry {
			switch dataOperator := dataOperator.(type) {
			case DataOperator:
				if strings.HasPrefix(urlPath, "/"+kUrlPrefix+"/") {

					context["api_token_id"] = r.Header.Get("api_token_id")
					context["api_token_key"] = r.Header.Get("api_token_key")
					if len(this.SessionKey) > 0 {
						cookieUser, err := r.Cookie("user")
						if cookieUser != nil && err == nil {
							mapCookies, err := ReadCookie(this.SessionKey, cookieUser.Value)
							if err == nil {
								userId := mapCookies["user_id"]
								tokenKey := mapCookies["token_key"]
								if userId != nil && tokenKey != nil {
									context["api_token_id"] = userId
									context["api_token_key"] = tokenKey
								}
							}
						}
					}

					dbo = dataOperator.(DataOperator)
					urlPrefix = kUrlPrefix
					break
				}
			case func(w http.ResponseWriter, r *http.Request):
				if r.Method == "OPTIONS" {
					continue
				}
				if urlPath == kUrlPrefix {

					if len(this.SessionKey) > 0 {
						cookieUser, err := r.Cookie("user")
						if cookieUser != nil && err == nil {
							mapCookies, err := ReadCookie(this.SessionKey, cookieUser.Value)
							if err == nil {
								userId := mapCookies["user_id"]
								tokenKey := mapCookies["token_key"]
								if userId != nil && tokenKey != nil {
									r.Header.Set("api_token_id", userId.(string))
									r.Header.Set("api_token_key", tokenKey.(string))
								}
							}
						}
					}

					dataOperator(w, r)
					return
				}
			default:
				fmt.Println("Unknow dbo.")
				return
			}
		}
		if dbo == nil {
			return
		}

		restUrl := urlPath[len(urlPrefix)+2:]
		restData := strings.Split(restUrl, "/")
		tableId := restData[0]

		switch r.Method {
		case "META":
			m, err := dbo.MetaData(tableId)
			if err != nil {
				fmt.Println(err)
			}
			jsonData, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}
			jsonString := string(jsonData)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, jsonString)
		case "GET":
			if len(restData) == 1 ||
				strings.HasPrefix(restData[1], "?") ||
				len(restData[1]) == 0 {
				//List records.
				t := r.FormValue("total")
				a := r.FormValue("array")
				filter := r.Form["filter"]
				field := r.Form["field"]
				sort := r.FormValue("sort")
				group := r.FormValue("group")
				s := r.FormValue("start")
				l := r.FormValue("limit")
				c := r.FormValue("case")
				context["case"] = c
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
					var headers []string
					var dataArray [][]string
					headers, dataArray, total, err = dbo.ListArray(tableId, field, filter, sort, group, start, limit, includeTotal, context)
					data = map[string]interface{}{
						"headers": headers,
						"data":    dataArray,
					}
				} else {
					data, total, err = dbo.ListMap(tableId, field, filter, sort, group, start, limit, includeTotal, context)
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
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				fmt.Fprint(w, jsonString)
			} else {
				// Load record by id.
				dataId := restData[1]
				c := r.FormValue("case")
				context["case"] = c

				field := r.Form["field"]

				data, err := dbo.Load(tableId, dataId, field, context)

				m := map[string]interface{}{
					"data": data,
				}
				if err != nil {
					m["err"] = err.Error()
				}
				jsonData, _ := json.Marshal(m)
				jsonString := string(jsonData)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				fmt.Fprint(w, jsonString)
			}
		case "POST":
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
			data, info, err := dbo.Create(tableId, mUpper, context)
			m = map[string]interface{}{
				"data": data,
				"info": info,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, jsonString)
		case "COPY":
			// Duplicate a new record.
			dataId := restData[1]
			data, info, err := dbo.Duplicate(tableId, dataId, context)

			m := map[string]interface{}{
				"data": data,
				"info": info,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, jsonString)
		case "PUT":
			// Update an existing record.
			metaValues := r.URL.Query()["meta"]
			meta := false
			if metaValues != nil && metaValues[0] == "1" {
				meta = true
			}
			context["meta"] = meta
			dataId := ""
			if len(restData) >= 2 {
				dataId = restData[1]
			}
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
			if dataId != "" {
				mUpper["ID"] = dataId
			}
			data, info, err := dbo.Update(tableId, mUpper, context)
			m = map[string]interface{}{
				"data": data,
				"info": info,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, jsonString)
		case "DELETE":
			// Remove the record.
			dataId := restData[1]

			load := false
			l := r.FormValue("load")
			if l == "1" {
				load = true
			}
			context["load"] = load

			data, info, err := dbo.Delete(tableId, dataId, context)

			m := map[string]interface{}{
				"data": data,
				"info": info,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, jsonString)
		case "OPTIONS":
		default:
			// Give an error message.
		}
	}
	http.HandleFunc("/", handler)

	if this.EnableHttp {
		go func() {
			fmt.Println(fmt.Sprint("Listening on http://", this.HostHttp, ":", this.PortHttp, "/"))
			err := http.ListenAndServe(fmt.Sprint(this.HostHttp, ":", this.PortHttp), nil)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	if this.EnableHttps {
		go func() {
			fmt.Println(fmt.Sprint("Listening on https://", this.HostHttps, ":", this.PortHttps, "/"))
			err := http.ListenAndServeTLS(fmt.Sprint(this.HostHttps, ":", this.PortHttps), this.CertFileHttps, this.KeyFileHttps, nil)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	if this.EnableHttp || this.EnableHttps {
		select {}
	} else {
		fmt.Println("Neither http nor https is listening, therefore I am quiting.")
	}
}
