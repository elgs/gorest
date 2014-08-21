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

func (this *MySqlDataOperator) Load(tableId string, id string, context map[string]interface{}) map[string]string {
	ret := make(map[string]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeLoad(db, context, tableId)
		if !ctn {
			return ret
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeLoad(db, context, tableId)
		if !ctn {
			return ret
		}
	}

	// Load the record
	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
	if err != nil {
		fmt.Println(err)
		return ret
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterLoad(db, context, m[0])
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterLoad(db, context, m[0])
	}

	if m != nil && len(m) == 1 {
		return m[0]
	} else {
		return ret
	}

}
func (this *MySqlDataOperator) ListMap(tableId string, where string, order string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64) {
	ret := make([]map[string]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeListMap(db, context, where, order, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeListMap(db, context, where, order, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}

	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", tableId,
			" WHERE 1=1 ", where, " ", order, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", tableId, " WHERE 1=1 ", where))
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterListMap(db, context, m, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterListMap(db, context, m, int64(cnt))
	}

	return m, int64(cnt)
}
func (this *MySqlDataOperator) ListArray(tableId string, where string, order string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64) {
	ret := make([][]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeListArray(db, context, where, order, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeListArray(db, context, where, order, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}

	a, err := gosqljson.QueryDbToArray(db, true,
		fmt.Sprint("SELECT * FROM ", tableId,
			" WHERE 1=1 ", where, " ", order, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", tableId, " WHERE 1=1 ", where))
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterListArray(db, context, a, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterListArray(db, context, a, int64(cnt))
	}

	return a, int64(cnt)
}
func (this *MySqlDataOperator) Create(tableId string, data map[string]interface{}, context map[string]interface{}) interface{} {
	tableId = normalizeTableId(tableId, this.Ds)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeCreate(db, context, data)
		if !ctn {
			return ""
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeCreate(db, context, data)
		if !ctn {
			return ""
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
		dataInterceptor.AfterCreate(db, context, data)
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterCreate(db, context, data)
	}

	return data["ID"]
}
func (this *MySqlDataOperator) Update(tableId string, data map[string]interface{}, context map[string]interface{}) int64 {
	tableId = normalizeTableId(tableId, this.Ds)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeUpdate(db, context, nil, nil)
		if !ctn {
			return 0
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeUpdate(db, context, nil, nil)
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
		return -1
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterUpdate(db, context, nil, nil)
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterUpdate(db, context, nil, nil)
	}

	return rowsAffected
}
func (this *MySqlDataOperator) Duplicate(tableId string, id string, context map[string]interface{}) interface{} {
	tableId = normalizeTableId(tableId, this.Ds)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeDuplicate(db, context, nil, nil)
		if !ctn {
			return ""
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDuplicate(db, context, nil, nil)
		if !ctn {
			return ""
		}
	}

	// Duplicate the record
	data, _ := gosqljson.QueryDbToMap(db, false,
		fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
	if data == nil || len(data) != 1 {
		return ""
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
		dataInterceptor.AfterDuplicate(db, context, nil, nil)
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterDuplicate(db, context, nil, nil)
	}

	return newId
}
func (this *MySqlDataOperator) Delete(tableId string, id string, context map[string]interface{}) int64 {
	tableId = normalizeTableId(tableId, this.Ds)
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeDelete(db, context, tableId)
		if !ctn {
			return 0
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeDelete(db, context, tableId)
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
		dataInterceptor.AfterDelete(db, context, tableId)
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterDelete(db, context, tableId)
	}

	return rowsAffected
}
func (this *MySqlDataOperator) QueryMap(tableId string, sqlSelect string, sqlSelectCount string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64) {
	ret := make([]map[string]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	if !isSelect(sqlSelect) {
		return ret, -1
	}
	if includeTotal && !isSelect(sqlSelectCount) {
		return ret, -1
	}
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeQueryMap(db, context, sqlSelect, sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeQueryMap(db, context, sqlSelect, sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}

	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint(sqlSelect, " LIMIT ?,?"), start, limit)
	cnt := -1
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false, sqlSelectCount)
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
		for _, v := range c[0] {
			cnt, err = strconv.Atoi(v)
		}
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterQueryMap(db, context, m, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterQueryMap(db, context, m, int64(cnt))
	}

	return m, int64(cnt)
}
func (this *MySqlDataOperator) QueryArray(tableId string, sqlSelect string, sqlSelectCount string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64) {
	ret := make([][]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	if !isSelect(sqlSelect) {
		return ret, -1
	}
	if includeTotal && !isSelect(sqlSelectCount) {
		return ret, -1
	}
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn := globalDataInterceptor.BeforeQueryArray(db, context, sqlSelect, sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn := dataInterceptor.BeforeQueryArray(db, context, sqlSelect, sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1
		}
	}

	a, err := gosqljson.QueryDbToArray(db, true,
		fmt.Sprint(sqlSelect, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false, sqlSelectCount)
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
		for _, v := range c[0] {
			cnt, err = strconv.Atoi(v)
		}
		if err != nil {
			fmt.Println(err)
			return ret, -1
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterQueryArray(db, context, a, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterQueryArray(db, context, a, int64(cnt))
	}

	return a, int64(cnt)
}

func isSelect(sqlSelect string) bool {
	return strings.HasPrefix(strings.ToUpper(sqlSelect), "SELECT ")
}

func getConn(ds string) (*sql.DB, error) {
	db, err := sql.Open("mysql", ds)
	db.SetMaxIdleConns(10)
	return db, err
}

func extractDbNameFromDs(ds string) string {
	a := strings.LastIndex(ds, "/")
	b := ds[a+1:]
	c := strings.Index(b, "?")
	if c < 0 {
		return b
	}
	return b[:c]
}

func normalizeTableId(tableId string, ds string) string {
	if strings.Contains(tableId, ".") {
		a := strings.Split(tableId, ".")
		return fmt.Sprint(a[0], ".", a[1])
	}
	db := extractDbNameFromDs(ds)
	return strings.Replace(fmt.Sprint(db, ".", tableId), "'", "''", -1)
}
