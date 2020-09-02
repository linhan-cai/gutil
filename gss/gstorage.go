package gss

type KVStorage interface {
	Get(key string) (value []byte, version int32, err error)
	Set(key string, value []byte, version int32) error
	Del(key string, version int32) error
}
