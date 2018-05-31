package main

import (
	"fmt"
	"io/ioutil"
)
// using io/ioutil as opposed to bufio as I want to read
// entire file as a single string
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
func q1b()(){
	// opens file
	b, err := ioutil.ReadFile("q1.txt")
	if err != nil{
		fmt.Print(err)
	}
	x := string(b)
	position := 0
	counter := 0
	for _,char := range x {
		if string(char) == "("{
			counter++
		} else if string(char) == ")"{
			counter--
		} else {
			continue // just ignore this
		}
		position++
		if counter == -1{
			break
		}
	}
	fmt.Println("The position is:",position)
}
func main()(){
	q1a()
	q1b()
}
	