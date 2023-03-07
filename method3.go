// This method starts with a word that has atleast
// 4 vowels in it to reduce the possible options.
package main

func Method3(game *Game, numberOfGuesses int) GameResult {

	gameresult := GameResult{
		Answer:  game.Answer,
		Guesses: 0,
		Solved:  false,
	}

	var err error

	// First guess should have atleast 4 unique vowls
	firstoptions, err := GetXLetterVowels(game.PossibleGuesses, 4)
	if err != nil {
		return gameresult
	}
	newguess, err := GetRandomWord(firstoptions)
	if err != nil {
		return gameresult
	}

	res := game.Guess(newguess)
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
		game.PossibleGuesses = RemoveIncorrectPosition(game.CorrectPosition[:], game.PossibleGuesses)
		gameresult.Solved = res.Solved
		gameresult.Guesses++
	}

	return gameresult
}
