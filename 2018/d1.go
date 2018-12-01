package main;

import (
	//"strconv"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func d18i() {
	answer := 0 // start at zero
	f, err := os.Open("d1.txt")
	if err != nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var numbers[] string // slice of numbers
	for scanner.Scan(){
		numbers = append(numbers,scanner.Text())
	}
	for i := range numbers{
		number := numbers[i]
		newNumber,_ := strconv.Atoi(number)
		answer += newNumber
	}
	fmt.Println(answer)
}

func d18ii() {
	answer := 0 // start at zero
	answerMap := make(map[int]int) 
	answerMap[0] = 1
	f, err := os.Open("d1.txt")
	if err != nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var numbers[] string // slice of numbers
	for scanner.Scan(){
		numbers = append(numbers,scanner.Text())
	}
	for (true){
		for i := range numbers{
			number := numbers[i]
			newNumber,_ := strconv.Atoi(number)
			answer += newNumber
			if answerMap[answer] == 0{
				answerMap[answer] = 1
			} else {
				fmt.Println(answer)
				os.Exit(0)
			}
		}
	}
}

func main() {
	d18i()
	d18ii()
}