package gorest

type DataOperator interface {
	Load(tableId string, id string, context map[string]interface{}) (map[string]string, error)
	ListMap(tableId string, filter []string, sort string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error)
	ListArray(tableId string, filter []string, sort string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error)
	Create(tableId string, data map[string]interface{}, context map[string]interface{}) (interface{}, error)
	Update(tableId string, data map[string]interface{}, context map[string]interface{}) (int64, error)
	Duplicate(tableId string, id string, context map[string]interface{}) (interface{}, error)
	Delete(tableId string, id string, context map[string]interface{}) (int64, error)
	QueryMap(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error)
	QueryArray(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error)
}

type DefaultDataOperator struct {
}

func (this *DefaultDataOperator) Load(tableId string, id string, context map[string]interface{}) (map[string]string, error) {
	return nil, nil
}
func (this *DefaultDataOperator) ListMap(tableId string, filter []string, sort string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error) {
	return nil, -1, nil
}
func (this *DefaultDataOperator) ListArray(tableId string, filter []string, sort string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error) {
	return nil, -1, nil
}
func (this *DefaultDataOperator) Create(tableId string, data map[string]interface{}, context map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (this *DefaultDataOperator) Update(tableId string, data map[string]interface{}, context map[string]interface{}) (int64, error) {
	return 0, nil
}
func (this *DefaultDataOperator) Duplicate(tableId string, id string, context map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (this *DefaultDataOperator) Delete(tableId string, id string, context map[string]interface{}) (int64, error) {
	return 0, nil
}
func (this *DefaultDataOperator) QueryMap(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error) {
	return nil, -1, nil
}
func (this *DefaultDataOperator) QueryArray(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error) {
	return nil, -1, nil
}
