package main

import (
	"testing"
)

// TestLoadWordlist
func TestLoadWordlist(t *testing.T) {
	var tests = []struct {
		name     string
		location string
		length   int
	}{
		{name: "all", location: "dictionaries/wordle-03022023.txt", length: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wl, err := LoadWordlist(tt.location)
			if err != nil {
				t.Errorf("could not load wordlist %s - %v", tt.location, err)
			}

			// Loop through and make sure they are all 5 letters
			for i, w := range wl {
				if len(w) != tt.length {
					t.Errorf("invalid word %s at position %d", w, i)
				}
			}
		})
	}
}

// TestGetRandomWord
func TestGetRandomWord(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
	}{
		{name: "Empty list", input: []string{}},
		{name: "Select word", input: []string{"seven", "eight", "three", "emtpy"}},
	}

	// Make sure it grabs from our list
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := GetRandomWord(tt.input)
			if err != nil {
				t.Error(err)
			}
			if len(ans) != 5 {
				t.Errorf("got %s", ans)
			}
		})
	}
}

// TestGetRandomWord2
// Actually pulls from dictionary to make sure the random
// number generator is working correctly
func TestGetRandomWord2(t *testing.T) {
	wordlist, err := LoadWordlist("dictionaries/wordle-03022023.txt")
	if err != nil {
		t.Errorf("could not load wordlist: %v", err)
	}

	var tests = []struct {
		name   string
		input  []string
		checks int
	}{
		{name: "Not twice", input: wordlist, checks: 10},
	}

	// Make sure it grabs from our list
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run multiple times to confirm the words are not the same
			counter := 0
			var answers []string
			for counter < tt.checks {
				ans, err := GetRandomWord(tt.input)
				if err != nil {
					t.Error(err)
				}
				answers = append(answers, ans)
				counter++
			}

			// Check all the responses
			visited := make(map[string]bool, 0)
			for _, a := range answers {
				if visited[a] == true {
					t.Errorf("duplicate %s", a)
				} else {
					visited[a] = true
				}
			}
		})
	}
}

// TestGetXLetterVowels
func TestGetXLetterVowels(t *testing.T) {
	var tests = []struct {
		name         string
		inputOptions []string
		numvowels    int
		want         int
	}{
		{name: "Two Vowels", inputOptions: []string{"raise", "cloud", "found", "audio", "later"}, numvowels: 2, want: 5},
		{name: "Three Vowels", inputOptions: []string{"raise", "cloud", "found", "audio", "later"}, numvowels: 3, want: 2},
		{name: "Four Vowels", inputOptions: []string{"raise", "cloud", "found", "audio", "later"}, numvowels: 4, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := GetXLetterVowels(tt.inputOptions, tt.numvowels)
			if len(ans) != tt.want {
				t.Errorf("got %d, want %d", len(ans), tt.want)
			}
		})
	}
}

// TestGetUniqueLetterOptions
func TestGetUniqueLetterOptions(t *testing.T) {
	var tests = []struct {
		name         string
		inputOptions []string
		want         int
	}{
		{name: "All unique", inputOptions: []string{"raise", "cloud", "found", "audio", "later"}, want: 5},
		{name: "Four", inputOptions: []string{"raise", "cloud", "found", "audio", "malls"}, want: 4},
		{name: "Three", inputOptions: []string{"raise", "cloud", "found", "salts", "malls"}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := GetUniqueLetterOptions(tt.inputOptions)
			if len(ans) != tt.want {
				t.Errorf("got %d, want %d", len(ans), tt.want)
			}
		})
	}
}

// TestRemoveWord
func TestRemoveWord(t *testing.T) {
	var tests = []struct {
		name         string
		inputWord    string
		inputOptions []string
		want         int
	}{
		{name: "Match", inputWord: "three", inputOptions: []string{"one", "two", "three", "four", "five"}, want: 4},
		{name: "No match", inputWord: "six", inputOptions: []string{"one", "two", "three", "four", "five"}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := RemoveWord(tt.inputWord, tt.inputOptions)
			if len(ans) != tt.want {
				t.Errorf("got %d, want %d", len(ans), tt.want)
			}
		})
	}
}

// TestRemoveDuplicateValues
func TestRemoveDuplicateValues(t *testing.T) {
	var tests = []struct {
		name  string
		input []rune
		want  int
	}{
		{name: "No Duplicates", input: []rune("abcdefg"), want: 7},
		{name: "One Duplicate", input: []rune("abcdefc"), want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := RemoveDuplicateValues(tt.input)
			if len(ans) != tt.want {
				t.Errorf("got %d, want %d", len(ans), tt.want)
			}
		})
	}
}

// TestRemoveIncorrectPosition
func TestRemoveIncorrectPosition(t *testing.T) {
	var tests = []struct {
		name         string
		inputCorrect []rune
		inputOptions []string
		want         int
	}{
		{name: "test1", inputCorrect: []rune{0, 0, 0, 0, 0}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 5},
		{name: "test2", inputCorrect: []rune{0, 0, 0, 105, 0}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 2},   // 105=i
		{name: "test3", inputCorrect: []rune{122, 121, 0, 0, 0}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 4}, // 122=z, 121=y
		{name: "test4", inputCorrect: []rune{0, 0, 0, 97, 0}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 1},    // 97=a
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := RemoveIncorrectPosition(tt.inputCorrect, tt.inputOptions)
			if len(ans) != tt.want {
				t.Errorf("got %d, want %d\t%v", len(ans), tt.want, ans)
			}
		})
	}
}

// TestRemoveWordsWithInvalidChars
func TestRemoveWordsWithInvalidChars(t *testing.T) {
	var tests = []struct {
		name         string
		invalid      []rune
		inputOptions []string
		want         int
	}{
		{name: "empty", invalid: []rune{}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 5},
		{name: "none", invalid: []rune{'b', 'd', 'x'}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 5},
		{name: "two", invalid: []rune{'b', 'd', 'i'}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 3},
		{name: "three", invalid: []rune{'b', 'm', 'i'}, inputOptions: []string{"zunis", "zygal", "zygon", "zymes", "zymic"}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := RemoveWordsWithInvalidChars(tt.invalid, tt.inputOptions)
			if len(ans) != tt.want {
				t.Errorf("got %d, want %d\t%v", len(ans), tt.want, ans)
			}
		})
	}
}

// TestGetTodaysAnswer
func TestGetTodaysAnswer(t *testing.T) {
	var tests = []struct {
		name string
	}{
		{name: "check valid answer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := GetTodaysAnswer()
			if err != nil {
				t.Errorf("got error %s", err)
			}
			if len(ans) != 5 {
				t.Errorf("got invalid answer %s", ans)
			}
		})
	}
}
