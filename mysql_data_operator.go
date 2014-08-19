package gorest

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"fmt"
	"github.com/elgs/gosqljson"
	"strconv"
)

type MySqlDataOperator struct {
	TableId string
	Db      *sql.DB
}

func (this *MySqlDataOperator) Load(id string) map[string]string {
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
	if m != nil && len(m) == 1 {
		return m[0]
	} else {
		return make(map[string]string, 0)
	}

}
func (this *MySqlDataOperator) List(where string, order string, start int64, limit int64) ([]map[string]string, int64) {
	m := gosqljson.QueryDbToMap(this.Db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId,
			" WHERE 1=1 ", where, " ", order, " LIMIT ?,?"), start, limit)
	c := gosqljson.QueryDbToMap(this.Db, false,
		fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", this.TableId, " WHERE 1=1 ", where))
	cnt, err := strconv.Atoi(c[0]["CNT"])
	if err != nil {
		fmt.Println(err)
	}
	return m, int64(cnt)
}
func (this *MySqlDataOperator) Create(data map[string]interface{}) interface{} {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeCreate(this.Db, data)
		if !ctn {
			return nil
		}
	}
	// Create the record
	if data["ID"] == nil {
		data["ID"] = uuid.New()
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
func (this *MySqlDataOperator) Update(data map[string]interface{}) int64 {
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
func (this *MySqlDataOperator) Duplicate(id string) interface{} {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDuplicate(this.Db, nil, nil)
		if !ctn {
			return nil
		}
	}
	// Duplicate the record
	data := gosqljson.QueryDbToMap(this.Db, false,
		fmt.Sprint("SELECT * FROM ", this.TableId, " WHERE ID=?"), id)
	if data == nil || len(data) != 1 {
		return nil
	}
	newData := make(map[string]interface{}, len(data[0]))
	for k, v := range data[0] {
		newData[fmt.Sprint(k)] = v
	}
	newId := uuid.New()
	newData["ID"] = newId

	newDataLen := len(newData)
	newValues := make([]interface{}, 0, newDataLen)
	var buffer bytes.Buffer
	for k, v := range newData {
		buffer.WriteString(fmt.Sprint(k, "=?,"))
		newValues = append(newValues, v)
	}
	sets := buffer.String()
	sets = sets[0 : len(sets)-1]
	gosqljson.ExecDb(this.Db, fmt.Sprint("INSERT INTO ", this.TableId, " SET ", sets), newValues...)

	if dataInterceptor != nil {
		dataInterceptor.AfterDuplicate(this.Db, nil, nil)
	}
	return newId
}
func (this *MySqlDataOperator) Delete(id string) int64 {
	dataInterceptor := GetDataInterceptor(this.TableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDelete(this.Db, this.TableId)
		if !ctn {
			return 0
		}
	}
	// Delete the record
	rowsAffected, err := gosqljson.ExecDb(this.Db, fmt.Sprint("DELETE FROM ", this.TableId, " WHERE ID=?"), id)
	if err != nil {
		fmt.Println(err)
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterDelete(this.Db, this.TableId)
	}
	return rowsAffected
}
