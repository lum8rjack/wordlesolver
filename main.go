package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	numberOfGuesses = 6
)

type Outcome struct {
	Gamesplayed  int
	Gameswon     int
	Totalguesses int
}

func main() {
	today := flag.Bool("today", false, "Use today's answer for all games (default false)")
	threads := flag.Int("threads", 25, "Number of threads to use")
	numberOfGames := flag.Int("games", 100, "Number of games to play")
	wordlist := flag.String("wordlist", "dictionaries/wordle-03022023.txt", "Which wordlist to use")
	flag.Parse()

	fmt.Println("Wordlesolver v1.0")

	// If today is set, get today's answer
	var TODAYSANSWER string
	var err error

	if *today {
		TODAYSANSWER, err = GetTodaysAnswer()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Using today's answer for all games")
	}

	// Load in the word list
	words, err := LoadWordlist(*wordlist)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wordlist: %s (%d words)\n", *wordlist, len(words))

	resultsOutcome1 := Outcome{
		Gamesplayed:  0,
		Gameswon:     0,
		Totalguesses: 0,
	}

	resultsOutcome2 := Outcome{
		Gamesplayed:  0,
		Gameswon:     0,
		Totalguesses: 0,
	}

	resultsOutcome3 := Outcome{
		Gamesplayed:  0,
		Gameswon:     0,
		Totalguesses: 0,
	}

	// Setup number of go routines
	var wg sync.WaitGroup
	sem := make(chan int, *threads)

	starttime := time.Now()
	fmt.Printf("Started %d games: %s\n", *numberOfGames, starttime.Format("2006-01-02 15:04:05"))

	counter := 0
	for counter < *numberOfGames {
		wg.Add(1)
		sem <- 1

		go func() {
			defer wg.Done()
			var answer string
			var err error

			if *today {
				answer = TODAYSANSWER
			} else {
				// Pick a random answer
				answer, err = GetRandomWord(words)
				if err != nil {
					fmt.Printf("Error getting random answer: %v\n", err)
					return
				}
			}

			// List of words
			wordlist1 := []string{}
			wordlist1 = append(wordlist1, words...)

			// Create a game for method 1
			game1, err := NewGame(answer, wordlist1)
			if err != nil {
				log.Fatal(err)
			}

			// Play the game for method 1
			gameresult1 := Method1(game1, numberOfGuesses)
			resultsOutcome1.Gamesplayed++
			if gameresult1.Solved {
				resultsOutcome1.Gameswon++
				resultsOutcome1.Totalguesses += gameresult1.Guesses
			}

			//////////////////////////////////////////////////////////////////////////
			// List of words
			wordlist2 := []string{}
			wordlist2 = append(wordlist2, words...)

			// Create a game for method 2
			game2, err := NewGame(answer, wordlist2)
			if err != nil {
				log.Fatal(err)
			}

			// Play the game for method 2
			gameresult2 := Method2(game2, numberOfGuesses)
			resultsOutcome2.Gamesplayed++
			if gameresult2.Solved {
				resultsOutcome2.Gameswon++
				resultsOutcome2.Totalguesses += gameresult2.Guesses
			}

			//////////////////////////////////////////////////////////////////////////
			// List of words
			wordlist3 := []string{}
			wordlist3 = append(wordlist3, words...)

			// Create a game for method 3
			game3, err := NewGame(answer, wordlist3)
			if err != nil {
				log.Fatal(err)
			}

			// Play the game for method 2
			gameresult3 := Method2(game3, numberOfGuesses)
			resultsOutcome3.Gamesplayed++
			if gameresult3.Solved {
				resultsOutcome3.Gameswon++
				resultsOutcome3.Totalguesses += gameresult3.Guesses
			}

			<-sem
		}()

		counter++
	}

	// Wait for scans to be done
	wg.Wait()
	close(sem)

	endtime := time.Now()
	completedtime := endtime.Sub(starttime).Seconds()

	fmt.Printf("Completed: %ds\n", int64(completedtime))

	// Calculate the results for game 1
	fmt.Println("Method 1 - Random:")
	fmt.Printf("\tGames played          : %d\n", resultsOutcome1.Gamesplayed)
	percentwon1 := float64(resultsOutcome1.Gameswon) / float64(resultsOutcome1.Gamesplayed)
	fmt.Printf("\tGames won             : %d (%.2f%%)\n", resultsOutcome1.Gameswon, percentwon1)
	percentguess1 := float64(resultsOutcome1.Totalguesses) / float64(resultsOutcome1.Gameswon)
	fmt.Printf("\tGuesses for games won : %d (%.2f)\n", resultsOutcome1.Totalguesses, percentguess1)

	// Calculate the results for game 2
	fmt.Println("Method 2 - 15 Characters:")
	fmt.Printf("\tGames played          : %d\n", resultsOutcome2.Gamesplayed)
	percentwon2 := float64(resultsOutcome2.Gameswon) / float64(resultsOutcome2.Gamesplayed)
	fmt.Printf("\tGames won             : %d (%.2f%%)\n", resultsOutcome2.Gameswon, percentwon2)
	percentguess2 := float64(resultsOutcome2.Totalguesses) / float64(resultsOutcome2.Gameswon)
	fmt.Printf("\tGuesses for games won : %d (%.2f)\n", resultsOutcome2.Totalguesses, percentguess2)

	// Calculate the results for game 3
	fmt.Println("Method 3 - 4 Vowels:")
	fmt.Printf("\tGames played          : %d\n", resultsOutcome3.Gamesplayed)
	percentwon3 := float64(resultsOutcome3.Gameswon) / float64(resultsOutcome3.Gamesplayed)
	fmt.Printf("\tGames won             : %d (%.2f%%)\n", resultsOutcome3.Gameswon, percentwon3)
	percentguess3 := float64(resultsOutcome3.Totalguesses) / float64(resultsOutcome3.Gameswon)
	fmt.Printf("\tGuesses for games won : %d (%.2f)\n", resultsOutcome3.Totalguesses, percentguess3)
}
