package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	VOWELS = []rune{'a', 'e', 'i', 'o', 'u'}
)

type TotalResults struct {
	Version              string  `json:"version"`
	GamesPlayed          int     `json:"games_played"`
	GamesWon             int     `json:"games_won"`
	GamesWonPercentage   int     `json:"games_won_percentage"`
	AverageGuessesPerWin float64 `json:"average_guesses_per_win"`
	DictionaryLength     int     `json:"dictionary_length"`
	TimeToComplete       int     `json:"time_to_complete"`
	Threads              int     `json:"threads"`
}

type GameResult struct {
	Answer  string
	Guesses int
	Solved  bool
}

type AnswerToday struct {
	ID              int    `json:"id"`
	Solution        string `json:"solution"`
	PrintDate       string `json:"print_date"`
	DaysSinceLaunch int    `json:"days_since_launch"`
	Editor          string `json:"editor"`
}

func LoadWordlist(fn string) ([]string, error) {
	var words []string

	readFile, err := os.Open(fn)
	if err != nil {
		return words, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, strings.ToLower(fileScanner.Text()))
	}

	return words, nil
}

func GetRandomWord(options []string) (string, error) {
	wordlength := int64(len(options))
	if wordlength <= 0 {
		return "empty", nil
	}
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(wordlength))
	if err != nil {
		return "", err
	}

	pick := options[randomIndex.Int64()]
	return pick, nil
}

func GetXLetterVowels(options []string, numvowels int) ([]string, error) {
	var newlist []string

	counter := 0

	for _, word := range options {
		for _, v := range VOWELS {
			if strings.ContainsRune(word, v) {
				counter++
			}
		}

		if counter >= numvowels {
			newlist = append(newlist, word)
		}
		counter = 0
	}

	return newlist, nil
}

func GetUniqueLetterOptions(options []string) []string {
	var newlist []string
	add := true

	for _, word := range options {
		nm := make(map[rune]bool)
		for _, r := range word {
			if nm[r] {
				add = false
				break
			} else {
				nm[r] = true
			}
		}
		if add {
			newlist = append(newlist, word)
		}
		add = true
	}
	return newlist
}

func RemoveWord(word string, list []string) []string {
	if len(list) == 0 {
		return list
	}

	index := -1
	for i, element := range list {
		if strings.EqualFold(word, element) {
			//if word == element {
			index = i
			break
		}
	}

	if index == -1 {
		return list
	}

	newlist := append(list[:index], list[index+1:]...)
	return newlist
}

func RemoveDuplicateValues(runeSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}

	for _, entry := range runeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveIncorrectPosition(correct []rune, options []string) []string {
	// Make sure the list isn't empty
	if len(correct) == 0 {
		return options
	}

	// Check if all values are 0
	allzero := true
	for _, r := range correct {
		if r != 0 {
			allzero = false
			break
		}
	}

	if allzero {
		return options
	}

	// Setup the list we will return
	var itemsToRemove []string

	// Loop through correct positions
	for position, character := range correct {
		if character != 0 {
			// Loop through options
			var newrune []rune
			for _, word := range options {
				newrune = []rune(word)
				if character != newrune[position] {
					itemsToRemove = append(itemsToRemove, word)
				}
			}
		}
	}

	// Remove the items
	if len(itemsToRemove) > 0 {
		for _, itr := range itemsToRemove {
			options = RemoveWord(itr, options)
		}
	}

	return options
}

func RemoveWordsWithInvalidChars(invalid []rune, options []string) []string {
	if len(invalid) == 0 || len(options) == 0 {
		return options
	}

	newlist := options

	// Loop through invalid chars
	for _, invalidChar := range invalid {
		templist := newlist
		// Loop through options
		for _, op := range templist {
			if strings.ContainsRune(op, invalidChar) {
				newlist = RemoveWord(op, newlist)
			}
		}
	}

	return newlist
}

// Web request to get today's Wordle answer
func GetTodaysAnswer() (string, error) {
	var answer string
	// ex: https://www.nytimes.com/svc/wordle/v2/2023-03-06.json
	url := "https://www.nytimes.com/svc/wordle/v2/"
	dt := time.Now()
	newurl := fmt.Sprintf("%s%s.json", url, dt.Format("2006-01-02"))

	webclient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, newurl, nil)
	if err != nil {
		return answer, err
	}

	req.Header.Set("User-Agent", "wordlesolver")

	res, err := webclient.Do(req)
	if err != nil {
		return answer, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return answer, err
	}

	atoday := AnswerToday{}
	err = json.Unmarshal(body, &atoday)
	if err != nil {
		return answer, nil
	}

	answer = atoday.Solution

	return answer, nil
}
