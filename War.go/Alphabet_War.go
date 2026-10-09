package main

import "fmt"

func AlphabetWar(fight string) string {
	leftPower := map[rune]int{
		'w': 4,
		'p': 3,
		'b': 2,
		's': 1,
		't': 0,
	}
	rightPower := map[rune]int{
		'm': 4,
		'q': 3,
		'd': 2,
		'z': 1,
		'j': 0,
	}
	convert := map[rune]rune{
		'm': 'w', 'w': 'm',
		'q': 'p', 'p': 'q',
		'd': 'b', 'b': 'd',
		'z': 's', 's': 'z',
	}
	word := []rune(fight)
	result := make([]rune, len(word))
	copy(result, word)
	for i := 0; i < len(word); i++ {
		if word[i] == 't' || word[i] == 'j' {
			if i+1 < len(word) && word[i+1] != 't' && word[i+1] != 'j' {
				_, existLeft := leftPower[word[i+1]]
				_, existRight := rightPower[word[i+1]]
				if existLeft && word[i] == 'j' && (i+2 >= len(word) || word[i+2] != 'j') {
					result[i+1] = convert[word[i+1]]
				}
				if existRight && word[i] == 't' && (i+2 >= len(word) || word[i+2] != 'j') {
					result[i+1] = convert[word[i+1]]
				}
			}
			if i-1 >= 0 && word[i-1] != 't' && word[i-1] != 'j' {
				_, existRightmines := rightPower[word[i-1]]
				_, existleftmineu := leftPower[word[i-1]]
				if existRightmines && word[i] == 't' && (i-2 < 0 || word[i-2] != 't') {
					result[i-1] = convert[word[i-1]]
				}
				if existleftmineu && word[i] == 'j' && (i-2 < 0 || word[i-2] != 'j') {
					result[i-1] = convert[word[i-1]]
				}
			}
		}
	}

	leftScore := 0
	rightScore := 0
	for _, v := range result {
		if power, okay := leftPower[v]; okay {
			leftScore += power
		}
		if power, okay := rightPower[v]; okay {
			rightScore += power
		}
	}

	if rightScore > leftScore {
		return "Right side wins!"
	}
	if leftScore > rightScore {
		return "Left side wins!"
	}
	return "Let's fight again!"
}

func main() {
	fmt.Println("=== Alphabet War ===")
	fmt.Println("This program takes a string of lowercase letters and decides which side wins.")
	fmt.Println("Left side:  w(4) p(3) b(2) s(1) t(priest)")
	fmt.Println("Right side: m(4) q(3) d(2) z(1) j(priest)")
	fmt.Println("Priests (t and j) convert an adjacent enemy letter into its friendly equivalent.")
	fmt.Println()

	var input string
	fmt.Print("Enter the fight string: ")
	fmt.Scan(&input)

	result := AlphabetWar(input)
	fmt.Println("Result:", result)
}
