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
	ds := "user:pass@tcp(host:3306)/test.test"
	dbo := &gorest.MySqlDataOperator{Ds: ds}
	r := &gorest.Gorest{
		Port:      8080,
		UrlPrefix: "api",
		Dbo:       dbo}
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

1. all table names and field names are in upper case;
2. each table has a field `ID` as the primary key;
3. the `ID` field is of char(36);

These restrictions will likely be removed, or be made configurable in the 
future, by improvement or another implementation.

The `DefaultDataInterceptor` is intended to be implemented to extend the 
business logic of the applications. And the `DefaultDataOperator` is intended 
to be implemented by database providers, in order to connect to other databases
or other data sources.

#API
TODO

#Sample Data Interceptor
```go
package gorest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	tableId := "test.TEST"
	RegisterDataInterceptor(tableId, &TDataInterceptor{TableId: tableId})
}

type TDataInterceptor struct {
	*DefaultDataInterceptor
	TableId string
}

func (this *TDataInterceptor) BeforeCreate(ds interface{}, data map[string]interface{}) bool {
	fmt.Println("Here I'm in BeforeCreate")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterCreate(ds interface{}, data map[string]interface{}) {
	fmt.Println("Here I'm in AfterCreate")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeLoad(ds interface{}, id string) bool {
	fmt.Println("Here I'm in BeforeLoad")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterLoad(ds interface{}, data map[string]string) {
	fmt.Println("Here I'm in AfterLoad")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	fmt.Println("Here I'm in BeforeUpdate")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
	fmt.Println("Here I'm in AfterUpdate")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	fmt.Println("Here I'm in BeforeDuplicate")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
	fmt.Println("Here I'm in AfterDuplicate")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeDelete(ds interface{}, id string) bool {
	fmt.Println("Here I'm in BeforeDelete")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterDelete(ds interface{}, id string) {
	fmt.Println("Here I'm in AfterDelete")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeListMap(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeListMap")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterListMap(ds interface{}, data []map[string]string, total int64) {
	fmt.Println("Here I'm in AfterListMap")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeListArray(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeListArray")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterListArray(ds interface{}, data [][]string, total int64) {
	fmt.Println("Here I'm in AfterListArray")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeQueryMap(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeQuerytMap")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterQueryMap(ds interface{}, data []map[string]string, total int64) {
	fmt.Println("Here I'm in AfterQueryMap")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
func (this *TDataInterceptor) BeforeQueryArray(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeQueryArray")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
	return true
}
func (this *TDataInterceptor) AfterQueryArray(ds interface{}, data [][]string, total int64) {
	fmt.Println("Here I'm in AfterQueryArray")
	if db, ok := ds.(*sql.DB); ok {
		_ = db
	}
}
```
