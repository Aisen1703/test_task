package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	targetNumber := rand.Intn(10) + 1

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Угадайте число от 1 до 10!")

	for {

		fmt.Print("Ваше число:")
		scanner.Scan()
		userInput := strings.TrimSpace(scanner.Text())

		guess, err := strconv.Atoi(userInput)
		if err != nil || guess < 1 || guess > 10 {
			fmt.Println("Ошибка! Введенное число не попадает в диапазон от 1 до 10!")
			continue
		}

		if guess < targetNumber {
			fmt.Println("Загаданное число больше!")
		} else if guess > targetNumber {
			fmt.Println("Загаданное число меньше!")
		} else {
			fmt.Println("Вы угадали! Это", targetNumber)
			return
		}
	}
}
