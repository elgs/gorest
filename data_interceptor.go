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
	BeforeLoad(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error)
	AfterLoad(resourceId string, ds interface{}, context map[string]interface{}, data map[string]string) error
	BeforeCreate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error)
	AfterCreate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) error
	BeforeUpdate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error)
	AfterUpdate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) error
	BeforeDuplicate(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error)
	AfterDuplicate(resourceId string, ds interface{}, context map[string]interface{}, oldId string, newId string) error
	BeforeDelete(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error)
	AfterDelete(resourceId string, ds interface{}, context map[string]interface{}, id string) error
	BeforeListMap(resourceId string, ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterListMap(resourceId string, ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error
	BeforeListArray(resourceId string, ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterListArray(resourceId string, ds interface{}, context map[string]interface{}, data [][]string, total int64) error
	BeforeQueryMap(resourceId string, ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterQueryMap(resourceId string, ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error
	BeforeQueryArray(resourceId string, ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterQueryArray(resourceId string, ds interface{}, context map[string]interface{}, data [][]string, total int64) error
}

type DefaultDataInterceptor struct{}

func (this *DefaultDataInterceptor) BeforeLoad(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterLoad(resourceId string, ds interface{}, context map[string]interface{}, data map[string]string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeCreate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterCreate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeUpdate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterUpdate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeDuplicate(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterDuplicate(resourceId string, ds interface{}, context map[string]interface{}, oldId string, newId string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeDelete(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterDelete(resourceId string, ds interface{}, context map[string]interface{}, id string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeListMap(resourceId string, ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterListMap(resourceId string, ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeListArray(resourceId string, ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterListArray(resourceId string, ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeQueryMap(resourceId string, ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterQueryMap(resourceId string, ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeQueryArray(resourceId string, ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterQueryArray(resourceId string, ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	return nil
}
