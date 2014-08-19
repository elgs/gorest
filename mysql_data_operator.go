package gorest

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"fmt"
	"github.com/elgs/gosqljson"
	"strconv"
	"strings"
)

type MySqlDataOperator struct {
	*DefaultDataOperator
	Ds string
}

func (this *MySqlDataOperator) Load(tableId string, id string) map[string]string {
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return nil
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeLoad(db, tableId)
		if !ctn {
			return nil
		}
	}
	// Load the record
	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
	if err != nil {
		fmt.Println()
		return nil
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterLoad(db, m[0])
	}
	if m != nil && len(m) == 1 {
		return m[0]
	} else {
		return make(map[string]string, 0)
	}

}
func (this *MySqlDataOperator) ListMap(tableId string, where string, order string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64) {
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return nil, -1
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeListMap(db, where, order, start, limit, includeTotal)
		if !ctn {
			return nil, -1
		}
	}
	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", tableId,
			" WHERE 1=1 ", where, " ", order, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println()
		return nil, -1
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", tableId, " WHERE 1=1 ", where))
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterListMap(db, m, int64(cnt))
	}
	return m, int64(cnt)
}
func (this *MySqlDataOperator) ListArray(tableId string, where string, order string, start int64, limit int64, includeTotal bool) ([][]string, int64) {
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return nil, -1
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeListArray(db, where, order, start, limit, includeTotal)
		if !ctn {
			return nil, -1
		}
	}
	a, err := gosqljson.QueryDbToArray(db, true,
		fmt.Sprint("SELECT * FROM ", tableId,
			" WHERE 1=1 ", where, " ", order, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return nil, -1
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", tableId, " WHERE 1=1 ", where))
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterListArray(db, a, int64(cnt))
	}
	return a, int64(cnt)
}
func (this *MySqlDataOperator) Create(tableId string, data map[string]interface{}) interface{} {
	dataInterceptor := GetDataInterceptor(tableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return nil
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
	gosqljson.ExecDb(db, fmt.Sprint("INSERT INTO ", tableId, " SET ", sets), values...)
	if dataInterceptor != nil {
		dataInterceptor.AfterCreate(db, data)
	}
	return data["ID"]
}
func (this *MySqlDataOperator) Update(tableId string, data map[string]interface{}) int64 {
	dataInterceptor := GetDataInterceptor(tableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return -1
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
	rowsAffected, err := gosqljson.ExecDb(db, fmt.Sprint("UPDATE ", tableId, " SET ", sets, " WHERE ID=?"), values...)
	if err != nil {
		fmt.Println(err)
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterUpdate(db, nil, nil)
	}
	return rowsAffected
}
func (this *MySqlDataOperator) Duplicate(tableId string, id string) interface{} {
	dataInterceptor := GetDataInterceptor(tableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return nil
	}
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDuplicate(db, nil, nil)
		if !ctn {
			return nil
		}
	}
	// Duplicate the record
	data, _ := gosqljson.QueryDbToMap(db, false,
		fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
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
	gosqljson.ExecDb(db, fmt.Sprint("INSERT INTO ", tableId, " SET ", sets), newValues...)

	if dataInterceptor != nil {
		dataInterceptor.AfterDuplicate(db, nil, nil)
	}
	return newId
}
func (this *MySqlDataOperator) Delete(tableId string, id string) int64 {
	dataInterceptor := GetDataInterceptor(tableId)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return -1
	}
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDelete(db, tableId)
		if !ctn {
			return 0
		}
	}
	// Delete the record
	rowsAffected, err := gosqljson.ExecDb(db, fmt.Sprint("DELETE FROM ", tableId, " WHERE ID=?"), id)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterDelete(db, tableId)
	}
	return rowsAffected
}
func (this *MySqlDataOperator) QueryMap(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([]map[string]string, int64) {
	if !isSelect(sqlSelect) || !isSelect(sqlSelectCount) {
		return nil, -1
	}
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return nil, -1
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeQueryMap(db, sqlSelect, sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return nil, -1
		}
	}
	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint(sqlSelect, " LIMIT ?,?"), start, limit)
	cnt := -1
	if err != nil {
		fmt.Println()
		return nil, -1
	}
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false, sqlSelectCount)
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
		for _, v := range c[0] {
			cnt, err = strconv.Atoi(v)
		}
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterQueryMap(db, m, int64(cnt))
	}
	return m, int64(cnt)
}
func (this *MySqlDataOperator) QueryArray(tableId string, sqlSelect string, sqlSelectCount string, start int64, limit int64, includeTotal bool) ([][]string, int64) {
	if !isSelect(sqlSelect) || !isSelect(sqlSelectCount) {
		return nil, -1
	}
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println()
		return nil, -1
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeQueryArray(db, sqlSelect, sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return nil, -1
		}
	}
	a, err := gosqljson.QueryDbToArray(db, true,
		fmt.Sprint(sqlSelect, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return nil, -1
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false, sqlSelectCount)
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
		for _, v := range c[0] {
			cnt, err = strconv.Atoi(v)
		}
		if err != nil {
			fmt.Println(err)
			return nil, -1
		}
	}
	if dataInterceptor != nil {
		dataInterceptor.AfterQueryArray(db, a, int64(cnt))
	}
	return a, int64(cnt)
}

func isSelect(sqlSelect string) bool {
	return strings.HasPrefix(strings.ToUpper(sqlSelect), "SELET ")
}

func getConn(ds string) (*sql.DB, error) {
	db, err := sql.Open("mysql", ds)
	db.SetMaxIdleConns(10)
	return db, err
}
