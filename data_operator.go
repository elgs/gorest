package gorest

type DataOperator interface {
	Load(id string) map[string]string
	List() []map[string]string
	Create(map[string]interface{}) string
	Update([]map[string]interface{})
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
func (this *DefaultDataOperator) Create(map[string]interface{}) string {
	return ""
}
func (this *DefaultDataOperator) Update([]map[string]interface{}) {
}
func (this *DefaultDataOperator) Duplicate(id string) {
}
func (this *DefaultDataOperator) Delete(id string) {
}
