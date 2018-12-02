package main

import (
	"fmt"
	"strings"
)

const text = `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`

func main() {
	//fmt.Printf("%s", text)

	textArray := strings.Split(text, "\n")

	//rowNum := 1
	for idx,line := range textArray {
		fmt.Printf("%d - %s\n", idx, line)

		charMap := countChars(line)

		//var countLine string
		for key, value := range charMap {

			fmt.Printf("'%s'=%d,", string(key), value)
			//countLine = "'"+string(key)+"'="+string(value)

			//fmt.Printf(countLine)
		}
		fmt.Println()
		//rowNum += 1
		//fmt.Printf("%d Size of Map = %d", rowNum, len(charMap))
		//keys := make([]rune, len(charMap))

		//fmt.Printf("%d %s\n", rowNum+1,keys)


	}
	//fmt.Printf("Num of lines is %d\n", len(textArray))

	//charCount = map[byte]int{}

	//for charIdx := range text {
	//	charCount[Inde] = charCount['y'] + 1
	//}

}

func countChars(s string)map[rune]int {
	//fmt.Println("Incoming String: " + s)
	//var m map[rune]int
	m := make(map[rune]int, 0)

	for _, value := range s {
		//fmt.Printf("%d -> %d\n", key, value)
		//fmt.Printf("%s - ", string(value))
		m[value] = m[value] + 1
	}
	//for _,c := range s {
	//	fmt.Printf("character - %s",c)
	//	m[c] = m[c] + 1
	//}

	return m
}
