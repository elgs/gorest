package gorest

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"errors"
	"fmt"
	"github.com/elgs/gosqljson"
	"strconv"
	"strings"
)

type MySqlDataOperator struct {
	*DefaultDataOperator
	Ds string
}

func (this *MySqlDataOperator) Load(tableId string, id string, context map[string]interface{}) (map[string]string, error) {
	ret := make(map[string]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, err
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeLoad(db, context, tableId)
		if !ctn {
			return ret, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeLoad(db, context, tableId)
		if !ctn {
			return ret, err
		}
	}

	// Load the record
	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
	if err != nil {
		fmt.Println(err)
		return ret, err
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterLoad(db, context, m[0])
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterLoad(db, context, m[0])
	}

	if m != nil && len(m) == 1 {
		return m[0], err
	} else {
		return ret, err
	}

}
func (this *MySqlDataOperator) ListMap(tableId string, filter []string, sort string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error) {
	ret := make([]map[string]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}

	sort = parseSort(sort)
	where := parseFilter(filter)
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeListMap(db, context, &where, &sort, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeListMap(db, context, &where, &sort, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}

	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint("SELECT * FROM ", tableId, where, sort, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", tableId, where))
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterListMap(db, context, m, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterListMap(db, context, m, int64(cnt))
	}

	return m, int64(cnt), err
}
func (this *MySqlDataOperator) ListArray(tableId string, filter []string, sort string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error) {
	ret := make([][]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}

	sort = parseSort(sort)
	where := parseFilter(filter)
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeListArray(db, context, &where, &sort, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeListArray(db, context, &where, &sort, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}

	a, err := gosqljson.QueryDbToArray(db, true,
		fmt.Sprint("SELECT * FROM ", tableId, where, sort, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false,
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM ", tableId, where))
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
		cnt, err = strconv.Atoi(c[0]["CNT"])
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterListArray(db, context, a, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterListArray(db, context, a, int64(cnt))
	}

	return a, int64(cnt), err
}
func (this *MySqlDataOperator) Create(tableId string, data map[string]interface{}, context map[string]interface{}) (interface{}, error) {
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeCreate(db, context, data)
		if !ctn {
			return "", err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeCreate(db, context, data)
		if !ctn {
			return "", err
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

	return data["ID"], err
}
func (this *MySqlDataOperator) Update(tableId string, data map[string]interface{}, context map[string]interface{}) (int64, error) {
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeUpdate(db, context, nil, nil)
		if !ctn {
			return 0, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeUpdate(db, context, nil, nil)
		if !ctn {
			return 0, err
		}
	}
	// Update the record
	id := data["ID"]
	if id == nil {
		fmt.Println("ID is not found.")
		return 0, err
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
		return -1, err
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterUpdate(db, context, nil, nil)
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterUpdate(db, context, nil, nil)
	}

	return rowsAffected, err
}
func (this *MySqlDataOperator) Duplicate(tableId string, id string, context map[string]interface{}) (interface{}, error) {
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeDuplicate(db, context, nil, nil)
		if !ctn {
			return "", err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeDuplicate(db, context, nil, nil)
		if !ctn {
			return "", err
		}
	}

	// Duplicate the record
	data, err := gosqljson.QueryDbToMap(db, false,
		fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
	if data == nil || len(data) != 1 {
		return "", err
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

	return newId, err
}
func (this *MySqlDataOperator) Delete(tableId string, id string, context map[string]interface{}) (int64, error) {
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeDelete(db, context, tableId)
		if !ctn {
			return 0, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeDelete(db, context, tableId)
		if !ctn {
			return 0, err
		}
	}
	// Delete the record
	rowsAffected, err := gosqljson.ExecDb(db, fmt.Sprint("DELETE FROM ", tableId, " WHERE ID=?"), id)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterDelete(db, context, tableId)
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterDelete(db, context, tableId)
	}

	return rowsAffected, err
}
func (this *MySqlDataOperator) QueryMap(tableId string, sqlSelect string, sqlSelectCount string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error) {
	ret := make([]map[string]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	if !isSelect(sqlSelect) {
		return ret, -1, nil
	}
	if includeTotal && !isSelect(sqlSelectCount) {
		return ret, -1, nil
	}
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeQueryMap(db, context, &sqlSelect, &sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeQueryMap(db, context, &sqlSelect, &sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}

	m, err := gosqljson.QueryDbToMap(db, true,
		fmt.Sprint(sqlSelect, " LIMIT ?,?"), start, limit)
	cnt := -1
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false, sqlSelectCount)
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
		for _, v := range c[0] {
			cnt, err = strconv.Atoi(v)
		}
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterQueryMap(db, context, m, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterQueryMap(db, context, m, int64(cnt))
	}

	return m, int64(cnt), err
}
func (this *MySqlDataOperator) QueryArray(tableId string, sqlSelect string, sqlSelectCount string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error) {
	ret := make([][]string, 0)
	tableId = normalizeTableId(tableId, this.Ds)
	context["table_id"] = tableId
	if !isSelect(sqlSelect) {
		return ret, -1, errors.New("Invalid query.")
	}
	if includeTotal && !isSelect(sqlSelectCount) {
		return ret, -1, errors.New("Invalid query.")
	}
	db, err := getConn(this.Ds)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeQueryArray(db, context, &sqlSelect, &sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeQueryArray(db, context, &sqlSelect, &sqlSelectCount, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}

	a, err := gosqljson.QueryDbToArray(db, true,
		fmt.Sprint(sqlSelect, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, false, sqlSelectCount)
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
		for _, v := range c[0] {
			cnt, err = strconv.Atoi(v)
		}
		if err != nil {
			fmt.Println(err)
			return ret, -1, err
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterQueryArray(db, context, a, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterQueryArray(db, context, a, int64(cnt))
	}

	return a, int64(cnt), err
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

func parseSort(sort string) string {
	if len(strings.TrimSpace(sort)) == 0 {
		return ""
	}
	return fmt.Sprint(" ORDER BY ", strings.ToUpper(strings.Replace(sort, ":", " ", -1)), " ")
}

func parseFilter(filter []string) string {
	if len(filter) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, v := range filter {

		t := strings.SplitN(v, ",", 3)
		if len(t) <= 1 {
			continue
		} else if len(t) == 2 {
			op := t[1]
			if op == "nu" || op == "nn" {
				f := strings.ToUpper(strings.Replace(strings.Replace(t[0], "'", "", -1), "--", "", -1))
				buffer.WriteString(fmt.Sprint(" AND ", f, ops[op]))
			} else {
				continue
			}
		} else if len(t) == 3 {
			op := t[1]
			f := strings.ToUpper(strings.Replace(strings.Replace(t[0], "'", "", -1), "--", "", -1))
			v := strings.Replace(t[2], "'", "''", -1)
			buffer.WriteString(fmt.Sprint(" AND ", f, ops[op], v))
		}
	}
	return fmt.Sprint(" WHERE 1=1 ", buffer.String())
}

var ops map[string]string = map[string]string{
	"eq": " = ",
	"ne": " != ",
	"gt": " > ",
	"lt": " < ",
	"ge": " >= ",
	"le": " <= ",
	"li": " LIKE ",
	"nl": " NOT LIKE ",
	"nu": " IS NULL ",
	"nn": " IS NOT NULL ",
	"rl": " RLIKE ",
}
