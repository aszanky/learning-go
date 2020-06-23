package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//Printing the output
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"

	fmt.Print(v1, v2, v3, v4)
	fmt.Println()
	fmt.Println(v1, v2, v3, v4)
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)

	//Using Standart Output
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

	//Reading from standart input
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
