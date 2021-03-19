package util

import (
	"time"
	"math/rand"
)

type StrArr []string

func (sa StrArr)Random2String() (s string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(sa)) {
		val := sa[i]
		s+=string(val)+","
	}
	if len(s)>0 {
		return s[:len(s) - 1]
	}
	return
}
func (sa StrArr)Random2Arr() (s []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(sa)) {
		val := sa[i]
		s=append(s,val)
	}
	return
}
