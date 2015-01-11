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
	"byte"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const wordsLocation = "/usr/share/dict/words"

type game struct {
	word    string
	guessed map[int]string
}

func (g *game) wordLength() int { return len(g.word) }

func (g *game) genStats() {
	var guessCount int
	var guesses bytes.Buffer
	for _, v := range g.guessed {
		guessCount++
		guesses.Write([]byte(fmt.Sprintf("%s, ", v)))
	}
	fmt.Printf("Guesses: %d, Guessed: %s", guessCount, guesses.String())
}

// selectWord will search the installed dictionary for a word that meets
// the length criteria
func selectWord(length int) string {
	file, err := os.Open(wordsLocation)
	if err != nil {
		log.Fatalln(err)
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
	return words[rand.Intn(len(words))]
}

func main() {
	//x := hangman{}
	fmt.Println(selectWord(5))
}
