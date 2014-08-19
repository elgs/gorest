package gorest

type DataOperator interface {
	Load(tableId string, id string) map[string]string
	ListMap(tableId string, where string, order string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64)
	ListArray(tableId string, where string, order string, start int64, limit int64, includeTotal bool) ([][]string, int64)
	Create(tableId string, data map[string]interface{}) interface{}
	Update(tableId string, data map[string]interface{}) int64
	Duplicate(tableId string, id string) interface{}
	Delete(tableId string, id string) int64
	QueryMap(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64)
	QueryArray(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([][]string, int64)
}

type DefaultDataOperator struct {
}

func (this *DefaultDataOperator) Load(tableId string, id string) map[string]string {
	return nil
}
func (this *DefaultDataOperator) ListMap(tableId string, where string, order string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64) {
	return nil, -1
}
func (this *DefaultDataOperator) ListArray(tableId string, where string, order string, start int64, limit int64, includeTotal bool) ([][]string, int64) {
	return nil, -1
}
func (this *DefaultDataOperator) Create(tableId string, data map[string]interface{}) (id interface{}) {
	return nil
}
func (this *DefaultDataOperator) Update(tableId string, data map[string]interface{}) int64 {
	return 0
}
func (this *DefaultDataOperator) Duplicate(tableId string, id string) interface{} {
	return nil
}
func (this *DefaultDataOperator) Delete(tableId string, id string) int64 {
	return 0
}
func (this *DefaultDataOperator) QueryMap(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64) {
	return nil, -1
}
func (this *DefaultDataOperator) QueryArray(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([][]string, int64) {
	return nil, -1
}
