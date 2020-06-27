package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("Guess a number 1 - 100. Are you ready?")
		numberGuessed, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(numberGuessed) == "yes" {
			break
		}
	}
	min := 0
	max := 101
	for {
		guessedNumber := min + (max-min)/2
		fmt.Println("Is number", guessedNumber, "<,>,= number you guessed?")
		comparisonAnswer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(comparisonAnswer) == ">" {
			min = guessedNumber
		}
		if strings.TrimSpace(comparisonAnswer) == "<" {
			max = guessedNumber
		}
		if strings.TrimSpace(comparisonAnswer) == "=" {
			fmt.Println("Well done")
			break
		}
	}
}
