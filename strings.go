package gutil

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ToStr(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func ToJson(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}

func ParseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func ParseInt32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}

func NewInt(i int) *int {
	i1 := i
	return &i1
}

func NewInt32(i int32) *int32 {
	i1 := i
	return &i1
}

func NewInt64(i int64) *int64 {
	i1 := i
	return &i1
}

func NewString(s string) *string {
	s1 := s
	return &s1
}

func MysqlDSN(dbHost string, dbPort uint16, dbUser, dbPass, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}
