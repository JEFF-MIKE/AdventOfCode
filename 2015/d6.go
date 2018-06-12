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
	var input [1000][1000]int // filled
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
			x2,_ := strconv.Atoi(n[4])
			y1,_ := strconv.Atoi(n[2])
			y2,_ := strconv.Atoi(n[5])
			xstart := x1
			ystart := y1
			for y2 >= ystart {
				for x2 >= xstart {
					if input[ystart][xstart] == 1{
						input[ystart][xstart] = 0 
					} else {
						input[ystart][xstart] = 1
					}
					xstart++
				}
				ystart++
				xstart = x1
			}
		} else {
			// either turn off the lights or turn them on
			x1,_ := strconv.Atoi(n[2])
			x2,_ := strconv.Atoi(n[5])
			y1,_ := strconv.Atoi(n[3])
			y2,_ := strconv.Atoi(n[6])
			xstart := x1
			ystart := y1
			num := 0
			if n[1] == "on"{
				num = 1
			} else {
				num = 0
			}
			for y2 >= ystart{
				for x2 >= xstart{
					input[ystart][xstart] = num
					xstart++
				}
				ystart++
				xstart = x1
			}
		}
	}
	total := 0 
	for p := 0;p<len(input);p++{
		for q := 0;q<len(input[p]);q++{
			total += input[p][q]
		}
	}
	fmt.Println(total)
	//fmt.Println(input)
}

func main(){
	q6a()
}