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
var signalChan = make(chan os.Signal, 1) // channel to catch ctrl-c

// game holds the data collected during game play
type game struct {
	attempts int    // track how many rounds played
	player   string // player name
	pAnswer  *int   // pointer to given answer
	cAnswer  int    // computer's answer
	results  []int  // array to hold wins, loses, and ties
}

// checkValidAnswer makes sure the given answer is valid
func checkValidAnswer(pa *int) bool {
	if *pa == lose || *pa == tie || *pa == win {
		return true
	}
	return false
}

// clearScreen runs a shell clear command
func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// genStats outputs the game play statistics
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
	os.Exit(1) // Since it was a ctrl-c, exit non-zero
}

// genComputerAnswer will randomly generate a number used as an answer
func genComputerAnswer() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(3)
}

func main() {
	clearScreen()
	fmt.Print("+ Rock-Paper-Scissors +\n\n")
	fmt.Println("Enter 0 for rock, 1 for paper, and 2 for scissors")
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &player)
	g := game{
		player:   player,
		attempts: 0,
		results:  make([]int, 0),
	}
	signal.Notify(signalChan, os.Interrupt)
	// setup go routine to catch a ctrl-c
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
		g.cAnswer = genComputerAnswer()
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
}
