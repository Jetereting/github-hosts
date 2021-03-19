package util

import (
	"strconv"
	"strings"
)

type Str string

// Convert Str to string
func (s Str) ToString() string {
	return string(s)
}

// Convert string to int64
func (s Str) ToInt64() int64 {
	i, err := strconv.ParseInt(s.ToString(), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// Convert string to int32
func (s Str) ToInt32() int32 {
	return int32(s.ToInt64())
}

// Convert string to int
func (s Str) ToInt() int {
	return int(s.ToInt32())
}

// Convert string to float64
func (s Str) ToFloat64() float64 {
	f, err := strconv.ParseFloat(s.ToString(), 64)
	if err != nil {
		return 0
	}
	return f
}

// Convert string to float32
func (s Str) ToFloat32() float32 {
	return float32(s.ToFloat64())
}

// ToMap Convert string to Map
func (s Str) ToMap() map[string]string {
	if s == "" {
		return make(map[string]string)
	}
	v := strings.Split(s.ToString()[2:len(s.ToString())-1], ",")
	mp := make(map[string]string)
	for _, m := range v {
		t := strings.Split(m, "\":\"")
		mp[strings.Trim(t[0], "\"")] = strings.Trim(t[1], "\"")
	}
	return mp
}
// GetBetweenStr 获取中间字符
func (str Str)GetBetween(start, end string) string {
	str0:=str.ToString()
	n := strings.Index(str0, start)
	if n == -1 {
		return ""
	}
	str0 = string([]byte(str0)[n+len(start):])
	m := strings.Index(str0, end)
	if m == -1 {
		m = len(str0)
	}
	str0 = string([]byte(str0)[:m])
	return str0
}