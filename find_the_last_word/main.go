package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getLastWordUsingFields(""))
	fmt.Println(getLastWord(""))

}

func getLastWordUsingFields(str string) (lastWord string) {
	words := strings.Fields(str)
	if len(words) > 0 {
		return words[len(words)-1]
	} else {
		return ""
	}
}

func getLastWord(str string) (lastWord string) {
	var i int
	for i = len(str)-1; i>=0; i--{ // Remove whitespace suffix
		if str[i] == ' '{
			continue
		} else {
			break
		}
	}
	for ; i >= 0; i--{
		if str[i] == ' '{
			break
		} else {
			lastWord = fmt.Sprintf("%s%s",string(str[i]), lastWord)
		}
	}
	return
}