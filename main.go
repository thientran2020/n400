package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	m "github.com/thientran2020/go-n400-study/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		// Clear screen + Welcome message
		fmt.Print("\033[H\033[2J")
		fmt.Printf(fmt.Sprintf("%s%s%s\n", m.BLUE, m.WELCOME_MESSAGE, m.COLOROFF))

		// Generate random number to query a random question
		number := rand.Intn(100)
		category_index := getIndex(m.CATEGORY_KEYS, number+1)
		sub_category_index := getIndex(m.SUBCATEGORY_KEYS, number+1)

		category := fmt.Sprintf("%s[__TOPIC_____] %s%s", m.GREEN, strings.ToUpper(m.CATEGORIES[category_index]), m.COLOROFF)
		sub_category := fmt.Sprintf("%s[__CATEGORY__] %s%s", m.YELLOW, strings.ToUpper(m.SUBCATEGORIES[sub_category_index]), m.COLOROFF)
		question := fmt.Sprintf("%s%s%s", m.UWHITE, fmt.Sprintf("QUESTION #%d:", number+1), m.COLOROFF)

		fmt.Printf("%s\n%s\n\n%s %s\n", category, sub_category, question, m.QUESTIONS[number])

		// Enter to display answer
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		asciiCode, err := consoleReader.ReadByte()
		if err != nil {
			os.Exit(0)
		}
		answer := fmt.Sprintf("%s%s%s", m.UWHITE, fmt.Sprintf("ANSWER #%d:", number+1), m.COLOROFF)
		if asciiCode == 10 {
			fmt.Printf("%s\n%s\n\n\n", answer, m.ANSWERS[number])
		}

		// Exit only if users enter 'Q' or 'q'
		fmt.Print(m.PROMT)
		consoleReader = bufio.NewReaderSize(os.Stdin, 1)
		asciiCode, err = consoleReader.ReadByte()
		if err != nil || asciiCode == 81 || asciiCode == 113 {
			fmt.Println(fmt.Sprintf("%s%s%s", m.RED, m.EXIT_MESSAGE, m.COLOROFF))
			os.Exit(0)
		}
	}
}

func getIndex(keys []int, number int) int {
	index := keys[len(keys)-1]
	for i := range keys {
		if keys[i] > number {
			index = keys[i-1]
			break
		}
	}
	return index
}
