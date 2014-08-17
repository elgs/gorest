package gorest

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"github.com/elgs/gosqljson"
	_ "github.com/nu7hatch/gouuid"
	_ "strings"
)

type DbOperator struct {
	TableId string
	Db      *sql.DB
}

func (this *DbOperator) Load(id string) map[string]string {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeLoad(this.TableId)
	}
	// Load the record
	m := gosqljson.QueryDbToMap(this.Db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId, " WHERE ID=?"), id)
	if dataInterceptor != nil {
		dataInterceptor.AfterLoad(nil)
	}
	return m[0]
}
func (this *DbOperator) List() []map[string]string {
	m := gosqljson.QueryDbToMap(this.Db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId))
	return m
}
func (this *DbOperator) Create(map[string]interface{}) string {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeCreate(nil)
	}
	// Create the record
	if dataInterceptor != nil {
		dataInterceptor.AfterCreate(nil)
	}
	return ""
}
func (this *DbOperator) Update([]map[string]interface{}) {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeUpdate(nil, nil)
	}
	// Update the record
	if dataInterceptor != nil {
		dataInterceptor.AfterUpdate(nil, nil)
	}
}
func (this *DbOperator) Duplicate(id string) {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeDuplicate(nil, nil)
	}
	// Duplicate the record
	if dataInterceptor != nil {
		dataInterceptor.AfterDuplicate(nil, nil)
	}
}
func (this *DbOperator) Delete(id string) {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeDelete(this.TableId)
	}
	// Delete the record
	if dataInterceptor != nil {
		dataInterceptor.AfterDelete(this.TableId)
	}
}
