package main

import (
	"fmt"

	"github.com/aaronclong/capybaradb/src/query/processor"
)

func main() {
	fmt.Println("Hello, Gopher!")
	processor.Lexer("SELECT * FROM 'DB'")
}
