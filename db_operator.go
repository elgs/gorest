package gorest

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/elgs/gosqljson"
	"github.com/nu7hatch/gouuid"
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
func (this *DbOperator) List(where string, order string) []map[string]string {
	m := gosqljson.QueryDbToMap(this.Db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId))
	return m
}
func (this *DbOperator) Create(data map[string]interface{}) interface{} {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		dataInterceptor.BeforeCreate(this.Db, data)
	}
	// Create the record
	if data["ID"] == nil {
		id, _ := uuid.NewV4()
		data["ID"] = id.String()
	}
	dataLen := len(data)
	values := make([]interface{}, 0, dataLen)
	var buffer bytes.Buffer
	for k, v := range data {
		buffer.WriteString(fmt.Sprint(k, "=?,"))
		values = append(values, v)
	}
	sets := buffer.String()
	sets = sets[0 : len(sets)-1]
	gosqljson.ExecDb(this.Db, fmt.Sprint("INSERT INTO ", this.TableId, " SET ", sets), values...)
	if dataInterceptor != nil {
		dataInterceptor.AfterCreate(this.Db, data)
	}
	return data["ID"]
}
func (this *DbOperator) Update(data map[string]interface{}) {
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
