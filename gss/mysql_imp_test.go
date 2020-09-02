package gss

import (
	"github.com/linhan-cai/gutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	imp, err := NewKVStorageMysqlImp(gutil.MysqlConf{})
	assert.NoError(t, err)

	value, version, err := imp.Get("abc")
	assert.NoError(t, err)
	t.Log(string(value), version)

	err = imp.Set("abc", append(value, 'a'), version)
	assert.NoError(t, err)

	value, version, err = imp.Get("abc")
	assert.NoError(t, err)
	t.Log(string(value), version)

	//err = imp.Del("abc", version)
	//assert.NoError(t, err)
}
