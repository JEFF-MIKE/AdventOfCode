package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"bytes"
)

func q5a(){
	f,err := os.Open("d5.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines) 
	var lines []string // holds out string input
	var stringList []string // will hold the split strings
	total := 0
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	end := 0
	vowelCount := 0 // counter for vowels
	double := 0 // need to account for double letters occuring
	v := map[rune]int{'a':1,'e':1,'i':1,'o':1,'u':1}
	// making a map for quick access time for exclusions
	ex := map[string]int{"ab":1,"cd":1,"pq":1,"xy":1}
	flag := true // doing a check 
	for i := range lines {
		vowelCount = 0
		// first check for vowels
		for _,rune := range lines[i]{
			if v[rune] == 1{
				vowelCount++
			}
		}
		if vowelCount < 3{
			continue // ignore, does not have 3 vowels
		}
		// now we split the string to carry out the other operations
		stringList = strings.Split(lines[i],"")
		p := 0
		q := 1
		end = len(stringList)
		double = 0
		flag = true
		for q < end {
			if stringList[p] == stringList[q]{
				double++ // are they the same?
			}
			buffer := bytes.Buffer{} // always create new buffer
			buffer.WriteString(stringList[p]) // p first
			buffer.WriteString(stringList[q]) // q second
			if ex[buffer.String()] == 1 {
				flag = false
				break
			}
			p++
			q++
		}
		if (flag) && (double > 0) {
			total++ // only reached if above are satisfied
		}
		
	}
	fmt.Println(total)
}
func q5b(){
	f,err := os.Open("d5.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines) 
	var lines []string // holds out string input
	var stringList []string // will hold the split strings
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	total := 0
	end := 0
	for i := range lines {
		// now we split the string to carry out the other operations
		stringList = strings.Split(lines[i],"")
		first := 0
		second := 1
		third := 2
		pair := map[string]int{}
		end = len(stringList)
		flag := false
		double := false
		k := 0
		q := 2
		fmt.Println(stringList)
		for q < end{
			if (stringList[k] == stringList[q]){
				fmt.Println("Flag is true for: ",stringList[k], "and", stringList[q])
				flag = true
				break
			} else {
				k++
				q++
			}
		}

		for third < end {
			if (stringList[first] == stringList[second]) && (stringList[second]==stringList[third]){
				flag = true
				first+=3
				second+=3
				third+=3
				continue // dont' record these pairs
			}
			firstBuffer := bytes.Buffer{} // always create new buffer
			secondBuffer := bytes.Buffer{}
			firstBuffer.WriteString(stringList[first])
			firstBuffer.WriteString(stringList[second])
			if pair[firstBuffer.String()] == 0 {
				pair[firstBuffer.String()] = 1
			} else {
				pair[firstBuffer.String()]++
			}
			secondBuffer.WriteString(stringList[second])
			secondBuffer.WriteString(stringList[third])
			if pair[secondBuffer.String()] == 0{
				pair[secondBuffer.String()] = 1 
			} else {
				pair[secondBuffer.String()]++
			}
			first+=3
			second+=3
			third+=3
		}
		for _,v := range pair {
			if v > 1 {
				double = true
				break
			}
		}
		if (flag) && (double) {
			fmt.Println(lines[i])
			fmt.Println(pair)
			total++ // only reached if above are satisfied
		}
		
	}
	fmt.Println(total)
}
func main() {
	q5a()
	q5b()
}



