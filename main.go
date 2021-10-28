package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)
var choices = [3]string{"rock", "paper", "scissors"}

func main() {
	opponentChoice := make(chan string, 1)
	var userChoice string
	var wg sync.WaitGroup
	wg.Add(1)
	go randomChoice(&wg, opponentChoice)
	fmt.Printf("Your opponent is thinking\nPick one (rock, paper, or scissors?)\n")
	fmt.Scanln(&userChoice)
	contains(choices[:], userChoice)
	wg.Wait()
	resultChecker(userChoice, <-opponentChoice)
}

func randomChoice(wg *sync.WaitGroup, choice chan string) {
	sleepTime := randN.Intn(10)

	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("Your opponent is ready\n")
	choice <- choices[randN.Intn(len(choices))]
	defer wg.Done()
}

func resultChecker(user, opponent string) {
	fmt.Printf("opponent choice: " + opponent + "\n")

	if user == "scissors" && opponent == "paper" || user == "rock" && opponent == "scissors" || user == "paper" && opponent == "rock" {
		fmt.Printf("result = you win!")
	} else if user == "rock" && opponent == "paper" || user == "scissors" && opponent == "rock" || user == "paper" && opponent == "scissors" {
		fmt.Printf("result = you lose!")
	} else {
		fmt.Printf("result = draw")
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	fmt.Printf("invalid input\n")
	os.Exit(1)
	return false
}
