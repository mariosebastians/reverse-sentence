package main

import "fmt"

func main() {
	test := "nama saya mario sebastian santoso"
	fmt.Println(reverses(test))
}

func reverses(sentence string) string {
	temp := ""
	sentences := []string{}

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			sentences = append(sentences, temp)
			temp = ""
		} else {
			temp += string(sentence[i])
		}
	}

	// memastikan tidak ada string yang ketinggalan
	if temp != "" {
		sentences = append(sentences, temp)
		temp = ""
	}

	for i := len(sentences) - 1; i >= 0; i-- {
		temp += sentences[i]
		if i != 0 {
			temp += " "
		}
	}

	return temp
}
