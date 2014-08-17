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
		ctn := dataInterceptor.BeforeLoad(this.Db, this.TableId)
		if !ctn {
			return nil
		}
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
		ctn := dataInterceptor.BeforeCreate(this.Db, data)
		if !ctn {
			return nil
		}
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
func (this *DbOperator) Update(data map[string]interface{}) int64 {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeUpdate(this.Db, nil, nil)
		if !ctn {
			return 0
		}
	}
	// Update the record
	id := data["ID"]
	if id == nil {
		fmt.Println("ID is not found.")
		return 0
	}
	delete(data, "ID")
	dataLen := len(data)
	values := make([]interface{}, 0, dataLen)
	var buffer bytes.Buffer
	for k, v := range data {
		buffer.WriteString(fmt.Sprint(k, "=?,"))
		values = append(values, v)
	}
	values = append(values, id)
	sets := buffer.String()
	sets = sets[0 : len(sets)-1]
	rowsAffected, err := gosqljson.ExecDb(this.Db, fmt.Sprint("UPDATE ", this.TableId, " SET ", sets, " WHERE ID=?"), values...)
	if err != nil {
		fmt.Println(err)
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterUpdate(this.Db, nil, nil)
	}
	return rowsAffected
}
func (this *DbOperator) Duplicate(id string) interface{} {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDuplicate(this.Db, nil, nil)
		if !ctn {
			return nil
		}
	}
	// Duplicate the record
	if dataInterceptor != nil {
		dataInterceptor.AfterDuplicate(this.Db, nil, nil)
	}
	return nil
}
func (this *DbOperator) Delete(id string) int64 {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDelete(this.Db, this.TableId)
		if !ctn {
			return 0
		}
	}
	// Delete the record
	if dataInterceptor != nil {
		dataInterceptor.AfterDelete(this.Db, this.TableId)
	}
	return 0
}
