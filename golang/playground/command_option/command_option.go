package command_option

import (
	"flag"
	"fmt"
)

func FlagTest() {
	wordPtr := flag.String("word","foo", "a string")
	var svar string
	flag.StringVar(&svar, "svar","bar", "a string var")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("svar:", svar)
}