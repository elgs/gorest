package gorest

type DataOperator interface {
	Load(id string) map[string]string
	List() []map[string]string
	Create(data map[string]interface{}) string
	Update(data map[string]interface{})
	Duplicate(id string)
	Delete(id string)
}

type DefaultDataOperator struct {
}

func (this *DefaultDataOperator) Load(id string) map[string]string {
	return nil
}
func (this *DefaultDataOperator) List() []map[string]string {
	return nil
}
func (this *DefaultDataOperator) Create(data map[string]interface{}) (id interface{}) {
	return ""
}
func (this *DefaultDataOperator) Update(data map[string]interface{}) {
}
func (this *DefaultDataOperator) Duplicate(id string) {
}
func (this *DefaultDataOperator) Delete(id string) {
}
