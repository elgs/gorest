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
	*DefaultDataOperator
	TableId string
	Ds      string
}

func (this *MySqlDataOperator) Load(id string) map[string]string {
	dataInterceptor := GetDataInterceptor(this.TableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeLoad(db, this.TableId)
		if !ctn {
			return nil
		}
	}
	// Load the record
	m, _ := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId, " WHERE ID=?"), id)
	if dataInterceptor != nil {
		dataInterceptor.AfterLoad(db, m[0])
	}
	if m != nil && len(m) == 1 {
		return m[0]
	} else {
		return make(map[string]string, 0)
	}

}
func (this *MySqlDataOperator) ListMap(where string, order string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64) {
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}
	m, _ := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId,
			" WHERE 1=1 ", where, " ", order, " LIMIT ?,?"), start, limit)
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", this.TableId, " WHERE 1=1 ", where))
		if err != nil {
			fmt.Println(err)
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
		}
	}
	return m, int64(cnt)
}
func (this *MySqlDataOperator) ListArray(where string, order string, start int64, limit int64, includeTotal bool) ([][]string, int64) {
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}
	a, _ := gosqljson.QueryDbToArray(db, true,
		fmt.Sprint("SELECT * FROM ", this.TableId,
			" WHERE 1=1 ", where, " ", order, " LIMIT ?,?"), start, limit)
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", this.TableId, " WHERE 1=1 ", where))
		if err != nil {
			fmt.Println(err)
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
		}
	}
	return a, int64(cnt)
}
func (this *MySqlDataOperator) Create(data map[string]interface{}) interface{} {
	dataInterceptor := GetDataInterceptor(this.TableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeCreate(db, data)
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
	gosqljson.ExecDb(db, fmt.Sprint("INSERT INTO ", this.TableId, " SET ", sets), values...)
	if dataInterceptor != nil {
		dataInterceptor.AfterCreate(db, data)
	}
	return data["ID"]
}
func (this *MySqlDataOperator) Update(data map[string]interface{}) int64 {
	dataInterceptor := GetDataInterceptor(this.TableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeUpdate(db, nil, nil)
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
	rowsAffected, err := gosqljson.ExecDb(db, fmt.Sprint("UPDATE ", this.TableId, " SET ", sets, " WHERE ID=?"), values...)
	if err != nil {
		fmt.Println(err)
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterUpdate(db, nil, nil)
	}
	return rowsAffected
}
func (this *MySqlDataOperator) Duplicate(id string) interface{} {
	dataInterceptor := GetDataInterceptor(this.TableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDuplicate(db, nil, nil)
		if !ctn {
			return nil
		}
	}
	// Duplicate the record
	data, _ := gosqljson.QueryDbToMap(db, false,
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
	gosqljson.ExecDb(db, fmt.Sprint("INSERT INTO ", this.TableId, " SET ", sets), newValues...)

	if dataInterceptor != nil {
		dataInterceptor.AfterDuplicate(db, nil, nil)
	}
	return newId
}
func (this *MySqlDataOperator) Delete(id string) int64 {
	dataInterceptor := GetDataInterceptor(this.TableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDelete(db, this.TableId)
		if !ctn {
			return 0
		}
	}
	// Delete the record
	rowsAffected, err := gosqljson.ExecDb(db, fmt.Sprint("DELETE FROM ", this.TableId, " WHERE ID=?"), id)
	if err != nil {
		fmt.Println(err)
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterDelete(db, this.TableId)
	}
	return rowsAffected
}
func getConn(ds string) (*sql.DB, error) {
	db, err := sql.Open("mysql", ds)
	db.SetMaxIdleConns(10)
	return db, err
}
