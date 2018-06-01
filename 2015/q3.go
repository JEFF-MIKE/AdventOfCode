package main

import (
	"fmt"
	b "bufio"
	"os"
	//"strconv"
	"strings"
)

func q3a(){
	f,err := os.Open("q3.txt")
	if err != nil {
		fmt.Println(err)
		return	
	}
	defer f.Close()
	scanner := b.NewScanner(f)
	scanner.Split(b.ScanLines)
	var lines []string
	var table [1000][1000]int //2d table of size 1000 
	for p,_ := range table{
		for g,_ := range table[p]{
			table[p][g] = 0
		}
	}
	x := 499
	y := 499
	for scanner.Scan() {
		lines = append(lines,scanner.Text())
	}
	for i,v := 0,len(lines);i<v;i++{
		n := strings.Split(lines[i],"")
		for z := 0;z<len(n);z++{
			switch n[z]{
			case "^":
				y -=1
			case ">":
				x += 1
			case "<":
				x -= 1
			case "v":
				y += 1
			default:
				fmt.Println("Yeah I don't even know tbh...")
			}
			table[x][y] += 1
		}
	}
	counter := 0
	for k,_ := range table{
		for v,_ := range table[k]{
			if table[k][v] > 0{
				counter++
			}
		}
	}
	fmt.Println(counter)
}


func main(){
	q3a()
}