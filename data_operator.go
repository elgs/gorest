package gorest

type DataOperator interface {
	Load(id string) map[string]string
	ListMap(where string, order string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64)
	ListArray(where string, order string, start int64, limit int64, includeTotal bool) ([][]string, int64)
	Create(data map[string]interface{}) interface{}
	Update(data map[string]interface{}) int64
	Duplicate(id string) interface{}
	Delete(id string) int64
	QueryMap(sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64)
	QueryArray(sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([][]string, int64)
}

type DefaultDataOperator struct {
}

func (this *DefaultDataOperator) Load(id string) map[string]string {
	return nil
}
func (this *DefaultDataOperator) ListMap(where string, order string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64) {
	return nil, -1
}
func (this *DefaultDataOperator) ListArray(where string, order string, start int64, limit int64, includeTotal bool) ([][]string, int64) {
	return nil, -1
}
func (this *DefaultDataOperator) Create(data map[string]interface{}) (id interface{}) {
	return nil
}
func (this *DefaultDataOperator) Update(data map[string]interface{}) int64 {
	return 0
}
func (this *DefaultDataOperator) Duplicate(id string) interface{} {
	return nil
}
func (this *DefaultDataOperator) Delete(id string) int64 {
	return 0
}
func (this *DefaultDataOperator) QueryMap(sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64) {
	return nil, -1
}
func (this *DefaultDataOperator) QueryArray(sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([][]string, int64) {
	return nil, -1
}
