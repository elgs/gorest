package gorest

import (
	"database/sql"
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
	BeforeLoad(resourceId string, db *sql.DB, field []string, context map[string]interface{}, id string) (bool, error)
	AfterLoad(resourceId string, db *sql.DB, field []string, context map[string]interface{}, data map[string]string) error
	BeforeCreate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) (bool, error)
	AfterCreate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) error
	BeforeUpdate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) (bool, error)
	AfterUpdate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) error
	BeforeDuplicate(resourceId string, db *sql.DB, context map[string]interface{}, id string) (bool, error)
	AfterDuplicate(resourceId string, db *sql.DB, context map[string]interface{}, oldId string, newId string) error
	BeforeDelete(resourceId string, db *sql.DB, context map[string]interface{}, id string) (bool, error)
	AfterDelete(resourceId string, db *sql.DB, context map[string]interface{}, id string) error
	BeforeListMap(resourceId string, db *sql.DB, field []string, context map[string]interface{}, filter *string, sort *string, group *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterListMap(resourceId string, db *sql.DB, field []string, context map[string]interface{}, data []map[string]string, total int64) error
	BeforeListArray(resourceId string, db *sql.DB, field []string, context map[string]interface{}, filter *string, sort *string, group *string, start int64, limit int64, includeTotal bool) (bool, error)
	AfterListArray(resourceId string, db *sql.DB, field []string, context map[string]interface{}, data [][]string, total int64) error
}

type DefaultDataInterceptor struct{}

func (this *DefaultDataInterceptor) BeforeLoad(resourceId string, db *sql.DB, field []string, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterLoad(resourceId string, db *sql.DB, field []string, context map[string]interface{}, data map[string]string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeCreate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterCreate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeUpdate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterUpdate(resourceId string, db *sql.DB, context map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeDuplicate(resourceId string, db *sql.DB, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterDuplicate(resourceId string, db *sql.DB, context map[string]interface{}, oldId string, newId string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeDelete(resourceId string, db *sql.DB, context map[string]interface{}, id string) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterDelete(resourceId string, db *sql.DB, context map[string]interface{}, id string) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeListMap(resourceId string, db *sql.DB, field []string, context map[string]interface{}, filter *string, sort *string, group *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterListMap(resourceId string, db *sql.DB, field []string, context map[string]interface{}, data []map[string]string, total int64) error {
	return nil
}
func (this *DefaultDataInterceptor) BeforeListArray(resourceId string, db *sql.DB, field []string, context map[string]interface{}, filter *string, sort *string, group *string, start int64, limit int64, includeTotal bool) (bool, error) {
	return true, nil
}
func (this *DefaultDataInterceptor) AfterListArray(resourceId string, db *sql.DB, field []string, context map[string]interface{}, data [][]string, total int64) error {
	return nil
}
