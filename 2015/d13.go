package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func d13a(){
	var lines []string // input
	//var p []map[string]map[string]int // holds maps with key:value pairs strings:ints
	var m = make(map[string]map[string]int) // key is name, value is x
	var marked = make(map[string]int)
	f,err := os.Open("d13.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for (scanner.Scan()){
		lines = append(lines,scanner.Text())
	}
	for i := range lines {
		n := strings.Split(lines[i]," ")
		// check to see if map already exists within dictionary
		if m[n[0]] == nil {
			// add it to our 2d map
			m[n[0]] = make(map[string]int) // make map before adding it			
		}
		number,_ := strconv.Atoi(n[3])
		recipient := strings.Trim(n[10],".")
		if n[2] == "gain" {
			m[n[0]][recipient] = number 
		} else {
			number,_ := strconv.Atoi(n[3])
			number *= -1
			m[n[0]][recipient] = number
		}
	}
	for key,_ := range m {
		marked[key] = 0
		}
	
	fmt.Println(m)
	fmt.Println(marked)
	
}

func main() {
	d13a()
}