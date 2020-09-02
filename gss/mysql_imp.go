package gss

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/linhan-cai/gutil"

	_ "github.com/go-sql-driver/mysql"
)

type KVStorageMysqlImp struct {
	dbHandle *sql.DB
}

func NewKVStorageMysqlImp(conf gutil.MysqlConf) (*KVStorageMysqlImp, error) {
	imp := new(KVStorageMysqlImp)
	db, err := sql.Open("mysql", gutil.MysqlDSN(conf.Host, conf.Port, conf.User, conf.Pass, conf.Name))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	imp.dbHandle = db
	return imp, nil
}

func (imp KVStorageMysqlImp) Get(key string) (value []byte, version int32, err error) {

	value = make([]byte, 0)

	stat := fmt.Sprintf("SELECT v, version FROM %v where k = ?", imp.tbSharp(key))
	rows, err := imp.dbHandle.Query(stat, key)
	if err != nil {
		return value, 0, ErrInnerError
	}

	for rows.Next() {
		if rows.Err() != nil {
			return value, 0, ErrInnerError
		}

		err = rows.Scan(&value, &version)
		if err != nil {
			return value, 0, ErrInnerError
		}

		break
	}

	return value, version, nil
}

func (imp KVStorageMysqlImp) Set(key string, value []byte, version int32) error {
	var stat string
	var args []interface{}
	if version == 0 {
		stat = fmt.Sprintf("insert into %v (k, v, version) values (?, ?, ?)", imp.tbSharp(key))
		args = []interface{}{key, value, version + 1}
	} else {
		stat = fmt.Sprintf("update %v set v=?, version = version + 1 where k = ? and version = ?", imp.tbSharp(key))
		args = []interface{}{value, key, version}
	}

	result, err := imp.dbHandle.Exec(stat, args...)
	if err != nil {
		if MySQLErrorParser(err).Code == ErrMySQLDuplicateKey {
			return ErrVersionNotMatch
		}
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return ErrInnerError
	}

	if affected != 1 {
		return ErrVersionNotMatch
	}

	return nil
}

func (imp KVStorageMysqlImp) Del(key string, version int32) error {
	var stat string
	var args []interface{}
	if version < 1 {
		stat = fmt.Sprintf("delete from %v where k = ?", imp.tbSharp(key))
		args = []interface{}{key}
	} else {
		stat = fmt.Sprintf("delete from %v where k = ? and version = ?", imp.tbSharp(key))
		args = []interface{}{key, version}
	}

	result, err := imp.dbHandle.Exec(stat, args...)
	if err != nil {
		return ErrInnerError
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return ErrInnerError
	}

	if affected != 1 {
		return ErrVersionNotMatch
	}

	return nil
}

func (imp KVStorageMysqlImp) dbSharp(key string) string {
	return "kv_storage"
}

func (imp KVStorageMysqlImp) tbSharp(key string) string {
	return "kv_1"
}

type MySQLError struct {
	Code    int32
	Message string
}

func MySQLErrorParser(err error) MySQLError {
	parser := MySQLError{}
	errMsg := err.Error()
	errParts := strings.Split(errMsg, ":")
	if len(errParts) == 2 && strings.HasPrefix(errParts[0], "Error") {
		parser.Code = gutil.ParseInt32(errParts[0][6:])
		parser.Message = errParts[1]
		return parser
	}

	parser.Code = -1
	parser.Message = errMsg
	return parser
}

const (
	ErrMySQLDuplicateKey = 1062
)
