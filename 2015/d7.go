package main;

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func q7a(){
	f,err := os.Open("d7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string // empty array to hold our input
	//var m map[string]string
	for scanner.Scan(){
		lines = append(lines,scanner.Text())	
	}
	// lines holds our input
	for i := range lines{
		x := strings.Split(lines[i]," ")
		if len(x) == 3 {
			fmt.Println(x)
		} else if (x[0] == "b") || x[2] == "b"{
			fmt.Println(x)
		}
		//fmt.Println(x,len(x))
	}
}

func main(){
	q7a()
}