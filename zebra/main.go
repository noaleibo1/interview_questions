package main

import "fmt"

func main() {

}

func mirrorLetters(str string) string {
	var lettersArray []string
	for i := 0; i < len(str); i++ {
		if isRelevant(string(str[i])) {
			lettersArray = append(lettersArray, string(str[i]))
		}
	}

	strToReturn := ""
	j := len(lettersArray) - 1
	for i := 0; i < len(str); i++ {
		if isRelevant(string(str[i])) {
			strToReturn = fmt.Sprintf("%s%s", strToReturn, lettersArray[j])
			j--
		} else {
			strToReturn = fmt.Sprintf("%s%c", strToReturn, str[i])
		}
	}

	return strToReturn
}

func isRelevant(c string) bool {
	return c == "a" || c == "e" || c == "o"
}
