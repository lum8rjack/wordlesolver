// This method starts with 3 guesses all with
// different letters to remove 15 characters.
package main

func Method2(game *Game, numberOfGuesses int) GameResult {

	gameresult := GameResult{
		Answer:  game.Answer,
		Guesses: 0,
		Solved:  false,
	}

	var err error

	// Set our first guess
	newguess := "raise"
	res := game.Guess(newguess)
	gameresult.Solved = res.Solved
	gameresult.Guesses++

	// Return if we guessed correctly
	if gameresult.Solved {
		return gameresult
	}

	game.PossibleGuesses = RemoveIncorrectPosition(game.CorrectPosition[:], game.PossibleGuesses)

	// Second guess
	newop := GetUniqueLetterOptions(game.PossibleGuesses)
	newguess, err = GetRandomWord(newop)
	if err != nil {
		return gameresult
	}
	res = game.Guess(newguess)
	gameresult.Solved = res.Solved
	gameresult.Guesses++

	// Return if we guessed correctly
	if gameresult.Solved {
		return gameresult
	}

	game.PossibleGuesses = RemoveIncorrectPosition(game.CorrectPosition[:], game.PossibleGuesses)

	// Third guess
	newop2 := GetUniqueLetterOptions(game.PossibleGuesses)
	newguess, err = GetRandomWord(newop2)
	if err != nil {
		return gameresult
	}
	res = game.Guess(newguess)
	gameresult.Solved = res.Solved
	gameresult.Guesses++

	// Continue with the remaining guesses
	for game.Guesses < numberOfGuesses && !gameresult.Solved {
		game.PossibleGuesses = RemoveIncorrectPosition(game.CorrectPosition[:], game.PossibleGuesses)

		// New guess
		newguess, err = game.GetRandomWord()
		if err != nil {
			break
		}

		res := game.Guess(newguess)
		gameresult.Solved = res.Solved
		gameresult.Guesses++
	}

	return gameresult
}
