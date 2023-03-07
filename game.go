package main

import (
	"errors"
	"strings"
)

type Game struct {
	Guesses         int
	PossibleGuesses []string
	Answer          string
	InvalidChars    []rune
	ValidChars      []rune
	CorrectPosition [5]rune
}

type Result struct {
	Solved  bool
	Guesses int
}

// Create a new game
func NewGame(answer string, words []string) (*Game, error) {
	newGame := Game{
		Answer:          answer,
		Guesses:         0,
		PossibleGuesses: words,
	}

	if len(answer) != 5 {
		return nil, errors.New("answer is not 5 characters")
	}

	return &newGame, nil
}

func (g *Game) GetRandomWord() (string, error) {
	return GetRandomWord(g.PossibleGuesses)
}

// Check if the guess is correct
func (g *Game) Guess(guess string) Result {
	// Add to the number of guesses
	g.Guesses++

	// Create result
	res := Result{
		Solved:  false,
		Guesses: g.Guesses,
	}

	if guess == g.Answer {
		res.Solved = true
		return res
	}

	// Check the characters of the guess
	for i, c := range guess {
		// Check if the character is in the answer
		if strings.ContainsRune(g.Answer, c) {
			// Check if the position is correct
			if rune(g.Answer[i]) == c {
				g.CorrectPosition[i] = c
			}

			g.ValidChars = append(g.ValidChars, c)

		} else {
			g.InvalidChars = append(g.InvalidChars, c)

			// Find words that contain the invalid char and add to
			// the list of items to remove
			var itemsToRemove []string

			for _, pg := range g.PossibleGuesses {
				if strings.ContainsRune(pg, c) {
					itemsToRemove = append(itemsToRemove, pg)
				}
			}

			if len(itemsToRemove) > 0 {
				for _, itr := range itemsToRemove {
					g.PossibleGuesses = RemoveWord(itr, g.PossibleGuesses)
				}
			}
		}
	}

	// Remove duplicates
	g.InvalidChars = RemoveDuplicateValues(g.InvalidChars)
	g.ValidChars = RemoveDuplicateValues(g.ValidChars)

	// Remove the guess from possible selection
	g.PossibleGuesses = RemoveWord(guess, g.PossibleGuesses)

	return res
}
