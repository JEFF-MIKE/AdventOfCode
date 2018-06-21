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
	var taken = make(map[string]map[string]int)	
	f,err := os.Open("d13a.txt")
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
		if m[recipient][n[0]] != 0 {
			// adjust the dictionary
			adjusted := m[n[0]][recipient] + m[recipient][n[0]]
			m[n[0]][recipient] = adjusted
			m[recipient][n[0]] = adjusted
		}
	}
	for key,_ := range m{
		// fill up taken with dictionaries.
		// if there are more than two in the nested dictionary,
		// we know an adjustment has to be made
		var z =  make(map[string]int)
		taken[key] = z
	}
	total := 0
	for key1,value1 := range m{
		// max2 is second highest, max1 is highest
		ref1,ref2 := "",""
		max1,max2 := 0,0
		for key2,_ := range value1{
			if m[key1][key2] > max2 {
				max2 = m[key1][key2]
				ref2 = key2 
				if max2 > max1{
					max2,max1 = max1,max2
					ref2,ref1 = ref1,ref2
				}
			}
		taken[key1][ref1] = max1
		taken[key1][ref2] = max2
		total += (max1 + max2)
		}
	}
	fmt.Println(m)
	fmt.Println(taken)
	fmt.Println(total)
	
}

func main() {
	d13a()
}