package main

// boards is a map that holds the different stages of game play
var boards = map[int]string{
	1: `+ Hangman +








`,
	2: `+ Hangman +







 ____|____
`,
	3: `+ Hangman +

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
     |/
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
     |
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
     |
     |
     |
     |
 ____|____
`,
	7: `+ Hangman +
     _________
     |/      |
     |      (_)
     |       |
     |
     |
     |
 ____|____
`,
	8: `+ Hangman +
     _________
     |/      |
     |      (_)
     |       |/
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
     |
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
     |
     |
 ____|____
`,
	11: `+ Hangman +
     _________
     |/      |
     |      (_)
     |      \|/
     |       |
     |        \
     |
 ____|____
`,
	12: `+ Hangman +
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