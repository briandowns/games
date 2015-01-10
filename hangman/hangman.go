// Copyright 2014 Brian J. Downs
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
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
)

const (
	wordsLocation = "/usr/share/dict/words"
)

type hangman struct {
	word   string
	length int
}

type player struct {
	guesses int
	guessed []string
}

//func (h *hangman) String() string {}
//func (p *player) String() string  {}

// selectWord will search the installed dictionary for a word that meets
// the length criteria
func (h *hangman) selectWord(length int) {
	file, err := os.Open(wordsLocation)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == length {
			words = append(words, scanner.Text())
		}
	}
	max := *big.NewInt(int64(len(words)))
	n, err := rand.Int(rand.Reader, &max)
	if err != nil {
		log.Fatalln(err)
	}
	h.word = words[int(n)]
}

func main() {
	x := hangman{}
	x.selectWord(5)
	fmt.Println(x.word)
}
