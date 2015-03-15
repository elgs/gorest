package gorest

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	//"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/elgs/exparser"
	"github.com/elgs/gosqljson"
)

type MySqlDataOperator struct {
	*DefaultDataOperator
	Ds         string
	DbType     string
	TokenTable string
	Db         *sql.DB
}

func (this *MySqlDataOperator) Load(tableId string, id string, field []string, context map[string]interface{}) (map[string]string, error) {
	ret := make(map[string]string, 0)
	tableId = normalizeTableId(tableId, this.DbType, this.Ds)
	context["token_table"] = this.TokenTable
	db, err := this.GetConn()

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeLoad(tableId, db, field, context, id)
		if !ctn {
			return ret, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeLoad(tableId, db, field, context, id)
		if !ctn {
			return ret, err
		}
	}

	// Load the record
	extraFilter := context["extra_filter"]
	if extraFilter == nil {
		extraFilter = ""
	}
	c := context["case"].(string)

	fields := parseFields(field)
	m, err := gosqljson.QueryDbToMap(db, c,
		fmt.Sprint("SELECT", fields, "FROM ", tableId, " WHERE ID=? ", extraFilter), id)
	if err != nil {
		fmt.Println(err)
		return ret, err
	}

	if len(m) == 0 {
		m = []map[string]string{
			make(map[string]string, 0),
		}
	}

	if dataInterceptor != nil {
		dataInterceptor.AfterLoad(tableId, db, field, context, m[0])
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterLoad(tableId, db, field, context, m[0])
	}

	if m != nil && len(m) == 1 {
		return m[0], err
	} else {
		return ret, err
	}

}
func (this *MySqlDataOperator) ListMap(tableId string, field []string, filter []string, sort string, group string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([]map[string]string, int64, error) {
	ret := make([]map[string]string, 0)
	tableId = normalizeTableId(tableId, this.DbType, this.Ds)
	context["token_table"] = this.TokenTable
	db, err := this.GetConn()

	sort = parseSort(sort)
	where := parseFilters(filter)
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeListMap(tableId, db, field, context, &where, &sort, &group, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeListMap(tableId, db, field, context, &where, &sort, &group, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}

	c := context["case"].(string)
	fields := parseFields(field)
	m, err := gosqljson.QueryDbToMap(db, c,
		fmt.Sprint("SELECT", fields, "FROM ", tableId, where, parseGroup(group), sort, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, "upper",
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM (", "SELECT", fields, "FROM ", tableId, where, parseGroup(group), ")a"))
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
		dataInterceptor.AfterListMap(tableId, db, field, context, m, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterListMap(tableId, db, field, context, m, int64(cnt))
	}

	return m, int64(cnt), err
}
func (this *MySqlDataOperator) ListArray(tableId string, field []string, filter []string, sort string, group string,
	start int64, limit int64, includeTotal bool, context map[string]interface{}) ([][]string, int64, error) {
	ret := make([][]string, 0)
	tableId = normalizeTableId(tableId, this.DbType, this.Ds)
	context["token_table"] = this.TokenTable
	db, err := this.GetConn()

	sort = parseSort(sort)
	where := parseFilters(filter)
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeListArray(tableId, db, field, context, &where, &sort, &group, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeListArray(tableId, db, field, context, &where, &sort, &group, start, limit, includeTotal)
		if !ctn {
			return ret, -1, err
		}
	}

	c := context["case"].(string)
	fields := parseFields(field)
	a, err := gosqljson.QueryDbToArray(db, c,
		fmt.Sprint("SELECT", fields, "FROM ", tableId, where, parseGroup(group), sort, " LIMIT ?,?"), start, limit)
	if err != nil {
		fmt.Println(err)
		return ret, -1, err
	}
	cnt := -1
	if includeTotal {
		c, err := gosqljson.QueryDbToMap(db, "upper",
			fmt.Sprint("SELECT COUNT(*) AS CNT FROM (", "SELECT", fields, "FROM ", tableId, where, parseGroup(group), ")a"))
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
		dataInterceptor.AfterListArray(tableId, db, field, context, a, int64(cnt))
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		globalDataInterceptor.AfterListArray(tableId, db, field, context, a, int64(cnt))
	}

	return a, int64(cnt), err
}
func (this *MySqlDataOperator) Create(tableId string, data map[string]interface{}, context map[string]interface{}) (interface{}, map[string]interface{}, error) {
	tableId = normalizeTableId(tableId, this.DbType, this.Ds)
	context["token_table"] = this.TokenTable
	info := make(map[string]interface{})
	db, err := this.GetConn()

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeCreate(tableId, db, context, info, data)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeCreate(tableId, db, context, info, data)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}

	// Create the record
	if data["ID"] == nil {
		data["ID"] = uuid.New()
	}
	if data["SEQ"] == nil {
		//data["SEQ"] = time.Now().Unix()
	}
	dataLen := len(data)
	values := make([]interface{}, 0, dataLen)
	var fieldBuffer bytes.Buffer
	var qmBuffer bytes.Buffer
	count := 0
	for k, v := range data {
		count++
		if count == dataLen {
			fieldBuffer.WriteString(k)
			qmBuffer.WriteString("?")
		} else {
			fieldBuffer.WriteString(fmt.Sprint(k, ","))
			qmBuffer.WriteString("?,")
		}
		values = append(values, v)
	}
	fields := fieldBuffer.String()
	qms := qmBuffer.String()
	if tx, ok := context["tx"].(*sql.Tx); ok {
		_, err = gosqljson.ExecTx(tx, fmt.Sprint("INSERT INTO ", tableId, " (", fields, ") VALUES (", qms, ")"), values...)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return nil, info, err
		}
	} else {
		_, err = gosqljson.ExecDb(db, fmt.Sprint("INSERT INTO ", tableId, " (", fields, ") VALUES (", qms, ")"), values...)
		if err != nil {
			fmt.Println(err)
			return nil, info, err
		}
	}

	if dataInterceptor != nil {
		err := dataInterceptor.AfterCreate(tableId, db, context, info, data)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		err := globalDataInterceptor.AfterCreate(tableId, db, context, info, data)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}

	if tx, ok := context["tx"].(*sql.Tx); ok {
		tx.Commit()
	}

	return data["ID"], info, err
}
func (this *MySqlDataOperator) Update(tableId string, data map[string]interface{}, context map[string]interface{}) (int64, map[string]interface{}, error) {
	tableId = normalizeTableId(tableId, this.DbType, this.Ds)
	context["token_table"] = this.TokenTable
	info := make(map[string]interface{})
	db, err := this.GetConn()

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeUpdate(tableId, db, context, info, data)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return 0, info, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeUpdate(tableId, db, context, info, data)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return 0, info, err
		}
	}
	// Update the record
	id := data["ID"]
	if id == nil {
		fmt.Println("ID is not found.")
		if tx, ok := context["tx"].(*sql.Tx); ok {
			tx.Rollback()
		}
		return 0, info, err
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
	var rowsAffected int64 = 0
	if tx, ok := context["tx"].(*sql.Tx); ok {
		rowsAffected, err = gosqljson.ExecTx(tx, fmt.Sprint("UPDATE ", tableId, " SET ", sets, " WHERE ID=?"), values...)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return -1, info, err
		}
	} else {
		rowsAffected, err = gosqljson.ExecDb(db, fmt.Sprint("UPDATE ", tableId, " SET ", sets, " WHERE ID=?"), values...)
		if err != nil {
			fmt.Println(err)
			return -1, info, err
		}
	}

	if dataInterceptor != nil {
		err := dataInterceptor.AfterUpdate(tableId, db, context, info, data)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return -1, info, err
		}
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		err := globalDataInterceptor.AfterUpdate(tableId, db, context, info, data)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return -1, info, err
		}
	}

	if tx, ok := context["tx"].(*sql.Tx); ok {
		tx.Commit()
	}

	return rowsAffected, info, err
}
func (this *MySqlDataOperator) Duplicate(tableId string, id string, context map[string]interface{}) (interface{}, map[string]interface{}, error) {
	tableId = normalizeTableId(tableId, this.DbType, this.Ds)
	context["token_table"] = this.TokenTable
	info := make(map[string]interface{})
	db, err := this.GetConn()

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeDuplicate(tableId, db, context, info, id)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeDuplicate(tableId, db, context, info, id)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}

	newId := uuid.New()
	// Duplicate the record
	if tx, ok := context["tx"].(*sql.Tx); ok {
		data, err := gosqljson.QueryTxToMap(tx, "upper",
			fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
		if data == nil || len(data) != 1 {
			tx.Rollback()
			return nil, info, err
		}
		newData := make(map[string]interface{}, len(data[0]))
		for k, v := range data[0] {
			newData[k] = v
		}
		newData["ID"] = newId

		newDataLen := len(newData)
		newValues := make([]interface{}, 0, newDataLen)
		var fieldBuffer bytes.Buffer
		var qmBuffer bytes.Buffer
		count := 0
		for k, v := range newData {
			count++
			if count == newDataLen {
				fieldBuffer.WriteString(k)
				qmBuffer.WriteString("?")
			} else {
				fieldBuffer.WriteString(fmt.Sprint(k, ","))
				qmBuffer.WriteString("?,")
			}
			newValues = append(newValues, v)
		}
		fields := fieldBuffer.String()
		qms := qmBuffer.String()
		_, err = gosqljson.ExecTx(tx, fmt.Sprint("INSERT INTO ", tableId, " (", fields, ") VALUES (", qms, ")"), newValues...)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return nil, info, err
		}
	} else {
		data, err := gosqljson.QueryDbToMap(db, "upper",
			fmt.Sprint("SELECT * FROM ", tableId, " WHERE ID=?"), id)
		if data == nil || len(data) != 1 {
			return nil, info, err
		}
		newData := make(map[string]interface{}, len(data[0]))
		for k, v := range data[0] {
			newData[k] = v
		}
		newData["ID"] = newId

		newDataLen := len(newData)
		newValues := make([]interface{}, 0, newDataLen)
		var fieldBuffer bytes.Buffer
		var qmBuffer bytes.Buffer
		count := 0
		for k, v := range newData {
			count++
			if count == newDataLen {
				fieldBuffer.WriteString(k)
				qmBuffer.WriteString("?")
			} else {
				fieldBuffer.WriteString(fmt.Sprint(k, ","))
				qmBuffer.WriteString("?,")
			}
			newValues = append(newValues, v)
		}
		fields := fieldBuffer.String()
		qms := qmBuffer.String()
		_, err = gosqljson.ExecDb(db, fmt.Sprint("INSERT INTO ", tableId, " (", fields, ") VALUES (", qms, ")"), newValues...)
		if err != nil {
			fmt.Println(err)
			return nil, info, err
		}
	}
	if dataInterceptor != nil {
		err := dataInterceptor.AfterDuplicate(tableId, db, context, info, id, newId)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		err := globalDataInterceptor.AfterDuplicate(tableId, db, context, info, id, newId)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return nil, info, err
		}
	}

	if tx, ok := context["tx"].(*sql.Tx); ok {
		tx.Commit()
	}

	return newId, info, err
}
func (this *MySqlDataOperator) Delete(tableId string, id string, context map[string]interface{}) (int64, map[string]interface{}, error) {
	tableId = normalizeTableId(tableId, this.DbType, this.Ds)
	context["token_table"] = this.TokenTable
	info := make(map[string]interface{})
	db, err := this.GetConn()

	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		ctn, err := globalDataInterceptor.BeforeDelete(tableId, db, context, info, id)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return 0, info, err
		}
	}
	dataInterceptor := GetDataInterceptor(tableId)
	if dataInterceptor != nil {
		ctn, err := dataInterceptor.BeforeDelete(tableId, db, context, info, id)
		if !ctn {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return 0, info, err
		}
	}

	var rowsAffected int64 = 0
	if tx, ok := context["tx"].(*sql.Tx); ok {
		load := context["load"].(bool)
		if load {
			data, err := gosqljson.QueryTxToMap(tx, "upper", "SELECT * FROM "+tableId+" WHERE ID=?", id)
			if err != nil {
				fmt.Println(err)
				tx.Rollback()
				return -1, info, err
			}
			if data == nil && len(data) != 1 {
				tx.Rollback()
				return -1, info, errors.New(id + " not found.")
			} else {
				context["data"] = data[0]
			}
		}

		// Delete the record
		rowsAffected, err = gosqljson.ExecTx(tx, fmt.Sprint("DELETE FROM ", tableId, " WHERE ID=?"), id)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return -1, info, err
		}
	} else {
		load := context["load"].(bool)
		if load {
			data, err := gosqljson.QueryDbToMap(db, "upper", "SELECT * FROM "+tableId+" WHERE ID=?", id)
			if err != nil {
				fmt.Println(err)
				return -1, info, err
			}
			if data == nil && len(data) != 1 {
				return -1, info, errors.New(id + " not found.")
			} else {
				context["data"] = data[0]
			}
		}

		// Delete the record
		rowsAffected, err = gosqljson.ExecDb(db, fmt.Sprint("DELETE FROM ", tableId, " WHERE ID=?"), id)
		if err != nil {
			fmt.Println(err)
			return -1, info, err
		}
	}

	if dataInterceptor != nil {
		err := dataInterceptor.AfterDelete(tableId, db, context, info, id)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return -1, info, err
		}
	}
	for _, globalDataInterceptor := range GlobalDataInterceptorRegistry {
		err := globalDataInterceptor.AfterDelete(tableId, db, context, info, id)
		if err != nil {
			if tx, ok := context["tx"].(*sql.Tx); ok {
				tx.Rollback()
			}
			return -1, info, err
		}
	}
	if tx, ok := context["tx"].(*sql.Tx); ok {
		tx.Commit()
	}
	return rowsAffected, info, err
}

//func isSelect(sqlSelect string) bool {
//	return strings.HasPrefix(strings.ToUpper(sqlSelect), "SELECT ")
//}

func (this *MySqlDataOperator) MetaData(resourceId string) ([]map[string]string, error) {
	db, err := this.GetConn()
	if err != nil {
		return nil, err
	}
	m, err := gosqljson.QueryDbToMap(db, "upper", "DESCRIBE "+resourceId)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (this *MySqlDataOperator) GetConn() (*sql.DB, error) {
	if this.Db == nil {
		if len(strings.TrimSpace(this.DbType)) == 0 {
			this.DbType = "mysql"
		}
		db, err := sql.Open(this.DbType, this.Ds)
		//fmt.Println("New db conn created.")
		if err != nil {
			return nil, err
		}
		this.Db = db
	}
	return this.Db, nil
}

func extractDbNameFromDs(dbType string, ds string) string {
	switch dbType {
	case "sqlite3":
		return ""
	default:
		a := strings.LastIndex(ds, "/")
		b := ds[a+1:]
		c := strings.Index(b, "?")
		if c < 0 {
			return b
		}
		return b[:c]
	}
}

func normalizeTableId(tableId string, dbType string, ds string) string {
	if strings.Contains(tableId, ".") {
		a := strings.Split(tableId, ".")
		return fmt.Sprint(a[0], ".", a[1])
	}
	db := extractDbNameFromDs(dbType, ds)

	MysqlSafe(&tableId)
	if len(strings.TrimSpace(db)) == 0 {
		return tableId
	} else {
		MysqlSafe(&db)
		return fmt.Sprint(db, ".", tableId)
	}
}

func MysqlSafe(s *string) {
	*s = strings.Replace(*s, "'", "''", -1)
	*s = strings.Replace(*s, "--", "", -1)
}

func parseSort(sort string) string {
	if len(strings.TrimSpace(sort)) == 0 {
		return ""
	}
	return fmt.Sprint(" ORDER BY ", strings.ToUpper(strings.Replace(sort, ":", " ", -1)), " ")
}

func parseFilter(filter string) string {
	r, err := parser.Calculate(filter)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return r
}

func parseFilters(filters []string) (r string) {
	for _, v := range filters {
		r += fmt.Sprint("AND ", parseFilter(v))
	}
	r = fmt.Sprint(" WHERE 1=1 ", r, " ")
	//fmt.Println(r)
	return
}

var parser = &exparser.Parser{
	Operators: exparser.MysqlOperators,
}

func parseFields(fields []string) (r string) {
	if fields == nil || len(fields) == 0 {
		return " * "
	}
	for i, v := range fields {
		if i == len(fields)-1 {
			r += fmt.Sprint(v)
		} else {
			r += fmt.Sprint(v, ",")
		}
	}
	r = fmt.Sprint(" ", r, " ")
	return
}
func parseGroup(group string) (r string) {
	if strings.TrimSpace(group) == "" {
		return ""
	}
	r = fmt.Sprint(" GROUP BY ", strings.ToUpper(group))
	return
}
