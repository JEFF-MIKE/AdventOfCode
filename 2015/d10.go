package main;

import (
	"fmt"
	"strings"
	"strconv"
)
func d10a(input string){
	lst := strings.Split(input,"")
	counter := 0
	startCount := 0
	var x string
	for counter < 50 {
		var newLst []string // every iteration we want to create a newList
		for i := range lst {
			if startCount == 0 {
				x = lst[i]
				startCount++
			} else if lst[i] == x {
				// it is equal to x, and we want to continue the loop
				// increment the startCount and continue
				startCount++
			} else {
				// x is not equal, and we have to reset the startCount
				// aswell as append to newLst
				newLst = append(newLst,strconv.Itoa(startCount)) 
				newLst = append(newLst,x)
				startCount = 1
				x = lst[i]
			}
		}
		if startCount > 0 {
			newLst = append(newLst,strconv.Itoa(startCount))
			newLst = append(newLst,lst[(len(lst)-1)])
		}
		startCount = 0 // just reset it here
		lst = newLst
		counter++
	}
	fmt.Println(lst,len(lst))
}

func main(){
	/*
	d10a("111221")
	d10a("1211")
	d10a("21")
	d10a("12")
	*/
	d10a("1113222113")
}