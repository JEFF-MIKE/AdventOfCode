package main

import (
	"fmt"
	b "bufio"
	"os"
	"strings"
	"strconv"
)
func q2a()(){
	total := 0
	f,err := os.Open("q2.txt")
	if err != nil {
		fmt.Println(err)
		return	
	}
	defer f.Close()
	scanner := b.NewScanner(f)
	scanner.Split(b.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines,scanner.Text())
	}
	for i,v :=0,len(lines);i < v;i++{
		sum := 0
		n := strings.Split(lines[i],"x")
		l,_ := strconv.Atoi(n[0])
		w,_ := strconv.Atoi(n[1])
		h,_ := strconv.Atoi(n[2])
		x := (l*w) 
		y := (w*h)
		z := (h*l)
		min := x
		if y < min{
			min = y
		}
		if z < min{
			min = z
		}
		sum += min
		// only 2 elements, no need for "for" loop
		sum += ((x*2)+(y*2)+(z*2))
		total += sum
	}
	fmt.Println(total)
}
func q2b()(){
	total := 0
	f,err := os.Open("q2.txt")
	if err != nil {
		fmt.Println(err)
		return	
	}
	defer f.Close()
	scanner := b.NewScanner(f)
	scanner.Split(b.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines,scanner.Text())
	}
	for i,v :=0,len(lines);i < v;i++{
		sum := 0
		n := strings.Split(lines[i],"x")
		l,_ := strconv.Atoi(n[0])
		w,_ := strconv.Atoi(n[1])
		h,_ := strconv.Atoi(n[2])
		max := l
		result := (w+w+h+h)
		if w > max{
			max = w
			result = (l+l+h+h)
		}
		if h > max{
			max = h
			result = (l+l+w+w)
		}
		//why do I have a feeling the above could be done
		// so much better...
		sum +=result
		sum += (w*l*h)
		total += sum
	}
	fmt.Println(total)
}
func main(){
	q2a()
	q2b()
}

