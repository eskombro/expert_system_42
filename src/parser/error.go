package parser

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		fmt.Println("Error:", e.Error())
		os.Exit(1)
	}
}

func throwParsingLineError(e string, line string) {
	fmt.Println("Error:", e)
	fmt.Println("     Line:", line)
	os.Exit(1)
}

func throwParsingError(e string) {
	fmt.Println("Error:", e)
	os.Exit(1)
}
