package util

import (
	"fmt"
	"reflect"
	"strconv"
)

type T interface {
	ToString() string
	ToInt64() int64
	ToInt32() int32
	ToInt() int
	ToUint64() uint64
	ToUint32() uint32
	ToFloat64() float64
	ToFloat32() float32
	ToBool() bool
	ToMap() map[string]interface{}
	ToSlice() []interface{}
}

func NewT(v interface{}) T {
	return &t{v}
}

type t struct {
	v interface{}
}

func (t *t) ToString() string {
	if t.v == nil {
		return ""
	}
	return fmt.Sprintf("%v", t.v)
}

func (t *t) ToInt64() int64 {
	v, err := strconv.ParseInt(t.ToString(), 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func (t *t) ToInt32() int32 {
	return int32(t.ToInt64())
}

func (t *t) ToInt() int {
	return int(t.ToInt64())
}

func (t *t) ToUint64() uint64 {
	v, err := strconv.ParseUint(t.ToString(), 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func (t *t) ToUint32() uint32 {
	return uint32(t.ToUint64())
}

func (t *t) ToFloat64() float64 {
	v, err := strconv.ParseFloat(t.ToString(), 64)
	if err != nil {
		return 0
	}
	return v
}

func (t *t) ToFloat32() float32 {
	return float32(t.ToFloat64())
}

func (t *t) ToBool() bool {
	v, err := strconv.ParseBool(t.ToString())
	if err != nil {
		return false
	}
	return v
}

func (t *t) ToMap() map[string]interface{} {
	if t.v == nil {
		return make(map[string]interface{})
	}
	v := reflect.ValueOf(t.v)
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Map {
		return make(map[string]interface{})
	}
	mp := make(map[string]interface{})
	for _, m := range v.MapKeys() {
		mp[m.String()] = v.MapIndex(m).Interface()
	}
	return mp
}

func (t *t) ToSlice() []interface{} {
	v := reflect.ValueOf(t.v)
	if v.IsNil() {
		return nil
	}
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Slice {
		return nil
	}
	s := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		s[i] = v.Index(i).Interface()
	}
	return s
}
