package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"net/http"
)

func Random() int {
	return int(C.random())
}

const (
	a = 7
	b
	c
)

func main() {
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
