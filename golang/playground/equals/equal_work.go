package main

import (
	"fmt"
	"reflect"
)

func equal(x, y reflect.Value) bool  {
	fmt.Println(x.Type())
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		fmt.Println("type not equal", x.Type())
		return false
	}
	switch x.Kind(){
	case reflect.Bool:
		fmt.Println(x.Bool(), x.Type())
		return x.Bool() == y.Bool()
	case reflect.Int:
		fmt.Println("Int", x.Int())
		return x.Int() == y.Int()
	case reflect.String:
		fmt.Println(x.String(), x.Type())
		return x.String() == y.String()
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		fmt.Println(x.Pointer(), x.Type())
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		fmt.Println("ptr", x.Elem())
		return equal(x.Elem(), y.Elem())
	case reflect.Struct:
		fmt.Println("struct", x.Interface())
		return false
	}
	return false
}


type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func(s S) Get()int {
	return s.Age
}

func(s *S) Set(age int) {
	s.Age = age
}
func main() {
	//i :=1
	//equal(reflect.ValueOf(S{}), reflect.ValueOf(S{}))
	var v interface{} = 10
	fmt.Println(reflect.ValueOf(&v), reflect.ValueOf(&v).Type())
	var t = reflect.ValueOf(&v).Type().Elem()
	fmt.Println(t.Kind())
	fmt.Println(t.Kind() == reflect.Interface)
}
