package tag

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStr(t *testing.T) {
	key := "k"
	val := "v"
	tag := Str(key, val)
	assert.Equal(t, TypeString, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestInt64(t *testing.T) {
	key := "k"
	val := int64(12345)
	tag := Int64(key, val)
	assert.Equal(t, TypeInt64, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestInt(t *testing.T) {
	key := "k"
	val := 12345
	tag := Int(key, val)
	assert.Equal(t, TypeInt, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestInt32(t *testing.T) {
	key := "k"
	val := int32(12345)
	tag := Int32(key, val)
	assert.Equal(t, TypeInt32, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestFloat64(t *testing.T) {
	key := "k"
	val := float64(12345)
	tag := Float64(key, val)
	assert.Equal(t, TypeFloat64, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestFloat32(t *testing.T) {
	key := "k"
	val := float32(12345)
	tag := Float32(key, val)
	assert.Equal(t, TypeFloat32, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestBool(t *testing.T) {
	key := "k"
	val := false
	tag := Bool(key, val)
	assert.Equal(t, TypeBoolean, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestErr(t *testing.T) {
	key := "k"
	val := errors.New("my custom error to test")
	tag := Err(key, val)
	assert.Equal(t, TypeError, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestDuration(t *testing.T) {
	key := "k"
	val := time.Duration(40000)
	tag := Duration(key, val)
	assert.Equal(t, TypeDuration, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestTime(t *testing.T) {
	key := "k"
	val := time.Now()
	tag := Time(key, val)
	assert.Equal(t, TypeTime, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}

func TestObj(t *testing.T) {
	key := "k"
	val := time.Now()
	tag := Obj(key, val)
	assert.Equal(t, TypeObject, tag.tagType)
	assert.Equal(t, key, tag.Key)
	assert.Equal(t, val, tag.Val)
}
