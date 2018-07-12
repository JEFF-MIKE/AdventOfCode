package main;

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func d16()(z int){
	d := map[string]int{"children":3,"cats":7,"samoyeds":2,"pomeranians":3,"akitas":0,"vizslas":0,"goldfish":5,"trees":3,"cars":2,"perfumes":1}
	f,err := os.Open("d16.txt")
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
		n := strings.Split(lines[i]," ")
		value1,_ := strconv.Atoi(strings.Trim(n[3],","))
		value2,_ := strconv.Atoi(strings.Trim(n[5],","))
		value3,_ := strconv.Atoi(strings.Trim(n[7],","))
		if value1 == d[strings.Trim(n[2],":")]{
			if value2 == d[strings.Trim(n[4],":")]{
				if value3 == d[strings.Trim(n[6],":")]{
					x := strings.Trim(n[1],":")
					m,_ := strconv.Atoi(x)
					z = m
					break
				}
			}
		} else {
			continue
		}
	}
	fmt.Println(z)
	return z
}

func d16ii()(z int){
	d := map[string]int{"children":3,"cats":7,"samoyeds":2,"pomeranians":3,"akitas":0,"vizslas":0,"goldfish":5,"trees":3,"cars":2,"perfumes":1}
	f,err := os.Open("d16.txt")
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
		c1 := false
		c2 := false
		c3 := false
		n := strings.Split(lines[i]," ")
		value1,_ := strconv.Atoi(strings.Trim(n[3],","))
		value2,_ := strconv.Atoi(strings.Trim(n[5],","))
		value3,_ := strconv.Atoi(strings.Trim(n[7],","))
		string1 := strings.Trim(n[2],":")
		string2 := strings.Trim(n[4],":")
		string3 := strings.Trim(n[6],":")
		newLst := [3]string{string1,string2,string3}
		for p := range newLst{
			if newLst[p] == "cats" || newLst[p] == "trees"{
				if p == 0 {
					if value1 > d[newLst[p]]{
						c1 = true
					}
				} else if p == 1 {
					if value2 > d[newLst[p]]{
						c2 = true
					}
				} else if p == 2 {
					if value3 > d[newLst[p]]{
						c3 = true
					}
				}
			} else if newLst[p] == "pomeranians" || newLst[p] == "goldfish"{
				if p == 0 {
					if value1 < d[newLst[p]]{
						c1 = true
					} 
				} else if p == 1 {
					if value2 < d[newLst[p]]{
						c2 = true
					}
				} else if p == 2 {
					if value3 < d[newLst[p]]{
						c3 = true
					}
				}
			}
		}
		temp,_ := strconv.Atoi(strings.Trim(n[1],":"))
		if temp != 103 {
			if value1 == d[string1] || c1 == true{
				if value2 == d[string2] || c2 == true{
					if value3 == d[string3] || c3 == true{
						x := strings.Trim(n[1],":")
						m,_ := strconv.Atoi(x)
						z = m
						break
					}
				}
			} else {
				continue
			}
		}
	}	
	fmt.Println(z)
	return z
}
func main(){
	d16()
	d16ii()
}