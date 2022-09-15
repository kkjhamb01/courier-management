package tag

import (
	"time"
)

const (
	TypeString = iota + 1
	TypeInt64
	TypeInt32
	TypeInt
	TypeFloat64
	TypeFloat32
	TypeBoolean
	TypeObject
	TypeError
	TypeTime
	TypeDuration
)

type Tag struct {
	tagType int
	Key     string
	Val     interface{}
}

func (t Tag) Type() int {
	return t.tagType
}

func Str(key string, value string) Tag {
	return Tag{
		tagType: TypeString,
		Key:     key,
		Val:     value,
	}
}

func Int64(key string, value int64) Tag {
	return Tag{
		tagType: TypeInt64,
		Key:     key,
		Val:     value,
	}
}

func Int(key string, value int) Tag {
	return Tag{
		tagType: TypeInt,
		Key:     key,
		Val:     value,
	}
}

func Int32(key string, value int32) Tag {
	return Tag{
		tagType: TypeInt32,
		Key:     key,
		Val:     value,
	}
}

func Float64(key string, value float64) Tag {
	return Tag{
		tagType: TypeFloat64,
		Key:     key,
		Val:     value,
	}
}

func Float32(key string, value float32) Tag {
	return Tag{
		tagType: TypeFloat32,
		Key:     key,
		Val:     value,
	}
}

func Bool(key string, value bool) Tag {
	return Tag{
		tagType: TypeBoolean,
		Key:     key,
		Val:     value,
	}
}

func Err(key string, value error) Tag {
	return Tag{
		tagType: TypeError,
		Key:     key,
		Val:     value,
	}
}

func Duration(key string, value time.Duration) Tag {
	return Tag{
		tagType: TypeDuration,
		Key:     key,
		Val:     value,
	}
}

func Time(key string, value time.Time) Tag {
	return Tag{
		tagType: TypeTime,
		Key:     key,
		Val:     value,
	}
}

func Obj(key string, value interface{}) Tag {
	return Tag{
		tagType: TypeObject,
		Key:     key,
		Val:     value,
	}
}
