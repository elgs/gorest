package gorest

import (
	"strings"
)

var dataInterceptorRegistry = make(map[string]DataInterceptor)

func RegisterDataInterceptor(id string, dataInterceptor DataInterceptor) {
	dataInterceptorRegistry[strings.ToUpper(id)] = dataInterceptor
}

func GetDataInterceptor(id string) DataInterceptor {
	return dataInterceptorRegistry[strings.ToUpper(id)]
}

var GlobalDataInterceptorRegistry = make([]DataInterceptor, 0)

func RegisterGlobalDataInterceptor(globalDataInterceptor DataInterceptor) {
	GlobalDataInterceptorRegistry = append(GlobalDataInterceptorRegistry, globalDataInterceptor)
}

type DataInterceptor interface {
	BeforeLoad(ds interface{}, id string) bool
	AfterLoad(ds interface{}, data map[string]string)
	BeforeCreate(ds interface{}, data map[string]interface{}) bool
	AfterCreate(ds interface{}, data map[string]interface{})
	BeforeUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool
	AfterUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{})
	BeforeDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool
	AfterDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{})
	BeforeDelete(ds interface{}, id string) bool
	AfterDelete(ds interface{}, id string)
	BeforeListMap(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool
	AfterListMap(ds interface{}, data []map[string]string, total int64)
	BeforeListArray(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool
	AfterListArray(ds interface{}, data [][]string, total int64)
	BeforeQueryMap(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool
	AfterQueryMap(ds interface{}, data []map[string]string, total int64)
	BeforeQueryArray(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool
	AfterQueryArray(ds interface{}, data [][]string, total int64)
}

type DefaultDataInterceptor struct{}

func (this *DefaultDataInterceptor) BeforeLoad(ds interface{}, id string) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterLoad(ds interface{}, data map[string]string) {}
func (this *DefaultDataInterceptor) BeforeCreate(ds interface{}, data map[string]interface{}) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterCreate(ds interface{}, data map[string]interface{}) {}
func (this *DefaultDataInterceptor) BeforeUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
}
func (this *DefaultDataInterceptor) BeforeDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
}
func (this *DefaultDataInterceptor) BeforeDelete(ds interface{}, id string) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterDelete(ds interface{}, id string) {}
func (this *DefaultDataInterceptor) BeforeListMap(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterListMap(ds interface{}, data []map[string]string, total int64) {
}
func (this *DefaultDataInterceptor) BeforeListArray(ds interface{}, where string, order string, start int64, limit int64, includeTotal bool) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterListArray(ds interface{}, data [][]string, total int64) {}
func (this *DefaultDataInterceptor) BeforeQueryMap(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterQueryMap(ds interface{}, data []map[string]string, total int64) {
}
func (this *DefaultDataInterceptor) BeforeQueryArray(ds interface{}, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) bool {
	return true
}
func (this *DefaultDataInterceptor) AfterQueryArray(ds interface{}, data [][]string, total int64) {}
