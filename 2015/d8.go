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

func d8ii()(){
	/*
	Part two is MUCH trickier than expected.Turns out
	that due to how the input is formatted, it is actually
	impossible to just count it out like in part one without
	having countless if statements. New approach is to just
	build a new string to represent the new encoding format
	required
	*/
	totalCount := 0
	totalStringCount := 0
	totalCharCount := 0
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
		charCount := 0
		stringCount := 0
		n := strings.Split(lines[i],"")
		charCount = charCounter(n)
		stringCount = stringCounter(lines[i])
		totalCharCount += charCount
		totalStringCount += stringCount
		}
	totalCount = totalStringCount-totalCharCount
	fmt.Println("Total String count is :",totalStringCount)
	fmt.Println("Total char count is :",totalCharCount)
	fmt.Println("This is totalCount",totalCount)
}

func charCounter(lst []string)(k int){
	// does the normal character counting method
	// for use with part 2
	retInt := 0
	x := 0
	for (x < len(lst)){
		if lst[x] != "\\"{
			retInt++
			x++
			continue
		} else{
			// it's a backslash
			if lst[x+1] == "\\" || lst[x+1] == "\""{
				retInt+=2
				x+=2
			} else if lst[x+1] == "x"{
				retInt+=4
				x+=4
				continue
			}
		}
	}
	fmt.Println(lst)
	fmt.Println("charCounter done,charcount is: ",retInt)
	return retInt
}

func stringCounter(oldstr string)(total int){
	n := strings.Split(oldstr,"")
	oldstr = strings.Replace(oldstr,"\\","\\\\",-1)
	oldstr = strings.Replace(oldstr,"\"","\\\"",-1)
	z := strings.Split(oldstr,"")
	fmt.Println(len(z)+2)
	fmt.Println(n)
	fmt.Println(z)
	return len(z)+2
}
func main(){
	d8i()
	d8ii()
}

