package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

const (
	lose int = iota
	tie
	win
)

var player string
var givenAnswer int
var signalChan = make(chan os.Signal, 1)

type game struct {
	attempts int
	player   string
	pAnswer  *int
	cAnswer  int
	results  []int
}

func checkValidAnswer(pa *int) bool {
	if *pa == lose || *pa == tie || *pa == win {
		return true
	}
	/*
		if *pa > 0 && *pa <= 2 {
			return true
		}
	*/
	return false
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func (g *game) genStats() {
	var wins, loses, ties int
	for _, i := range g.results {
		switch {
		case i == win:
			wins++
		case i == lose:
			loses++
		case i == tie:
			ties++
		}
	}
	fmt.Printf("\n\n%s, here are your stats...\n", g.player)
	fmt.Printf("Rounds: %d, Wins: %d, Loses: %d, Ties: %d\n\n", len(g.results), wins, loses, ties)
	os.Exit(1)
}

func genAnswer() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(3)
}

func main() {
	clearScreen()
	fmt.Println("Rock-Paper-Scissors")
	fmt.Println("Enter 0 for rock, 1 for paper, and 2 for scissors")
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &player)
	g := game{
		player:   player,
		attempts: 0,
		results:  make([]int, 0),
	}
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			g.genStats()
		}
	}()
	for {
		fmt.Print("Enter answer: ")
		fmt.Scanf("%d", &givenAnswer)
		if !checkValidAnswer(&givenAnswer) {
			fmt.Println("invalid answer, try again")
			continue
		}
		g.attempts = g.attempts + 1
		g.pAnswer = &givenAnswer
		g.cAnswer = genAnswer()
		switch {
		case g.cAnswer%3+1 == *g.pAnswer:
			g.results = append(g.results, 2)
			fmt.Println("Win")
		case *g.pAnswer%3+1 == g.cAnswer:
			g.results = append(g.results, 0)
			fmt.Println("lose")
		default:
			g.results = append(g.results, 1)
			fmt.Println("tie")
		}
	}
	os.Exit(0)
}
