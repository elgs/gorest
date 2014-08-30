package gorest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elgs/gosqljson"
	"io/ioutil"
)

func init() {
	loadACL()
	RegisterGlobalDataInterceptor(&GlobalTokenInterceptor{Id: "GlobalTokenInterceptor"})
}

var acl = make(map[string]map[string]bool)
var tokenRegistry = make(map[string][]string) // id:[token_key, user_id]

func checkToken(db *sql.DB, id string, key string, context map[string]interface{}) (bool, error) {
	if id != "" && key != "" && len(tokenRegistry[id]) > 0 && tokenRegistry[id][0] == key {
		return true, nil
	}
	tokenTable := context["token_table"]
	if tokenTable != nil {
		tokenTableString := tokenTable.(string)
		if len(tokenTableString) > 0 && context["table_id"] == tokenTable {
			return false, errors.New("We think you are invading the system.")
		}

		mysqlSafe(&tokenTableString)
		data, err := gosqljson.QueryDbToMap(db, false, fmt.Sprint("SELECT * FROM ", tokenTableString, " WHERE ID=? AND TOKEN_KEY=? AND STATUS=?"), id, key, "0")
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		if data != nil && len(data) == 1 {
			record := data[0]
			tokenRegistry[record["ID"]] = []string{record["TOKEN_KEY"], record["USER_ID"]}
			return true, nil
		}
	}
	return false, errors.New("Authentication failed.")
}

func loadACL() {
	// load acl from configuration files.
	configFile := "gorest_acl.json"
	aclConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(configFile, " not found, default policies are used.")
	}
	err = json.Unmarshal(aclConfig, &acl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(len(acl), acl)
}

func checkACL(context map[string]interface{}, op string) (bool, error) {
	tableId := context["table_id"].(string)
	if acl[tableId] != nil {
		if acl[tableId][op] {
			return true, nil
		}
	}
	return false, errors.New("Access denied.")
}

type GlobalTokenInterceptor struct {
	*DefaultDataInterceptor
	Id string
}

func (this *GlobalTokenInterceptor) BeforeCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) (bool, error) {
	if ok, err := checkACL(context, "create"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterCreate(ds interface{}, context map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeLoad(ds interface{}, context map[string]interface{}, id string) (bool, error) {
	if ok, err := checkACL(context, "load"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterLoad(ds interface{}, context map[string]interface{}, data map[string]string) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error) {
	if ok, err := checkACL(context, "update"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterUpdate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) (bool, error) {
	if ok, err := checkACL(context, "duplicate"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterDuplicate(ds interface{}, context map[string]interface{}, oldData map[string]interface{}, data map[string]interface{}) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeDelete(ds interface{}, context map[string]interface{}, id string) (bool, error) {
	if ok, err := checkACL(context, "delete"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterDelete(ds interface{}, context map[string]interface{}, id string) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeListMap(ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	if ok, err := checkACL(context, "list"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterListMap(ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeListArray(ds interface{}, context map[string]interface{}, filter *string, sort *string, start int64, limit int64, includeTotal bool) (bool, error) {
	if ok, err := checkACL(context, "list"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterListArray(ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeQueryMap(ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	if ok, err := checkACL(context, "query"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterQueryMap(ds interface{}, context map[string]interface{}, data []map[string]string, total int64) error {
	return nil
}
func (this *GlobalTokenInterceptor) BeforeQueryArray(ds interface{}, context map[string]interface{}, sqlSelect *string, sqlSelectCount *string, start int64, limit int64, includeTotal bool) (bool, error) {
	if ok, err := checkACL(context, "query"); !ok {
		return false, err
	}
	db := ds.(*sql.DB)
	return checkToken(db, context["api_token_id"].(string), context["api_token_key"].(string), context)
}
func (this *GlobalTokenInterceptor) AfterQueryArray(ds interface{}, context map[string]interface{}, data [][]string, total int64) error {
	return nil
}
