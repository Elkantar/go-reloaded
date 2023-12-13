package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	var sentence string
	var sentence2 string
	var sentence3 []string
	var sentence4 []string
	var sentence5 string
	var low string
	var i int

	lengthOfArg := len(os.Args[1:])
	if Params(lengthOfArg) {

		file, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_WRONLY, 0o600)
		defer file.Close() // on ferme automatiquement à la fin de notre programme

		if err != nil {
			panic(err)
		}

		e := os.Remove("results.txt")
		if e != nil {
			fmt.Println(err)
			return
		}

		// data, _ := ioutil.ReadFile("sample.txt")
		file2, _ := os.OpenFile("results.txt", os.O_CREATE|os.O_WRONLY, 0o600)
		defer file2.Close() // on ferme automatiquement à la fin de notre programme

		content, _ := ioutil.ReadFile("sample.txt")
		sentence2 = string(content)
		sentence3 = SplitWhiteSpaces(sentence2)
		for a := range sentence3 {
			if a != 0 {
				if string(sentence3[a]) == "(cap)" {
					sentence3[a] = ""
					sentence3[a-1] = CAP(string(sentence3[a-1]))
				} else if string(sentence3[a]) == "(low)" {
					sentence3[a] = ""
					sentence3[a-1] = LOW(string(sentence3[a-1]))
					a++
				} else if string(sentence3[a]) == "(up)" {
					sentence3[a] = ""
					sentence3[a-1] = UP(string(sentence3[a-1]))
				} else if string(sentence3[a]) == "(bin)" {
					output, err := strconv.ParseInt(sentence3[a-1], 2, 64)
					if err != nil {
						fmt.Println(err)
						return
					}
					sentence3[a] = ""
					sentence3[a-1] = strconv.Itoa(int(output))
				} else if string(sentence3[a]) == "(hex)" {
					output, err := strconv.ParseInt(sentence3[a-1], 16, 64)
					if err != nil {
						fmt.Println(err)
						return
					}
					sentence3[a] = ""
					sentence3[a-1] = strconv.Itoa(int(output))
				} else if string(sentence3[a]) == "(cap," {
					low = sentence3[a+1]
					sentence3[a] = ""

					for _, c := range low {
						if c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
							i = RuneToInt(c)
						}
					}

					for i > 0 {
						sentence3[a-i] = CAP(string(sentence3[a-i]))
						i--
					}

					sentence3[a] = ""
					sentence3[a+1] = ""

				} else if string(sentence3[a]) == "(up," {
					low = sentence3[a+1]
					sentence3[a] = ""
					for _, c := range low {
						if c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
							i = RuneToInt(c)
						}
					}
					for i > 0 {
						sentence3[a-i] = UP(string(sentence3[a-i]))
						i--
					}
					sentence3[a] = ""
					sentence3[a+1] = ""
				} else if string(sentence3[a]) == "(low," {
					low = sentence3[a+1]
					sentence3[a] = ""
					for _, c := range low {
						if c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
							i = RuneToInt(c)
						}
					}
					for i > 0 {
						sentence3[a-i] = LOW(string(sentence3[a-i]))
						i--
					}
					sentence3[a] = ""
					sentence3[a+1] = ""
				} else if string(sentence3[a]) == "(up," {
					low = sentence3[a+1]
					sentence3[a] = ""
					for _, c := range low {
						if c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
							i = RuneToInt(c)
						}
					}
					for i > 0 {
						sentence3[a-i] = UP(string(sentence3[a-i]))
						i--
					}

					sentence3[a] = ""
					sentence3[a+1] = ""

				} else if string(sentence3[a]) == "(hex," {
					low = sentence3[a+1]
					sentence3[a] = ""
					for _, c := range low {
						if c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
							i = RuneToInt(c)
						}
					}

					for i > 0 {
						output, err := strconv.ParseInt(sentence3[a-i], 16, 64)
						if err != nil {
							fmt.Println(err)
							return
						}
						sentence3[a-i] = strconv.Itoa(int(output))
						i--
					}
					sentence3[a] = ""
					sentence3[a+1] = ""

				} else if string(sentence3[a]) == "(bin," {
					low = sentence3[a+1]
					sentence3[a] = ""
					for _, c := range low {
						if c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
							i = RuneToInt(c)
						}
					}

					for i > 0 {
						output, err := strconv.ParseInt(sentence3[a-i], 2, 64)
						if err != nil {
							fmt.Println(err)
							return
						}
						sentence3[a-i] = strconv.Itoa(int(output))
						i--
					}
					sentence3[a] = ""
					sentence3[a+1] = ""
				}
			}
		}
		var savecha string

		for p, cha := range sentence3 {
			savecha = ""
			for _, runes := range cha {
				if Punctuation(string(runes)) {
					savecha = savecha + string(runes) + " "
				} else {
					savecha = savecha + string(runes)
				}
			}
			sentence3[p] = savecha
		}
		for _, cha := range sentence3 {
			if cha != " " || cha != "" {
				sentence4 = append(sentence4, cha)
			}
		}

		for _, f := range sentence4 {
			sentence5 += f + " "
		}

		sentence6 := SplitWhiteSpaces(sentence5)
		for g, ch := range sentence6 {
			Punc := 0
			if g+1 < len(sentence6) {
				char2 := sentence6[g+1]
				for h := range char2 {
					if h+1 < len(char2) {
						if Punctuation(string(char2[h])) && Punctuation(string(char2[h+1])) {
							Punc = 1
						}
					}
				}
			}
			if ch == "a" {
				if g+1 < len(sentence6) {
					char2 := string(sentence6[g+1])
					for range char2 {
						if Voyelle(rune(char2[0])) {
							ch = "an"
						}
					}
				}
			}
			var apo int
			var saveapo string
			var savech string
			for n, t := range ch {
				if t == '‘' && n == 0 {
					apo = 1
					saveapo = "‘"
				} else if t == '‘' && n != 0 {
					apo = 2
					saveapo = "‘"
				}
			}
			for range ch {
				if apostrophe(string(ch[0])) && apostrophe(string(ch[len(ch)-1])) {
					apo = 3
				} else if apostrophe(string(ch[0])) {
					apo = 1
					saveapo = string(ch[0])
				} else if apostrophe(string(ch[len(ch)-1])) {
					apo = 2
					saveapo = string(ch[len(ch)-1])
				}
			}
			if apo == 2 {
				for l, apostrophe := range ch {
					if l == 0 {
						savech += saveapo
					}
					savech += string(apostrophe)

				}
				ch = savech
			}
			if apo == 1 {
				for l, apostrophe := range ch {
					savech += string(apostrophe)
					if l == len(ch)-1 {
						savech += saveapo
					}
				}
				ch = savech
			}

			len := len(sentence6)
			if g == 0 {
				if !Punctuation(sentence6[g+1]) && Punc != 1 && !apostrophe(sentence6[g]) && sentence6[g+1] != "‘" && sentence6[g] != "‘" {
					sentence = sentence + ch + " "
				} else {
					sentence = sentence + ch
				}
			} else if g+1 < len && g != 0 {
				if !Punctuation(sentence6[g+1]) && Punc != 1 && !apostrophe(sentence6[g]) && sentence6[g+1] != "‘" && sentence6[g] != "‘" {
					sentence = sentence + ch + " "
				} else {
					sentence = sentence + ch
				}
			} else {
				sentence += ch + "\n"
			}
		}
		_, err = file2.WriteString(string(sentence)) // écrire dans le fichier
		if err != nil {
			panic(err)
		}

		/*data2, err := ioutil.ReadFile("results.txt") // lire le fichier
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
		fmt.Println(string(data2))*/
	}
}
