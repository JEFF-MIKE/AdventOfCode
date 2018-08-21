package main;
import (

	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type Reindeer struct{
	kms int
	sec int
	rest int
}

type TravelTime struct{
	timeOut int
	secondsLeft int
}


/*
	Two part to this question
	First part is the faster method, but as it turns out 
	the for loop condition is now completely different so 
	the 2nd function is forced to go through the loop by every second as opposed to 
	just their seconds travelled and rest intervals.
*/

func d14i()(){
	d := make(map[string]int)
	totalTimer := 2503 // amount of seconds we use to travel
	// setup a map of reindeer:distance travelled
	dps := make(map[string]Reindeer)
	// dps is a map containing our reindeer structs
	f,err := os.Open("d14.txt")
	if err != nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	for i := range lines{
		splitter := strings.Split(lines[i]," ")
		currentReindeer := splitter[0]
		d[currentReindeer] = 0
		kms,err := strconv.Atoi(splitter[3])
		if err != nil{
			fmt.Println(err)
		}
		sec,err := strconv.Atoi(splitter[6])
		if err != nil{
			fmt.Println(err)
		}
		rest,err := strconv.Atoi(splitter[13])
		if err != nil{
			fmt.Println(err)
		}
		x := Reindeer{kms,sec,rest}
		dps[currentReindeer] = x
		// now deal with the reindeer actually travelling
		localTimer := totalTimer
		for localTimer > 0{
			// Only go through the loop one piece at a time when there
			// is less time remaining than the amount spent
			if ((localTimer - dps[currentReindeer].sec) < 0){
				// make a for loop here 
				timeLeft := localTimer
				d[currentReindeer] += (dps[currentReindeer].kms * timeLeft)
				break
			} else{
				// first register the deer travel time
				localTimer -= dps[currentReindeer].sec 
				d[currentReindeer] += (dps[currentReindeer].kms * dps[currentReindeer].sec)
				// then negate the rest time and continue the loop
				localTimer -= dps[currentReindeer].rest
			}
		}
	}
	fmt.Println(dps)
	fmt.Println(d)
}

func d14ii()(){
	score := make(map[string]int) // score of each reindeer
	d := make(map[string]int) 
	totalTimer := 2503 // amount of seconds we use to travel
	// setup a map of reindeer:distance travelled
	dps := make(map[string]Reindeer)
	reindeerStatus := make(map[string]*TravelTime) // keep a record of the reindeers who are out and those who are in
	// dps is a map containing our reindeer structs
	f,err := os.Open("d14.txt")
	if err != nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	for i := range lines{
		splitter := strings.Split(lines[i]," ")
		currentReindeer := splitter[0]
		d[currentReindeer] = 0
		kms,err := strconv.Atoi(splitter[3])
		if err != nil{
			fmt.Println(err)
		}
		sec,err := strconv.Atoi(splitter[6])
		if err != nil{
			fmt.Println(err)
		}
		rest,err := strconv.Atoi(splitter[13])
		if err != nil{
			fmt.Println(err)
		}
		x := Reindeer{kms,sec,rest}
		dps[currentReindeer] = x
		score[currentReindeer] = 0
		reindeerStatus[currentReindeer] = &TravelTime{secondsLeft:sec,timeOut:0}
	}
	// gotta deal with the reindeer travelling at every second
	for i:=0;i<totalTimer;i++{
		for key,_ := range d{
			if reindeerStatus[key].secondsLeft > 0{
				reindeerStatus[key].secondsLeft-=1
				d[key] += dps[key].kms
				if reindeerStatus[key].secondsLeft == 0{
					reindeerStatus[key].timeOut = dps[key].rest
				}
			} else {
				reindeerStatus[key].timeOut -= 1
				if reindeerStatus[key].timeOut == 0{
					reindeerStatus[key].secondsLeft = dps[key].sec
				}
			}
		}
		max := 0
		fastestReindeers := make(map[string]int)
		fastestReindeer := ""
		// do one loop to find the absolute maximum so far
		for key,value := range d {
			if max < value{
				max = value
				fastestReindeer = key
			} 
		}
		fastestReindeers[fastestReindeer] = max
		for key,value := range d {
			if value == max && fastestReindeers[key] == 0 {
				fastestReindeers[key] = value
			}
		}
		for name,_ := range fastestReindeers {
			score[name] += 1
		}
	}
	fmt.Println(score)
}

func main(){
	fmt.Println("Answer to part 1: ")
	d14i()
	fmt.Println("Answer to part 2: ")
	d14ii()
}