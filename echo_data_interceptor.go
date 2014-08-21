package gorest

import (
	"fmt"
)

type EchoDataInterceptor struct {
	*DefaultDataInterceptor
}

func (this *EchoDataInterceptor) BeforeCreate(ds interface{}, data map[string]interface{}) bool {
	fmt.Println("Here I'm in BeforeCreate")
	//if db, ok := ds.(*sql.DB); ok {
	//	_ = db
	//}
	return true
}
func (this *EchoDataInterceptor) AfterCreate(ds interface{}, data map[string]interface{}) {
	fmt.Println("Here I'm in AfterCreate")
}
func (this *EchoDataInterceptor) BeforeLoad(ds interface{}, id string) bool {
	fmt.Println("Here I'm in BeforeLoad")
	return true
}
func (this *EchoDataInterceptor) AfterLoad(ds interface{}, data map[string]string) {
	fmt.Println("Here I'm in AfterLoad")
}
func (this *EchoDataInterceptor) BeforeUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	fmt.Println("Here I'm in BeforeUpdate")
	return true
}
func (this *EchoDataInterceptor) AfterUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
	fmt.Println("Here I'm in AfterUpdate")
}
func (this *EchoDataInterceptor) BeforeDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	fmt.Println("Here I'm in BeforeDuplicate")
	return true
}
func (this *EchoDataInterceptor) AfterDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
	fmt.Println("Here I'm in AfterDuplicate")
}
func (this *EchoDataInterceptor) BeforeDelete(ds interface{}, id string) bool {
	fmt.Println("Here I'm in BeforeDelete")
	return true
}
func (this *EchoDataInterceptor) AfterDelete(ds interface{}, id string) {
	fmt.Println("Here I'm in AfterDelete")
}
func (this *EchoDataInterceptor) BeforeListMap(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeListMap")
	return true
}
func (this *EchoDataInterceptor) AfterListMap(ds interface{}, data []map[string]string, total int64) {
	fmt.Println("Here I'm in AfterListMap")
}
func (this *EchoDataInterceptor) BeforeListArray(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeListArray")
	return true
}
func (this *EchoDataInterceptor) AfterListArray(ds interface{}, data [][]string, total int64) {
	fmt.Println("Here I'm in AfterListArray")
}
func (this *EchoDataInterceptor) BeforeQueryMap(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeQuerytMap")
	return true
}
func (this *EchoDataInterceptor) AfterQueryMap(ds interface{}, data []map[string]string, total int64) {
	fmt.Println("Here I'm in AfterQueryMap")
}
func (this *EchoDataInterceptor) BeforeQueryArray(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool {
	fmt.Println("Here I'm in BeforeQueryArray")
	return true
}
func (this *EchoDataInterceptor) AfterQueryArray(ds interface{}, data [][]string, total int64) {
	fmt.Println("Here I'm in AfterQueryArray")
}
