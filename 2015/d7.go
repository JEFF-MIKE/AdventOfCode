package main;

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func q7a(){
	f,err := os.Open("q7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string // empty array to hold our input
	for scanner.Scan(){
		lines = append(lines,scanner.Text())	
	}
	// lines holds our input
	for i := range lines{
		m := strings.Split(lines[i]," ")
		fmt.Println(m)
	}
}

func main(){
	q7a()
}