gorest
======

A restful framework in Golang.

#Installation
`go get -u "github.com/elgs/gorest"`

#Sample Code
```go
package main

import (
	"github.com/elgs/gorest"
)

func main() {
	ds := "user:pass@tcp(host:3306)/test"
	dbType := "mysql"
	tokenTable := "db_name.table_name"
	dbo := &gorest.MySqlDataOperator{
		Ds:         ds,
		DbType:     dbType,
		TokenTable: tokenTable,
	}
	gorest.RegisterDataOperator("api", dbo)
	r := &gorest.Gorest{
		EnableHttp: true,
		HostHttp:   "0.0.0.0",
		PortHttp:   8080,

		EnableHttps:   true,
		HostHttps:     "0.0.0.0",
		PortHttps:     8433,
		CertFileHttps: "cert.crt",
		KeyFileHttps:  "private.key",
	}
	r.Serve()
}
```
where in the test.TEST table:
```
ID	NAME		CREATE_TIME

0	Alicia	2014-08-18 18:43:38
1	Brian	2014-08-18 18:43:38
2	Chloe	2014-08-18 18:43:38
4	Bianca	2014-08-18 18:43:38
5	Leo		2014-08-18 18:43:38
6	Joy		2014-08-18 18:43:38
7	Samuel	2014-08-18 18:43:38
```

#Client
`curl http://localhost:8080/api/test.test`
outputs(beautified manually):
```json
{
	"data": [{
		"create_time": "2014-08-18 18:43:38",
		"id": "0",
		"name": "Alicia"
	}, {
		"create_time": "2014-08-18 18:43:38",
		"id": "1",
		"name": "Brian"
	}, {
		"create_time": "2014-08-18 18:43:38",
		"id": "2",
		"name": "Chloe"
	}, {
		"create_time": "2014-08-18 18:43:38",
		"id": "4",
		"name": "Bianca"
	}, {
		"create_time": "2014-08-18 18:43:38",
		"id": "5",
		"name": "Leo"
	}, {
		"create_time": "2014-08-18 18:43:38",
		"id": "6",
		"name": "Joy"
	}, {
		"create_time": "2014-08-18 18:43:38",
		"id": "7",
		"name": "Samuel"
	}],
	"total": -1
}
```

#Caveat
The default implementation of `MySqlDataOperator` assumes:

1. all field names are in upper case;
2. each table has a field `ID` as the primary key;
3. the `ID` field is of char(36);

These restrictions will likely be removed, or be made configurable in the 
future, by improvement or another implementation.

The `DefaultDataInterceptor` is intended to be implemented to extend the 
business logic of the applications. And the `DefaultDataOperator` is intended 
to be implemented by database providers, in order to connect to other databases
or other data sources.

#API
Gets from test table data in the first page(1-25)(in JSON):
```
curl -X GET -H "api_token_id:0" -H "api_token_key:0" "http://localhost:8080/api/test"
```
Returns:
```json
{"data":[{"create_time":"2014-08-18 18:43:38","id":"0","name":"Alicia"},{"create_time":"2014-08-18 18:43:38","id":"1","name":"Brian"},{"create_time":"2014-08-18 18:43:38","id":"2","name":"Chloe"},{"create_time":"2014-08-18 18:43:38","id":"4","name":"Bianca"},{"create_time":"2014-08-18 18:43:38","id":"5","name":"Leo"},{"create_time":"2014-08-18 18:43:38","id":"6","name":"Joy"},{"create_time":"2014-08-18 18:43:38","id":"7","name":"Samuel"}],"total":-1}
```

Gets from test table data(3-6)(in JSON):
```
curl -X GET -H "api_token_id:0" -H "api_token_key:0" "http://localhost:8080/api/test?start=3&limit=4"
```
Returns:
```json
{"data":[{"create_time":"2014-08-18 18:43:38","id":"4","name":"Bianca"},{"create_time":"2014-08-18 18:43:38","id":"5","name":"Leo"},{"create_time":"2014-08-18 18:43:38","id":"6","name":"Joy"},{"create_time":"2014-08-18 18:43:38","id":"7","name":"Samuel"}],"total":-1}
```

Gets from test table data(3-6)(In array, including total rows):
```
curl -X GET -H "api_token_id:0" -H "api_token_key:0" "http://localhost:8080/api/test?start=3&limit=4&array=1&total=1"
```
Returns:
```json
{"data":[["id","name","create_time"],["4","Bianca","2014-08-18 18:43:38"],["5","Leo","2014-08-18 18:43:38"],["6","Joy","2014-08-18 18:43:38"],["7","Samuel","2014-08-18 18:43:38"]],"total":7}
```

Gets from test table record with ID=1:
```
curl -X GET -H "api_token_id:0" -H "api_token_key:0" "http://localhost:8080/api/test/1"
```
Returns:
```json
{"data":{"create_time":"2014-08-18 18:43:38","id":"1","name":"Brian"}}
```

Duplicates from test table the record with ID=1:
```
curl -X COPY -H "api_token_id:0" -H "api_token_key:0" "http://localhost:8080/api/test/1"
```
Returns (ID of new record):
```json
{"data":"d2480a37-88da-492c-a379-a4cdee3049b8"}
```

Deletes from test table the record  with ID=d2480a37-88da-492c-a379-a4cdee3049b8:
```
curl -X DELETE -H "api_token_id:0" -H "api_token_key:0" "http://localhost:8080/api/test/d2480a37-88da-492c-a379-a4cdee3049b8"
```
Returns (Rows affected):
```json
{"data":1}
```

Creates in test table a new record:
```
curl -X POST -H "api_token_id:0" -H "api_token_key:0" -d '{"name": "Elgs"}' "http://localhost:8080/api/test"
```
Returns (ID of new record):
```json
{"data":"192ec8b5-5085-49a1-a9f1-c01a2a682b41"}
```

Updates in the test table the record with ID=192ec8b5-5085-49a1-a9f1-c01a2a682b41:
```
curl -X PUT -H "api_token_id:0" -H "api_token_key:0" -d '{"id":"192ec8b5-5085-49a1-a9f1-c01a2a682b41","name":"Peter"}' "http://localhost:8080/api/test"
```
Returns (Rows affected):
```json
{"data":1}
```

Invalid token:
```
curl -X GET -H "api_token_id:0" -H "api_token_key:1" "http://localhost:8080/api/test"
```
Returns:
```json
{"data":[],"err":"Authentication failed.","total":-1}
```
#Operators

* "eq": " = ",
* "ne": " != ",
* "gt": " > ",
* "lt": " < ",
* "ge": " >= ",
* "le": " <= ",
* "li": " LIKE ",
* "nl": " NOT LIKE ",
* "nu": " IS NULL ",
* "nn": " IS NOT NULL ",
* "rl": " RLIKE ",

#Sample Data Interceptor
```go
package gorest

import (
	"fmt"
)

type EchoDataInterceptor struct {
	*DefaultDataInterceptor
}

func (this *EchoDataInterceptor) BeforeCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	fmt.Println("Here I'm in BeforeCreate")
	//if db, ok := ds.(*sql.DB); ok {
	//	_ = db
	//}
	return true, nil
}
func (this *EchoDataInterceptor) AfterCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) error {
	fmt.Println("Here I'm in AfterCreate")
	return nil
}
func (this *EchoDataInterceptor) BeforeLoad(ds interface{}, context map[string]interface{}, id string) (bool, error) {
	fmt.Println("Here I'm in BeforeLoad")
	return true, nil
}
func (this *EchoDataInterceptor) AfterLoad(ds interface{}, context map[string]interface{}, data map[string]string) error {
	fmt.Println("Here I'm in AfterLoad")
	return nil
}
func (this *EchoDataInterceptor) BeforeUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error) {
	fmt.Println("Here I'm in BeforeUpdate")
	return true, nil
}
func (this *EchoDataInterceptor) AfterUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error {
	fmt.Println("Here I'm in AfterUpdate")
	return nil
}
func (this *EchoDataInterceptor) BeforeDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error) {
	fmt.Println("Here I'm in BeforeDuplicate")
	return true, nil
}
func (this *EchoDataInterceptor) AfterDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error {
	fmt.Println("Here I'm in AfterDuplicate")
	return nil
}
func (this *EchoDataInterceptor) BeforeDelete(ds interface{}, context map[string]interface{}, id string) (bool, error) {
	fmt.Println("Here I'm in BeforeDelete")
	return true, nil
}
func (this *EchoDataInterceptor) AfterDelete(ds interface{}, context map[string]interface{}, id string) error {
	fmt.Println("Here I'm in AfterDelete")
	return nil
}
func (this *EchoDataInterceptor) BeforeListMap(ds interface{}, context map[string]interface{}, where string, order string, start int64, limit int64, includeTotal bool) (bool, error) {
	fmt.Println("Here I'm in BeforeListMap")
	return true, nil
}
func (this *EchoDataInterceptor) AfterListMap(ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	fmt.Println("Here I'm in AfterListMap")
	return nil
}
func (this *EchoDataInterceptor) BeforeListArray(ds interface{}, context map[string]interface{}, where string, order string, start int64, limit int64, includeTotal bool) (bool, error) {
	fmt.Println("Here I'm in BeforeListArray")
	return true, nil
}
func (this *EchoDataInterceptor) AfterListArray(ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	fmt.Println("Here I'm in AfterListArray")
	return nil
}
```
