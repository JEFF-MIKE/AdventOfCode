package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()(){
	// opens file
	x := 0
	f, _ := os.Open("q1.txt")
	scanner := bufio.NewScanner(f)
	for _,char := range scanner.Text(){
		if (string(char) == "("){
			x++
		} else {
			x--
		}
	}
	fmt.Println(x)
}