// This method is probably the common one used.
// Each guess is random after removing ones that
// don't have the invalid characters and only
// have valid ones in the correct position.
package main

func Method1(game *Game, numberOfGuesses int) GameResult {

	gameresult := GameResult{
		Answer:  game.Answer,
		Guesses: 0,
		Solved:  false,
	}

	var err error

	newguess, err := GetRandomWord(game.PossibleGuesses)
	if err != nil {
		return gameresult
	}

	// Counter for how many times we can guess
	for game.Guesses < numberOfGuesses && !gameresult.Solved {
		res := game.Guess(newguess)

		// Check the correct positions and remove options that don't match
		game.PossibleGuesses = RemoveIncorrectPosition(game.CorrectPosition[:], game.PossibleGuesses)

		gameresult.Solved = res.Solved
		gameresult.Guesses++
		// New guess
		newguess, err = game.GetRandomWord()
		if err != nil {
			break
		}
	}

	return gameresult
}
