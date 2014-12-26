package gorest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

type Gorest struct {
	EnableHttp bool
	PortHttp   uint16
	HostHttp   string

	EnableHttps   bool
	PortHttps     uint16
	HostHttps     string
	CertFileHttps string
	CertRootHttps string
	KeyFileHttps  string

	SessionKey   string
	FileBasePath string
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
		for kUrlPrefix, dataOperator := range dataOperatorRegistry {
			switch dataOperator := dataOperator.(type) {
			case DataOperator:
				if strings.HasPrefix(urlPath, "/"+kUrlPrefix+"/") {
					dbo = dataOperator.(DataOperator)
					urlPrefix = kUrlPrefix
					break
				}
			case func(w http.ResponseWriter, r *http.Request):
				if urlPath == kUrlPrefix {
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

		context := make(map[string]interface{})
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
		switch r.Method {
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
				if c != "upper" && c != "camel" {
					c = "lower"
				}
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
					data, total, err = dbo.ListArray(tableId, field, filter, sort, group, start, limit, includeTotal, context)
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
				if c != "upper" && c != "camel" {
					c = "lower"
				}
				context["case"] = c

				field := r.Form["field"]

				bin := false
				b := r.FormValue("bin")
				if b == "1" {
					bin = true
				}
				context["bin"] = bin

				data, err := dbo.Load(tableId, dataId, field, context)

				if bin && err == nil {
					filePath := context["file_path"].(string)
					fileName := context["file_name"].(string)
					filesize := context["file_size"].(int64)
					file, _ := os.Open(this.FileBasePath + filePath + dataId)
					defer file.Close()
					file.Seek(0, os.SEEK_SET)
					n, _ := io.CopyN(w, file, filesize)
					w.Header().Set("Content-Length", strconv.FormatInt(n, 10))
					w.Header().Set("Content-Type", "application/octet-stream")
					w.Header().Set("Content-disposition", "attachment; filename='"+fileName+"'")
				} else {
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
			}
		case "POST":
			// Create the record.
			metaValues := r.URL.Query()["meta"]
			meta := false
			if metaValues != nil && metaValues[0] == "1" {
				meta = true
			}
			context["meta"] = meta

			bin := false
			b := r.FormValue("bin")
			if b == "1" {
				bin = true
			}

			if bin {
				context["meta"] = true
				m := make(map[string]interface{})

				file, header, err := r.FormFile("file")
				defer file.Close()

				if err != nil {
					fmt.Println(err)
					m["err"] = err.Error()
				}

				id := uuid.New()
				filePath := fmt.Sprint(this.FileBasePath, string(os.PathSeparator), id[0:2], string(os.PathSeparator), id)
				os.MkdirAll(fmt.Sprint(this.FileBasePath, string(os.PathSeparator), id[0:2]), os.FileMode(0755))
				out, err := os.Create(filePath)
				if err != nil {
					fmt.Println(err)
					m["err"] = err.Error()
				}
				defer out.Close()

				// write the content from POST to the file
				written, err := io.Copy(out, file)
				if err != nil {
					m["err"] = err.Error()
				}

				buf := make([]byte, written)
				hash := sha256.New()
				hash.Write(buf)
				md := hash.Sum(nil)
				mdStr := hex.EncodeToString(md)

				m["NAME"] = header.Filename
				m["PATH"] = "/" + id[0:2] + "/"
				m["SIZE"] = written
				m["CHECKSUM"] = mdStr
				m["ID"] = id

				data, err := dbo.Create(tableId, m, context)
				m = map[string]interface{}{
					"data": data,
				}
				if err != nil {
					m["err"] = err.Error()
				}

				jsonData, err := json.Marshal(m)
				jsonString := string(jsonData)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				fmt.Fprint(w, jsonString)
			} else {
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
				data, err := dbo.Create(tableId, mUpper, context)
				m = map[string]interface{}{
					"data": data,
				}
				if err != nil {
					m["err"] = err.Error()
				}
				jsonData, err := json.Marshal(m)
				jsonString := string(jsonData)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				fmt.Fprint(w, jsonString)
			}
		case "COPY":
			// Duplicate a new record.
			dataId := restData[1]
			data, err := dbo.Duplicate(tableId, dataId, context)

			m := map[string]interface{}{
				"data": data,
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
			data, err := dbo.Update(tableId, mUpper, context)
			m = map[string]interface{}{
				"data": data,
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

			bin := false
			b := r.FormValue("bin")
			if b == "1" {
				bin = true
			}
			context["bin"] = bin

			load := false
			l := r.FormValue("load")
			if l == "1" {
				load = true
			}
			context["load"] = load

			context["file_base_path"] = this.FileBasePath

			data, err := dbo.Delete(tableId, dataId, context)

			m := map[string]interface{}{
				"data": data,
			}
			if err != nil {
				m["err"] = err.Error()
			}
			jsonData, err := json.Marshal(m)
			jsonString := string(jsonData)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, jsonString)
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
