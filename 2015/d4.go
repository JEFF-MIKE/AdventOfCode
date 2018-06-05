package main

import (
	"fmt"
	"strconv"
	"bytes"
	"encoding/hex"
	"crypto/md5"
)
func q4a(x string)(){ 
	start := []byte{0,0,0,0,0}
	// takes in string x and hashes it 
	// to have at least 5 zeros at the beginning
	var output []byte
	i := 0
	counter := 0
	for {
		// don't know when it will end at all
		i++
		b := bytes.Buffer{} // create buffer for creating hex number
		b.WriteString(x)
		b.WriteString(strconv.Itoa(1048970)) // convert i to a string, and append
		text := b.String()
		hash := md5.Sum([]byte(text))
		output = make([]byte, hex.EncodedLen(len(hash)))
		hex.Encode(output,hash[:])
		counter = 0
		for i,_ := range start{
			if output[i] == start[i]{
				counter++
				if counter == 5{
					fmt.Println(true)
					fmt.Printf(hex.EncodeToString(output))
				}
			}
		}
			fmt.Println(true)
			fmt.Println(hash)
		}
}

func main(){
	q4a("pqrstuv")
}