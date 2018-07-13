package main;

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func d8i()(){
	f,err := os.Open("d8.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	for i := range lines{
		n := strings.Split(lines[i],"")
		fmt.Println(n)
	}
}

func main(){
	d8i()
}