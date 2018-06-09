package main;

import (
	"strconv"
	"os"
	"fmt"
	"bufio"
	"strings"
	//"strconv"
)

func q6a(){
	// 0's are off
	// 1's are on
	f,err := os.Open("d6.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//var input [1000][1000]int // filled
	var lines []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	replace := strings.NewReplacer(","," ")
	for _, char := range lines{
		char = replace.Replace(char)
		n := strings.Split(char," ")
		if n[0] == "toggle"{
			// we gotta invert the lights that were turned off,
			// and the lights that were turned on
			//fmt.Println((strconv.Atoi(string(lines[i][1]))))
			x1,_ := strconv.Atoi(n[1])
			x2,_ := strconv.Atoi(n[2])
			y1,_ := strconv.Atoi(n[4])
			y2,_ := strconv.Atoi(n[5])
			xstart := x1
			ystart := y1
			for x2 >= xstart {
				for y2 >= ystart {
					fmt.Println(y1)
					y1++
				}
				fmt.Println(x1)
				x1++
			}
			fmt.Println(x1,x2,y1,y2)
		} else if n[1] == "on" {
			// turn on the lights
		} else {
			// turn off the lights
		}
	}
}

func main(){
	q6a()
}