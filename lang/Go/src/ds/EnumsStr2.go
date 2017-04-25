////////////////////////////////////////////////////////////////////////////
// Porgram: EnumsStr2.go
// Purpose: Go Enum and its string demo, using reflection
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Egon https://groups.google.com/d/msg/golang-nuts/fCdBSRNNUY8/P45qC_03LoAJ
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"reflect"
)

var strs = make(map[reflect.Type]map[int]string)

func Register(t reflect.Type, v int, s string) {
	m, ok := strs[t]
	if !ok {
		m = make(map[int]string)
		strs[t] = m
	}
	m[v] = s
}

func GetStr(e interface{}) string {
	t := reflect.TypeOf(e)
	v := int(reflect.ValueOf(e).Int())
	return strs[t][v]
}

func InitEnums(estr interface{}) {
	v := reflect.ValueOf(estr).Elem()
	vt := v.Type()
	for i, n := 0, v.NumField(); i < n; i += 1 {
		f := v.Field(i)
		Register(f.Type(), i, string(vt.Field(i).Tag))
		f.SetInt(int64(i))
	}
}

type Enum int

func (e Enum) String() string { return GetStr(e) }

var Enums struct {
	Alpha Enum "Alpha"
	Beta  Enum "Beta"
}

func main() {
	InitEnums(&Enums)

	fmt.Printf("%+v\n", Enums)

	fmt.Printf("%v\n", Enums.Alpha.String())
	fmt.Printf("%v\n", Enums.Beta)
}

/*

{Alpha:Alpha Beta:Beta}
Alpha
Beta

*/
