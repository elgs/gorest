package main

var dataInterceptorRegistry = make(map[string]DataInterceptor)

func RegisterDataInterceptor(id string, dataInterceptor DataInterceptor) {
	dataInterceptorRegistry[id] = dataInterceptor
}

func GetDataInterceptor(id string) DataInterceptor {
	return dataInterceptorRegistry[id]
}

type DataInterceptor interface {
	BeforeLoad(id string) bool
	AfterLoad(data map[string]string)
	BeforeCreate(data map[string]interface{}) bool
	AfterCreate(data map[string]interface{})
	BeforeUpdate(oldData map[string]interface{}, data map[string]interface{}) bool
	AfterUpdate(oldData map[string]interface{}, data map[string]interface{})
	BeforeDuplicate(oldData map[string]interface{}, data map[string]interface{}) bool
	AfterDuplicate(oldData map[string]interface{}, data map[string]interface{})
	BeforeDelete(id string) bool
	AfterDelete(id string)
}

type DefaultDataInterceptor struct{}

func (this *DefaultDataInterceptor) BeforeLoad(id string) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterLoad(data map[string]string) {
}
func (this *DefaultDataInterceptor) BeforeCreate(data map[string]interface{}) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterCreate(data map[string]interface{}) {}
func (this *DefaultDataInterceptor) BeforeUpdate(oldData map[string]interface{}, data map[string]interface{}) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterUpdate(oldData map[string]interface{}, data map[string]interface{}) {
}
func (this *DefaultDataInterceptor) BeforeDuplicate(oldData map[string]interface{}, data map[string]interface{}) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterDuplicate(oldData map[string]interface{}, data map[string]interface{}) {
}
func (this *DefaultDataInterceptor) BeforeDelete(id string) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterDelete(id string) {}

type MyDataInterceptor struct {
	*DefaultDataInterceptor
}
