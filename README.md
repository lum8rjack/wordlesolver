# wordlesolver

## Overview
I created this project to try and identify the most effective method for solving [Wordle ](https://www.nytimes.com/games/wordle/index.html). I came up with 3 different methods to test:

1. This method starts with a random word and randomly selects the remaining guesses based on valid/invalid characters identified.
2. This method starts with 3 random guesses all with different letters to remove 15 characters.
3. This method starts with a word that has atleast 4 vowels in it to reduce the possible options.

Thousands of games are played for each method to get statistics on how often the game was solved and the average number of guesses required.

```bash
Usage of ./wordlesolver.bin:
  -games int
        Number of games to play (default 100)
  -threads int
        Number of threads to use (default 25)
  -today
        Use today's answer for all games (default false)
  -wordlist string
        Which wordlist to use (default "dictionaries/wordle-03022023.txt")
```

## Results
Below are the results from a few different test runs.

```bash
./wordlesolver.bin -games 1000
Wordlesolver v1.0
Wordlist: dictionaries/wordle-03022023.txt (14855 words)
Started 1000 games
Completed: 397s
Method 1 - Random:
	Games played          : 1000
	Games won             : 371 (0.37%)
	Guesses for games won : 1909 (5.15)
Method 2 - 15 Characters:
	Games played          : 1000
	Games won             : 436 (0.44%)
	Guesses for games won : 2205 (5.06)
Method 3 - 4 Vowels:
	Games played          : 1000
	Games won             : 431 (0.43%)
	Guesses for games won : 2164 (5.02)


./wordlesolver.bin -games 10000
Wordlesolver v1.0
Wordlist: dictionaries/wordle-03022023.txt (14855 words)
Started 10000 games
Completed: 3822s
Method 1 - Random:
	Games played          : 10000
	Games won             : 3717 (0.37%)
	Guesses for games won : 19368 (5.21)
Method 2 - 15 Characters:
	Games played          : 10000
	Games won             : 4310 (0.43%)
	Guesses for games won : 21835 (5.07)
Method 3 - 4 Vowels:
	Games played          : 10000
	Games won             : 4312 (0.43%)
	Guesses for games won : 21854 (5.07)
```

## Today's Answer
You can also have every game use today's answer by passing the the "today" flag.

```bash
./wordlesolver.bin -today=true -games 1000
Wordlesolver v1.0
Using today's answer for all games
Wordlist: dictionaries/wordle-03022023.txt (14855 words)
Started 1000 games
Completed: 416s
Method 1 - Random:
        Games played          : 1000
        Games won             : 491 (0.49%)
        Guesses for games won : 2499 (5.09)
Method 2 - 15 Characters:
        Games played          : 1000
        Games won             : 526 (0.53%)
        Guesses for games won : 2590 (4.92)
Method 3 - 4 Vowels:
        Games played          : 1000
        Games won             : 500 (0.50%)
        Guesses for games won : 2450 (4.90)
```

## Improvements
The next step is to use Go's pprof profiling tool to identify areas of the code that could be improved to speed up the games.

## References
I pulled the dictionary from Wordle's source code, but you could also create your own list using the following word lists:
-  http://www.gwicks.net/textlists/english3.zip
- /usr/share/dict/words on linux machines

Then run the following command to get only 5 character words:
```bash
strings /usr/share/dict/words | grep -E '^[[:alpha:]]{5}$' > 5characterwords.txt
```

