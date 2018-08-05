/* 
day  20 of advent of code 2015
*/ 
package main;

import (
	"fmt"
)
func d20i(s int)(z int){
	// number is 34,000,000
	houseList := [34000000]int{} // holds all the houses
	elfNumber := 1 
	index := 0
	multiplier := 10
	flag := true
	for flag {
		houseList[index] += elfNumber * multiplier
		index += elfNumber
		if index >= len(houseList){
			elfNumber++
			index = elfNumber - 1
		}
		if elfNumber == (len(houseList)+1){
			flag = false
		}
	}
	for i := range houseList {
		if houseList[i] >= s{
			return (i+1)
		}
	}
	return 0
}

func d20ii(s int)(z int){
	// number is 34,000,000
	houseList := [34000000]int{} // holds all the houses
	elfNumber := 1 
	index := 0
	multiplier := 11
	flag := true
	presentCount := 0 //when this goes to 50, elf number increases
	for flag {
		houseList[index] += elfNumber * multiplier
		index += elfNumber
		presentCount++
		if index >= len(houseList){
			// We do not want them to 
			// visit the same house twice, so
			// we increasethe elfNumber here anyway
			elfNumber++
			index = elfNumber-1
		}
		if presentCount == 50{
			presentCount = 0
			elfNumber++
			index = elfNumber-1
		}
		if elfNumber == (len(houseList)+1){
			flag = false
		}
	}
	for i := range houseList {
		if houseList[i] >= s{
			return (i+1)
		}
	}
	return 0
}

func main(){
	//fmt.Println(d20i(20))
	//fmt.Println(d20i(34000000))
	fmt.Println(d20ii(34000000))
}