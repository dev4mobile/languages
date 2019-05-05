package main

import (
	"fmt"
	"reflect"
)

type MyMath struct {
	Pi float64
}

func (myMath MyMath) Sum(a, b int) int {
	return a + b
}

func (MyMath MyMath) Desc(a, b int) int {
	return a - b
}

func main() {
	m := MyMath{3.14}
	rValue := reflect.ValueOf(m)
	fmt.Println(rValue.NumMethod())
	method := rValue.Method(0)
	paramList := []reflect.Value{reflect.ValueOf(30), reflect.ValueOf(20)}
	r := method.Call(paramList)
	fmt.Println(r[0].Int())
}
