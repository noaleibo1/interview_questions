package main

const (
	ErrorTreeIsGraph = "This is a graph"
	ErrorInvalidTree = "This is an invalid tree"

	// source: https://codereview.stackexchange.com/questions/121924/regular-expression-matching-with-string-slice-in-go
	RegexForDeserializationInput = `^\\w+?,\\s\\w+?,\\s\\w+(?:\\s\\w+)?,\\s\\d{8},\\s(?:\\d{3}\\s){2}\\d{4}$`
)
