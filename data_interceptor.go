package gorest

var dataInterceptorRegistry = make(map[string]DataInterceptor)

func RegisterDataInterceptor(id string, dataInterceptor DataInterceptor) {
	dataInterceptorRegistry[id] = dataInterceptor
}

func GetDataInterceptor(id string) DataInterceptor {
	return dataInterceptorRegistry[id]
}

type DataInterceptor interface {
	BeforeLoad(ds interface{}, id string) bool
	AfterLoad(ds interface{}, data map[string]string)
	BeforeCreate(ds interface{}, data map[string]interface{}) bool
	AfterCreate(ds interface{}, data map[string]interface{})
	BeforeUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool
	AfterUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{})
	BeforeDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool
	AfterDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{})
	BeforeDelete(ds interface{}, id string) bool
	AfterDelete(ds interface{}, id string)
}

type DefaultDataInterceptor struct{}

func (this *DefaultDataInterceptor) BeforeLoad(ds interface{}, id string) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterLoad(ds interface{}, data map[string]string) {
}
func (this *DefaultDataInterceptor) BeforeCreate(ds interface{}, data map[string]interface{}) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterCreate(ds interface{}, data map[string]interface{}) {}
func (this *DefaultDataInterceptor) BeforeUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterUpdate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
}
func (this *DefaultDataInterceptor) BeforeDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterDuplicate(ds interface{}, oldData map[string]interface{}, data map[string]interface{}) {
}
func (this *DefaultDataInterceptor) BeforeDelete(ds interface{}, id string) bool {
	return false
}
func (this *DefaultDataInterceptor) AfterDelete(ds interface{}, id string) {}
