package main;

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
/* 
	easy enough, the parameters of function d23
	are x and y which represent a and b respectively.
*/
func d23(x,y int){
	var lines [][]string
	f,err := os.Open("d23.txt")
	if err != nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		lines = append(lines,strings.Split(scanner.Text()," "))
	}
	// these must be integers
	var a int = x;
	var b int = y;
	index := 0
	for index < len(lines){
		if len(lines[index]) == 2{
			switch lines[index][0]{
			case "hlf":
				if lines[index][1] == "a"{
					a=a/2
				} else {
					b=b/2
				}
				index++
			case "tpl":
				if lines[index][1] == "a"{
					a=a*3
				} else {
					b=b*3
				}
				index++
			case "inc":
				if lines[index][1] == "a"{
					a++
				} else {
					b++
				}
				index++
			case "jmp":
				e,_ := strconv.Atoi(lines[index][1])
				index += e		
			continue
			}
		} else {
			reg := strings.Trim(lines[index][1],",")
			if lines[index][0] == "jie"{				
				if reg == "a" {
					if a % 2 == 0{
						e,_ := strconv.Atoi(lines[index][2])
						index += e
						continue
					}
				} else{
					if b % 2 == 0{
						e,_ := strconv.Atoi(lines[index][2])
						index += e
						continue
					}
				}
				} else {
				if reg == "a"{
					if a == 1{
						e,_ := strconv.Atoi(lines[index][2])
						index += e
						continue
					}
				} else {
					if b == 1{
						e,_ := strconv.Atoi(lines[index][2])
						index += e
						continue
					}
				}
			}
			/* if it reaches here, that means the above jump
				commands failed. We just skip to the next
				instruction instead
			*/ 
			index++
			continue
		}
	}
	fmt.Println("a: ",a,"\n","b: ",b)
}
func main(){
	d23(0,0) // part 1, a=0,b=0
	d23(1,0) // part 2,a=1,b=0
}