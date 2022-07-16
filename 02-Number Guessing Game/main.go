/*
Implement a number-guessing game in which the computer computes a four digit number
      as a secret number and a player tries to guess that number correctly. Player
      would enter her guess and the computer would produce a feedback on the positions of
      the digits.
      Four-digit number can't start with 0 and have repeating digits.
      Let's say the computer computes 2658 as a secret number to be guessed by the player.
      When player enters her guess such as 1234, the computer would display -1
      meaning that only one digit of 1234 exist in the secret number and its position
      is wrong. When the player enters 5678 the similarly the computer displays +2-1. And
      the game goes on until the player correctly guess the secret number and the computer
      displays +4. The game also keeps track of all guesses entered by the players so far
      and lists them when it displays its feedback to the player so that the player can
      compute her next guess correctly.
*/
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

func randNum() int {
	rand.Seed(time.Now().Unix())
	return (rand.Intn(10000-1000) + 1000)
}

func getInput() int {
	fmt.Print("Guess:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)

	return num
}

func checkInput(r, x int) int {
	rS := strconv.Itoa(r)
	xS := strconv.Itoa(x)
	res := 0
	for i := 0; i < len(xS); i++ {
		for j := 0; j < len(xS); j++ {
			if rS[i] == xS[j] {
				res++
			}
		}
	}
	return res
}

func main() {
	exit := 0
	for exit != 1 {
		fmt.Println("For exit please press 1")
		fmt.Println("For showing answer and exit press 2")
		r := randNum() //random number created
		var history string

		for true {
			x := getInput()
			if x == 1 {
				exit = 1
				break
			}
			if x == 2 {
				fmt.Println("Answer:", r)
				//exit = 1
				//break
			}
			res := checkInput(r, x)
			resS := strconv.Itoa(res)
			history += ("-" + resS)
			fmt.Println("History:", history)
			if res >= 4 && x == r {
				break
			}
		}
		fmt.Println("Congratulations")
	}
}
