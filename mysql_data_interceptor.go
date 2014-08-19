package gorest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	tableId := "test.test"
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
