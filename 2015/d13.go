package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
// decided to completely rework the code as
// it was just not initially working out at all
// so OOP is here to save the day.
// previous work down below

type person struct{
	name string // hold their name for removal
	favourablePerson map[string]int
}
func (p person) addPartner(matchString string,matchInt int)(){
	p.favourablePerson[matchString] = matchInt
}
func (p person) getMax()(string,int){
	maxPerson := ""
	maxInt := 0
	for key,value := range p.favourablePerson{
		if value > maxInt{
			// get the max first, then
			maxInt = value
			maxPerson = key
		}
	}
	return maxPerson,maxInt
}

func (p person) check()(bool){
	if len(p.favourablePerson) > 6{
		return true
	} else {
		return false
	}
}

func (p person) removeMax(){
	x,_ := p.getMax()
	// remove the max once we use it
	delete(p.favourablePerson,x)
}

type table struct{
	// an array of people
	seats []person
	totalCount int
}

func (t *table) addPerson(k person)([]person){
	t.seats = append(t.seats,k)
	return t.seats
}
func (t *table) giveMax()(string,int){
	// iterates through ALL the people to see who
	// has the highest value
	maxInt := 0
	maxName := ""
	// we will use index variable for removing from the
	// table list
	index := 0 
	for k := range t.seats {
		x,y := t.seats[k].getMax()
		if y > maxInt{
			maxInt = y
			maxName = x
			index = k
		}
	}
	t.totalCount += maxInt
	t.seats[index].removeMax()
	t.checkPerson(index)
	return maxName,maxInt
}
func (t table) checkPerson(i int)(){
	// it only has six people
	if t.seats[i].check() == false{
		t.seats[i],t.seats[(len(t.seats) - 1)] = t.seats[(len(t.seats)-1)],t.seats[i]
		t.seats = t.seats[:(len(t.seats)-1)]
	} else {
		
	}
}

func d13a(){
	var a *person
	var x []person
	seating := table{seats: x,totalCount: 0}
	// when personCounter goes to 7, we reset it
	personCounter := 0
	listIterator := 0
	var lines []string // input
	f,err := os.Open("d13.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for (scanner.Scan()){
		lines = append(lines,scanner.Text())
	}
	for i := range lines{
		n := strings.Split(lines[i]," ")
		if personCounter == 0{
			var o = make(map[string]int)
			p := person{name:n[0],favourablePerson: o}
			a = &p
			seating.addPerson(p)
		}
		number,_ := strconv.Atoi(n[3])
		if n[2] == "lose"{
			number = number *-1
		}
		a.addPartner(strings.Trim(n[10],"."),number)
		personCounter++
		if personCounter == 7 {
			personCounter = 0
			listIterator++
		}
	}
	for true{
		
	}
}

func main(){
	d13a()
}
/*
func d13a(){
	var lines []string // input
	//var p []map[string]map[string]int // holds maps with key:value pairs strings:ints
	var m = make(map[string]map[string]int) // key is name, value is x
	var taken = make(map[string]map[string]int)	
	f,err := os.Open("d13a.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for (scanner.Scan()){
		lines = append(lines,scanner.Text())
	}
	for i := range lines {
		n := strings.Split(lines[i]," ")
		// check to see if map already exists within dictionary
		if m[n[0]] == nil {
			// add it to our 2d map
			m[n[0]] = make(map[string]int) // make map before adding it			
		}
		number,_ := strconv.Atoi(n[3])
		recipient := strings.Trim(n[10],".")
		if n[2] == "gain" {
			m[n[0]][recipient] = number 
		} else {
			number,_ := strconv.Atoi(n[3])
			number *= -1
			m[n[0]][recipient] = number
		}
		if m[recipient][n[0]] != 0 {
			// adjust the dictionary
			adjusted := m[n[0]][recipient] + m[recipient][n[0]]
			m[n[0]][recipient] = adjusted
			m[recipient][n[0]] = adjusted
		}
	}
	for key,_ := range m{
		// fill up taken with dictionaries.
		// if there are more than two in the nested dictionary,
		// we know an adjustment has to be made
		var z =  make(map[string]int)
		taken[key] = z
	}
	total := 0
	for key1,value1 := range m{
		// max2 is second highest, max1 is highest
		ref1,ref2 := "",""
		max1,max2 := 0,0
		for key2,_ := range value1{
			if m[key1][key2] > max2 {
				max2 = m[key1][key2]
				ref2 = key2 
				if max2 > max1{
					max2,max1 = max1,max2
					ref2,ref1 = ref1,ref2
				}
			}
		taken[key1][ref1] = max1
		taken[key1][ref2] = max2
		total += (max1 + max2)
		}
	}
	fmt.Println(m)
	fmt.Println(taken)
	fmt.Println(total)
	
}

func main() {
	d13a()
}
*/