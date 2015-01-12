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

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

const wordsLocation = "/usr/share/dict/words"

var err error
var foundWord string
var wordLength int
var playerGuess string
var gameWord string
var signalChan = make(chan os.Signal, 1) // channel to catch ctrl-c

// clearScreen runs a shell clear command
func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

type game struct {
	word       string
	guessed    map[int]string
	blankWord  string
	maxGuesses int
}

func (g *game) wordLength() int { return len(g.word) }

func (g *game) genStats() {
	var guessCount int
	var guesses bytes.Buffer
	for _, v := range g.guessed {
		guessCount++
		guesses.Write([]byte(fmt.Sprintf("%s, ", v)))
	}
	fmt.Printf("\nGuesses: %d, Guessed: %s", guessCount, guesses.String())
	os.Exit(1)
}

// selectWord will search the installed dictionary for a word that meets
// the length criteria
func selectWord(length int) (string, error) {
	file, err := os.Open(wordsLocation)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var words []string
	rand.Seed(time.Now().UTC().UnixNano())
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == length {
			words = append(words, scanner.Text())
		}
	}
	return words[rand.Intn(len(words))], nil
}

func (g *game) setup() {
	fmt.Print("  Press 0 to enter your own word or 1 to generate one: ")
	var wordSetupAnswer int
	fmt.Scanf("%d", &wordSetupAnswer)
	for {
		switch {
		case wordSetupAnswer == 0:
			fmt.Print("  Enter word: ")
			fmt.Scanf("%s", &gameWord)
			g.word = gameWord
			break
		case wordSetupAnswer == 1:
			fmt.Print("  Enter length of word: ")
			fmt.Scanf("%d", &wordLength)
			if foundWord, err = selectWord(wordLength); err != nil {
				fmt.Println(err)
				continue
			}
			g.word = foundWord
			break
		default:
			fmt.Println("  not a valid entry.  Try again.")
			break
		}
		break
	}
	g.genBlankWord()
}

func (g *game) genBlankWord() {
	var preBlank bytes.Buffer
	preBlank.Write([]byte("\n "))
	for i := 1; i <= g.wordLength(); i++ {
		preBlank.Write([]byte("_ "))
	}
	g.blankWord = preBlank.String()
}

func (g *game) letterInWord(pg string) bool {
	for _, i := range g.word {
		if string(i) == pg {
			return true
		}
	}
	return false
}

func main() {
	clearScreen()
	fmt.Println("+ Hangman +")
	g := game{}
	signal.Notify(signalChan, os.Interrupt)
	// setup go routine to catch a ctrl-c
	go func() {
		for range signalChan {
			g.genStats()
		}
	}()
	g.setup()
	for {
		clearScreen()
		fmt.Println(boards[len(g.guessed)])
		fmt.Println(g.blankWord)
		fmt.Print("\nEnter guess: ")
		fmt.Scanf("%s", &playerGuess)
		if g.letterInWord(playerGuess) {
			//
		} else {
			clearScreen()
			g.guessed[len(g.guessed)+1] = playerGuess
			fmt.Println(boards[len(g.guessed)])
			fmt.Println(g.blankWord)
			continue
		}
	}
	os.Exit(0)
}

// boards is a map that holds the different stages of game play
var boards = map[int]string{
	0: `+ Hangman +








`,
	1: `+ Hangman +







 ____|____
`,
	2: `+ Hangman +

     |/
     |
     |
     |
     |
     |
 ____|____
`,
	3: `+ Hangman +
     _________
     |/
     |
     |
     |
     |
     |
 ____|____
`,
	4: `+ Hangman +
     _________
     |/      |
     |
     |
     |
     |
     |
 ____|____
`,
	5: `+ Hangman +
     _________
     |/      |
     |      (_)
     |
     |
     |
     |
 ____|____
`,
	6: `+ Hangman +
     _________
     |/      |
     |      (_)
     |       |
     |
     |
     |
 ____|____
`,
	7: `+ Hangman +
     _________
     |/      |
     |      (_)
     |       |/
     |
     |
     |
 ____|____
`,
	8: `+ Hangman +
     _________
     |/      |
     |      (_)
     |      \|/
     |
     |
     |
 ____|____
`,
	9: `+ Hangman +
     _________
     |/      |
     |      (_)
     |      \|/
     |       |
     |
     |
 ____|____
`,
	10: `+ Hangman +
     _________
     |/      |
     |      (_)
     |      \|/
     |       |
     |        \
     |
 ____|____
`,
	11: `+ Hangman +
     _________
     |/      |
     |      (_)
     |      \|/
     |       |
     |      / \
     |
 ____|____
`,
}
