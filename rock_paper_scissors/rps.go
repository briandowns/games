// Copyright 2015 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This is an implementation of the game rock, paper, scissors.
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

var compStr string
var err error
var compAnswer int
var givenAnswer int

// channel variables
var (
	answerChan = make(chan *int, 1)
	validChan  = make(chan *int, 1)      // channel to receive answer for validity check
	validResp  = make(chan bool, 1)      // channel to send validity answer on
	signalChan = make(chan os.Signal, 1) // channel to catch ctrl-c
)

// game holds the data collected during game play
type game struct {
	attempts int   // track how many rounds played
	pAnswer  *int  // pointer to given answer
	cAnswer  int   // computer's answer
	results  []int // array to hold wins, loses, and ties
}

// checkValidAnswer makes sure the given answer is valid
func checkValidAnswer() {
	for {
		select {
		case a := <-validChan:
			if *a == lose || *a == tie || *a == win {
				validResp <- true
			} else {
				validResp <- false
			}
		}
	}
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
	fmt.Printf("\n\nRounds: %d, Wins: %d, Loses: %d, Ties: %d\n\n", len(g.results), wins, loses, ties)
	os.Exit(1) // Since it was a ctrl-c, exit non-zero
}

// genComputerAnswer will randomly generate a number used as an answer
func genComputerAnswer() (int, string, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	a := rand.Intn(3)
	switch {
	case a == 0:
		return a, "rock", nil
	case a == 1:
		return a, "paper", nil
	case a == 2:
		return a, "scissors", nil
	}
	return 127, "", fmt.Errorf("unknown value")
}

func main() {
	clearScreen()
	fmt.Print("+ Rock-Paper-Scissors (Enter 0 for ROCK, 1 for PAPER, and 2 for SCISSORS)\n")
	g := game{
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
	go checkValidAnswer()
Loop:
	for {
		clearScreen()
		fmt.Print("+ Rock-Paper-Scissors (Enter 0 for ROCK, 1 for PAPER, and 2 for SCISSORS)\n")
		fmt.Print("\nEnter answer: ")
		fmt.Scanf("%d", &givenAnswer)
		validChan <- &givenAnswer
		for i := range validResp {
			if !i {
				fmt.Println("invalid answer, try again")
				continue Loop
			}
			break
		}
		g.attempts = g.attempts + 1
		g.pAnswer = &givenAnswer
		if compAnswer, compStr, err = genComputerAnswer(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		g.cAnswer = compAnswer
		switch {
		case g.cAnswer%3+1 == *g.pAnswer:
			g.results = append(g.results, win)
			fmt.Printf("Computer: %s  WIN!\n", compStr)
		case *g.pAnswer%3+1 == g.cAnswer:
			g.results = append(g.results, lose)
			fmt.Printf("Computer: %s  LOSE\n", compStr)
		default:
			g.results = append(g.results, tie)
			fmt.Printf("Computer: %s  TIE\n", compStr)
		}
	}
}
