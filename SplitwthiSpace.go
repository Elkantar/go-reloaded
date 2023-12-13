package main

import "fmt"

func Punctuation(char string) bool {
	return char == "." || char == "," || char == "!" || char == "?" || char == ":" || char == ";"
}

func Voyelle(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'y' || c == 'h' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'Y' || c == 'H'
}

func apostrophe(c string) bool {
	return c == "'" || c == "`" || c == "′"
}

func RuneToInt(c rune) int {
	var i int
	if c > '0' && c <= '9' {
		if c == '2' {
			i = 2
		}
		if c == '3' {
			i = 3
		}
		if c == '4' {
			i = 4
		}
		if c == '5' {
			i = 5
		}
		if c == '6' {
			i = 6
		}
		if c == '7' {
			i = 7
		}
		if c == '8' {
			i = 8
		}
		if c == '9' {
			i = 9
		}
	}
	return i
}

func SplitWhiteSpaces(s string) []string {
	var out []string = []string{}
	var current string = ""

	for _, char := range s {
		if char == '\n' || char == '\t' || char == ' ' {
			if current != "" {
				out = append(out, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}

	if current != "" {
		out = append(out, current)
	}

	return out
}

func Params(nbr int) bool {
	if nbr == 2 {
		return true
	} else if nbr >= 2 {
		fmt.Println("Too many arguments")
		return false
	} else {
		fmt.Println("File name missing")
		return false
	}
}

// Upper letter of the word
func UP(s string) string {
	sentence := []rune(s)
	for ch := range sentence {
		if sentence[ch] >= 'a' && sentence[ch] <= 'z' {
			sentence[ch] = sentence[ch] - 32
		}
	}
	return string(sentence)
}

// Lower letter of the word
func LOW(s string) string {
	sentence := []rune(s)
	for ch := range sentence {
		if sentence[ch] >= 'A' && sentence[ch] <= 'Z' {
			sentence[ch] = sentence[ch] + 32
		}
	}
	return string(sentence)
}

// Capitalize 1st Letter of the word
func CAP(s string) string {
	ar := []rune(s)

	lettre := true

	for i := 0; i < len(s); i++ {
		if AlphaNumeric(ar[i]) && lettre {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = 'A' - 'a' + ar[i]
			}
			lettre = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] = 'a' - 'A' + ar[i]
		} else if !AlphaNumeric(ar[i]) {
			lettre = true
		}
	}
	return string(ar)
}

func AlphaNumeric(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

/*
func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}*/
