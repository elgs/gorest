package gorest

type DataOperator interface {
	Load(id string) map[string]string
	List() []map[string]string
	Create(map[string]interface{}) string
	Update([]map[string]interface{})
	Duplicate(id string)
	Delete(id string)
}
