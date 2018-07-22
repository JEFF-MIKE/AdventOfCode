package main;

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func d18i(){
	f,err := os.Open("d18.txt")
	if err != nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	lightArray := [100][100]string{}
	//g := &lightArray
	// will setup inital array first before
	// making a separate function
	for i := range lines{
		n := strings.Split(lines[i],"")
		fmt.Println(n)
		for x := range n{
			if n[x] == "#"{
				lightArray[i][x] = "#"
			} else {
				// it's off
				lightArray[i][x] = "."
			}
		}
	}
	fmt.Println(lightArray)
}

func main(){
	d18i()
}