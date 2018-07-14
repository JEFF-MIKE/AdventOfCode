package main;

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func d8i()(){
	totalCount := 0
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
		stringCount := 0
		charCount := 0
		x := 0
		for (x < (len(n)-1)) {
			if n[x] != "\\" {
				charCount += 1
				stringCount += 1
				x++
				continue
			} else {
				if (n[x+1] == "\\") || (n[x+1] == "\""){
					stringCount+=2 // count backslash
					charCount++
					x+=2
					continue
				} else if n[x+1] == "x"{
					stringCount+=4 // count backslash + following characters
					charCount+=1 // hex counts as single character
					x+=4 // jump 4 spaces			
				}
			}	
		}
		fmt.Println("CharCount:",charCount-1,"stringCount",stringCount+1)
		totalCount += ((stringCount+1)-(charCount-1))
	}
	fmt.Println(totalCount)
}

func main(){
	d8i()
}