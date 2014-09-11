package gorest

import (
	"fmt"
)

type EchoDataInterceptor struct {
	*DefaultDataInterceptor
}

func (this *EchoDataInterceptor) BeforeCreate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	fmt.Println("Here I'm in BeforeCreate")
	//if db, ok := ds.(*sql.DB); ok {
	//	_ = db
	//}
	return true, nil
}
func (this *EchoDataInterceptor) AfterCreate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) error {
	fmt.Println("Here I'm in AfterCreate")
	return nil
}
func (this *EchoDataInterceptor) BeforeLoad(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error) {
	fmt.Println("Here I'm in BeforeLoad")
	return true, nil
}
func (this *EchoDataInterceptor) AfterLoad(resourceId string, ds interface{}, context map[string]interface{}, data map[string]string) error {
	fmt.Println("Here I'm in AfterLoad")
	return nil
}
func (this *EchoDataInterceptor) BeforeUpdate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	fmt.Println("Here I'm in BeforeUpdate")
	return true, nil
}
func (this *EchoDataInterceptor) AfterUpdate(resourceId string, ds interface{}, context map[string]interface{}, data map[string]interface{}) error {
	fmt.Println("Here I'm in AfterUpdate")
	return nil
}
func (this *EchoDataInterceptor) BeforeDuplicate(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error) {
	fmt.Println("Here I'm in BeforeDuplicate")
	return true, nil
}
func (this *EchoDataInterceptor) AfterDuplicate(resourceId string, ds interface{}, context map[string]interface{}, oldId string, newId string) error {
	fmt.Println("Here I'm in AfterDuplicate")
	return nil
}
func (this *EchoDataInterceptor) BeforeDelete(resourceId string, ds interface{}, context map[string]interface{}, id string) (bool, error) {
	fmt.Println("Here I'm in BeforeDelete")
	return true, nil
}
func (this *EchoDataInterceptor) AfterDelete(resourceId string, ds interface{}, context map[string]interface{}, id string) error {
	fmt.Println("Here I'm in AfterDelete")
	return nil
}
func (this *EchoDataInterceptor) BeforeListMap(resourceId string, ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	fmt.Println("Here I'm in BeforeListMap")
	return true, nil
}
func (this *EchoDataInterceptor) AfterListMap(resourceId string, ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	fmt.Println("Here I'm in AfterListMap")
	return nil
}
func (this *EchoDataInterceptor) BeforeListArray(resourceId string, ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	fmt.Println("Here I'm in BeforeListArray")
	return true, nil
}
func (this *EchoDataInterceptor) AfterListArray(resourceId string, ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	fmt.Println("Here I'm in AfterListArray")
	return nil
}
func (this *EchoDataInterceptor) BeforeQueryMap(resourceId string, ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	fmt.Println("Here I'm in BeforeQuerytMap")
	return true, nil
}
func (this *EchoDataInterceptor) AfterQueryMap(resourceId string, ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	fmt.Println("Here I'm in AfterQueryMap")
	return nil
}
func (this *EchoDataInterceptor) BeforeQueryArray(resourceId string, ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	fmt.Println("Here I'm in BeforeQueryArray")
	return true, nil
}
func (this *EchoDataInterceptor) AfterQueryArray(resourceId string, ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	fmt.Println("Here I'm in AfterQueryArray")
	return nil
}
