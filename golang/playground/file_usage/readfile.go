package main

import (
	"fmt"
	"os"
)

func main() {
	file, e := os.Open("textfile.txt")
	if e != nil {
		panic(e)
	}

	var (
		L int32
		M float32
		N bool
		O string
	)

	fmt.Fscanf(file, "%d %f %t %s", &L, &M, &N, &O)
	fmt.Println(L, M, N, O)
	fmt.Fscanf(file, "%d %f %t %s", &L, &M, &N, &O)
	fmt.Println(L, M, N, O)
}
