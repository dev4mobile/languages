package main

import (
	"fmt"
	"playground/data_type"
	"playground/format"
	"playground/object"
	"playground/ops_channel"
	"playground/read_mutex"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}

func TestWriteFile(t *testing.T) {
	//file_usage.WriteFile()

}

func TestReadFile(t *testing.T)  {
	//file_usage.ReadFile()
}


func TestObject(t *testing.T) {
	person := object.Person{12}
	person.SetAge(23)
	fmt.Printf("age:%d\n", person.Age)
}

func TestDataType(t *testing.T) {
	data_type.DataType()
}

func TestMutexRead(t *testing.T) {
	read_mutex.ReadMutex()
}

func TestPrintf(t *testing.T)  {
	format.Printf()
}

func TestChannel(t *testing.T)  {
	go ops_channel.SamplChannel()
	time.Sleep(4 * time.Second)
	fmt.Printf("hello")
}

func TestSplit(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		fmt.Println("no equal")
	}

	var a, b []string = nil, []string{}
	var c, d map[string]int = nil, make(map[string]int)
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(c, d))
	fmt.Println(reflect.ValueOf(nil).Kind())

}