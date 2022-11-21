package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var pl = fmt.Println
	s1 := "A word"
	replacer := strings.NewReplacer("A", "Another") // Word replace
	s2 := replacer.Replace(s1)
	
	pl(s2, len(s2))
	pl("Contains another :", strings.Contains(s2, "Another")) // check if it contains the word
	// Index returns the index of the first instance of substr in s
	pl("o index :", strings.Index(s2, "o"))
	pl(strings.Replace(s2, "o", "0", -1)) // if n < 0 there is no limit on the number of replacecment

	s3 := "\nSome Words\n"
	s3 = strings.TrimSpace(s3) // Remove the space
	pl("split: ", strings.Split("a-b-c-d", "-"))
	pl("lower: ", strings.ToLower(s3))
	pl("upper: ", strings.ToUpper(s3))
	pl("prefix: ", strings.HasPrefix("tacocat", "taco"))
	pl("suffix: ", strings.HasSuffix("tacocat", "taco"))


	str_rand := "abcdefg"
	pl("Rune Count: ", utf8.RuneCountInString(str_rand))
	for i, runeVal := range str_rand {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

}