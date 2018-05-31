package main

import (
	"fmt"
	"io/ioutil"
)

func q1a()(){
	// opens file
	b, err := ioutil.ReadFile("q1.txt")
	if err != nil{
		fmt.Print(err)
	}
	x := string(b)
	counter := 0
	for _,char := range x {
		if string(char) == "("{
			counter++
		} else if string(char) == ")"{
			counter--
		} else {
			continue // just ignore this
		}
	}
	fmt.Println("The answer is:",counter)
}
func main()(){
	q1a()
}
	