package main;

import (
	"io/ioutil"
	"fmt"
	//"reflect"
	"encoding/json"
	
)

func d12(){
	f,_ := ioutil.ReadFile("d12.txt")
	var obj map[string]interface{}
	x := json.Unmarshal([]byte(f),&obj)
	if x != nil {
		panic(x)
	}
	//fmt.Println(obj)
	fmt.Println(obj["e"])
	//for _,value := range 
}

func main(){
	d12()
}