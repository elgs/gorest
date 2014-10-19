package gorest

import (
	"database/sql"
)

var dataOperatorRegistry = make(map[string]interface{})

func RegisterDataOperator(id string, dataOperator interface{}) {
	dataOperatorRegistry[id] = dataOperator
}

func GetDataOperator(id string) interface{} {
	return dataOperatorRegistry[id]
}

type DataOperator interface {
	Load(resourceId string, id string, field []string, context map[string]interface{}) (map[string]string, error)
	ListMap(resourceId string, field []string, filter []string, sort string, group string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error)
	ListArray(resourceId string, field []string, filter []string, sort string, group string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error)
	Create(resourceId string, data map[string]interface{}, context map[string]interface{}) (interface{}, error)
	Update(resourceId string, data map[string]interface{}, context map[string]interface{}) (int64, error)
	Duplicate(resourceId string, id string, context map[string]interface{}) (interface{}, error)
	Delete(resourceId string, id string, context map[string]interface{}) (int64, error)
	GetConn() (*sql.DB, error)
}

type DefaultDataOperator struct {
}

func (this *DefaultDataOperator) Load(resourceId string, id string, field []string, context map[string]interface{}) (map[string]string, error) {
	return nil, nil
}
func (this *DefaultDataOperator) ListMap(resourceId string, field []string, filter []string, sort string, group string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error) {
	return nil, -1, nil
}
func (this *DefaultDataOperator) ListArray(resourceId string, field []string, filter []string, sort string, group string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error) {
	return nil, -1, nil
}
func (this *DefaultDataOperator) Create(resourceId string, data map[string]interface{}, context map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (this *DefaultDataOperator) Update(resourceId string, data map[string]interface{}, context map[string]interface{}) (int64, error) {
	return 0, nil
}
func (this *DefaultDataOperator) Duplicate(resourceId string, id string, context map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (this *DefaultDataOperator) Delete(resourceId string, id string, context map[string]interface{}) (int64, error) {
	return 0, nil
}
func (this *DefaultDataOperator) GetConn() (*sql.DB, error) {
	return nil, nil
}
