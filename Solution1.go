package main
import (
	"fmt"
	"strings"
)
func freqOfLetters( setOfStrings []string,char chan string){
	for i:=0;i<len(setOfStrings);i++ {
		sli :=  strings.Split(setOfStrings[i], "")
		for j:=0;j<len(sli); j++ {
			char <- sli[j]
		}
	}
	close(char)
}
func main() {
	m:=make(map[string]int)
	setofStrings := []string{"tarun","varun"}
	char:= make(chan string)
	go freqOfLetters(setofStrings,char)
	for chars:= range char{
		m[chars]+=1
	}
	asciiNum:=97
	fmt.Println("{")
	for i:=0;i<26;i++{
		character := string(asciiNum+i)
		fmt.Printf("  %q: ",character)
		fmt.Println(m[character])
	}
	fmt.Println("}")
}