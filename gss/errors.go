package gss

import "errors"

var (
	// 更新错误版本号错误
	// 1. insert的时候, 唯一索引导致插入失败
	// 2. 更新的时候, version不对导致更新错误
	ErrVersionNotMatch = errors.New("version not match")

	ErrInnerError      = errors.New("storage inner error")
	ErrRecordNotExists = errors.New("record not exists")
)
