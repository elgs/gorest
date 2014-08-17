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
		dataInterceptor.BeforeLoad(this.Db, this.TableId)
	}
	// Load the record
	m := gosqljson.QueryDbToMap(this.Db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId, " WHERE ID=?"), id)
	if dataInterceptor != nil {
		dataInterceptor.AfterLoad(this.Db, m[0])
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
		dataInterceptor.BeforeCreate(this.Db, nil)
	}
	// Create the record
	if dataInterceptor != nil {
		dataInterceptor.AfterCreate(this.Db, nil)
	}
	return ""
}
func (this *DbOperator) Update([]map[string]interface{}) {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeUpdate(this.Db, nil, nil)
	}
	// Update the record
	if dataInterceptor != nil {
		dataInterceptor.AfterUpdate(this.Db, nil, nil)
	}
}
func (this *DbOperator) Duplicate(id string) {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeDuplicate(this.Db, nil, nil)
	}
	// Duplicate the record
	if dataInterceptor != nil {
		dataInterceptor.AfterDuplicate(this.Db, nil, nil)
	}
}
func (this *DbOperator) Delete(id string) {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeDelete(this.Db, this.TableId)
	}
	// Delete the record
	if dataInterceptor != nil {
		dataInterceptor.AfterDelete(this.Db, this.TableId)
	}
}
