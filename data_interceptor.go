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
	BeforeLoad(ds interface{}, context map[string]interface{}, id string) (bool, error)
	AfterLoad(ds interface{}, context map[string]interface{}, data map[string]string) error
	BeforeCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error)
	AfterCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) error
	BeforeUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error)
	AfterUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error
	BeforeDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error)
	AfterDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error
	BeforeDelete(ds interface{}, context map[string]interface{}, id string) (bool, error)
	AfterDelete(ds interface{}, context map[string]interface{}, id string) error
	BeforeListMap(ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterListMap(ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error
	BeforeListArray(ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterListArray(ds interface{}, context map[string]interface{}, data [][]string, total int64) error
	BeforeQueryMap(ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterQueryMap(ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error
	BeforeQueryArray(ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterQueryArray(ds interface{}, context map[string]interface{}, data [][]string, total int64) error
}

type DefaultDataInterceptor struct{}

func (this *DefaultDataInterceptor) BeforeLoad(ds interface{}, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterLoad(ds interface{}, context map[string]interface{}, data map[string]string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeDelete(ds interface{}, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterDelete(ds interface{}, context map[string]interface{}, id string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeListMap(ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterListMap(ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeListArray(ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterListArray(ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeQueryMap(ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterQueryMap(ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeQueryArray(ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterQueryArray(ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	return nil
}
