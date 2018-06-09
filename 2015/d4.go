package main

import (
	"fmt"
	"strconv"
	"encoding/hex"
	"crypto/md5"
)
func q4a(key string)(){ 
	start := [5]byte{0,0,0,0,0}
	// takes in string x and hashes it 
	// to have at least 5 zeros at the beginning
	var output []byte
	i := 0
	counter := 0
	for {
		// don't know when it will end at all
		i++
		text := key + strconv.Itoa(i)
		hash := md5.Sum([]byte(text))
		output = make([]byte, hex.EncodedLen(len(hash)))
		hex.Encode(output,hash[:])
		counter = 0
		for i,_ := range start{
			if output[i] == start[i]{
				counter++
				if counter == 5{
					fmt.Println(true)
					fmt.Println(output)
					break
				}
			}
		}
	}
}

func main(){
	q4a("pqrstuv")
}