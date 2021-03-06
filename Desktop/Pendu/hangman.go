package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/* print a string with space between every characters */
func spaceOut(s string) string {
	if len(s) < 1 {
		panic("Input string must have at least lenght of 1. Lenght: " + strconv.Itoa(len(s)))
	}
	if len(s) == 1 {
		return s
	}

	finalSentence := string(s[0])
	for i := 1; i < len(s); i++ {
		finalSentence += " " + string(s[i])
	}
	return finalSentence
}

/* Replaces characters with an underscore if their index in hidden is true */
func hideLetters(s string, hidden []bool) string {
	if len(s) < 1 {
		panic("Input string must have at least lenght of 1. Lenght: " + strconv.Itoa(len(s)))
	}

	finalString := ""
	if hidden[0] {
		finalString += "_"
	} else {
		finalString += string(s[0])
	}
	for i := 1; i < len(s)-1; i++ {
		if hidden[i] {
			finalString += " _"
		} else {
			finalString += " " + string(s[i])
		}
	}
	return finalString
}

/* Returns the index of s if array contains s. Returns -1 if s is not contained within array. */
func alreadyGuessed(array []string, s string) int {
	for index := range array {
		if array[index] == s {
			return index
		}
	}
	return -1
}

/* Verification if the string s is only composed by letter and not number...*/
func error(s string) bool {
	bol := true
	for _, v := range s {
		if !(v >= rune(97) && v <= rune(122)) && !(v >= rune(65) && v <= rune(90)) {
			bol = false
		}
	}
	return bol
}

func main() {
	///import the different words.txt
	if len(os.Args) == 2 {
		arr, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Print(err)
		}
		separation := "\n"
		myString := string(arr)
		wordsman1 := []rune(myString)
		var wordsman2 string
		for lama1 := 0; lama1 < len(wordsman1); lama1++ {
			if wordsman1[lama1] == 130 || (wordsman1[lama1] >= 232 && wordsman1[lama1] <= 235) {
				wordsman2 += "e"
			} else {
				wordsman2 += string(wordsman1[lama1])
			}
		}
		wordsman := strings.Split(wordsman2, separation)
		/// import hangman.txt
		death2, err := ioutil.ReadFile("hangman.txt")
		if err != nil {
			fmt.Print(err)
		}
		separation2 := "\n"
		myString2 := string(death2)
		death := strings.Split(myString2, separation2)
		/// game starting

		attempt := 57
		countDeath := 0
		fmt.Println("Good Luck, you have 10 attempts.")
		/// import random word and random letter
		rand.Seed(time.Now().UnixNano())
		word := strings.ToUpper(wordsman[rand.Intn(len(wordsman))])
		letterFind := strings.ToUpper(string(word[rand.Intn(len(word)-1)]))
		/// information about random letter
		count2 := 1
		position := 0
		tableau := []rune(letterFind)
		tableau2 := []rune(word)
		for i := 0; i < len(tableau2); i++ {
			if tableau2[i] == tableau[0] {
				position = count2
			} else {
				count2++
			}
		}
		/// printing about the letter
		fmt.Printf("\n")
		fmt.Printf("You have a %v in your word at the ", letterFind)
		fmt.Print(position)
		fmt.Println("eme position")
		///variable declaration
		hiddenSpaces := len(word) - 1
		hiddenLetters := make([]bool, hiddenSpaces)
		for i := range hiddenLetters {
			hiddenLetters[i] = true
		}
		var previousGuesses []string
		/// verification if it's win
		reader := bufio.NewReader(os.Stdin)
		for {
			if hiddenSpaces == 0 {
				fmt.Println(spaceOut(word))
				fmt.Printf("\n")
				fmt.Println("Congrats!")
				fmt.Printf("\n")
				return
			}
			///printing and verification about the letter enter by the player
			fmt.Println(hideLetters(word, hiddenLetters))
			fmt.Printf("\n")
			fmt.Print("Choose: ")
			text, _ := reader.ReadString('\n')
			text = strings.ToUpper(strings.TrimSpace(text))
			if !(error(text)) {
				fmt.Println("You must enter a letter of the Latin alphabet.")
			}
			/// verification about a entire word enter by the player
			if len(text) > 1 {
				if text == "STOP" {
					fmt.Println("BOUGE!")
					return
				}
				if text == word {
					fmt.Printf("\n")
					fmt.Println("Congrats!!  Bon j'avoue, celui-l??, il ??tait simple...")
					return
				} else if countDeath >= 8 {
					fmt.Println("Wooooops! Wrong word!")
					fmt.Println("You have no more attempts :/")
					fmt.Println(death[63])
					fmt.Println(death[64])
					fmt.Println(death[65])
					fmt.Println(death[66])
					fmt.Println(death[67])
					fmt.Println(death[68])
					fmt.Println(death[69])
					fmt.Println("The word was :")
					fmt.Println(word)

					break
				} else {
					fmt.Println("Wooooops! Wrong word!")
					fmt.Println("You just lose two attempts x/")
					countDeath += 2
					attempt -= 2
					continue
				}
			} else if len(text) < 1 {
				continue
			}
			///letter already guessed
			if alreadyGuessed(previousGuesses, text) != -1 {
				fmt.Println("You have already guessed \"" + text + "\"")
				continue
			}
			/// verification if it's a good letter or not
			correctGuess := false
			for index, value := range word {
				if string(text[0]) == string(value) {
					if !correctGuess {
						correctGuess = true
						previousGuesses = append(previousGuesses, string(text[0]))
					}
					hiddenLetters[index] = false
					hiddenSpaces--
				}
			}
			///print jos?? the hangman
			if !correctGuess {
				fmt.Println("Not present in the word, " + string(rune(attempt)) + " attempts remaining")
				if countDeath == 0 {
					fmt.Println(death[0])
					fmt.Println(death[1])
					fmt.Println(death[2])
					fmt.Println(death[3])
					fmt.Println(death[4])
					fmt.Println(death[5])
					fmt.Println(death[6])
				}
				if countDeath == 1 {
					fmt.Println(death[7])
					fmt.Println(death[8])
					fmt.Println(death[9])
					fmt.Println(death[10])
					fmt.Println(death[11])
					fmt.Println(death[12])
					fmt.Println(death[13])
				}
				if countDeath == 2 {
					fmt.Println(death[14])
					fmt.Println(death[15])
					fmt.Println(death[16])
					fmt.Println(death[17])
					fmt.Println(death[18])
					fmt.Println(death[19])
					fmt.Println(death[20])
				}
				if countDeath == 3 {
					fmt.Println(death[21])
					fmt.Println(death[22])
					fmt.Println(death[23])
					fmt.Println(death[24])
					fmt.Println(death[25])
					fmt.Println(death[26])
					fmt.Println(death[27])
				}
				if countDeath == 4 {
					fmt.Println(death[28])
					fmt.Println(death[29])
					fmt.Println(death[30])
					fmt.Println(death[31])
					fmt.Println(death[32])
					fmt.Println(death[33])
					fmt.Println(death[34])
				}
				if countDeath == 5 {
					fmt.Println(death[35])
					fmt.Println(death[36])
					fmt.Println(death[37])
					fmt.Println(death[38])
					fmt.Println(death[39])
					fmt.Println(death[40])
					fmt.Println(death[41])
				}
				if countDeath == 6 {
					fmt.Println(death[42])
					fmt.Println(death[43])
					fmt.Println(death[44])
					fmt.Println(death[45])
					fmt.Println(death[46])
					fmt.Println(death[47])
					fmt.Println(death[48])
				}
				if countDeath == 7 {
					fmt.Println(death[49])
					fmt.Println(death[50])
					fmt.Println(death[51])
					fmt.Println(death[52])
					fmt.Println(death[53])
					fmt.Println(death[54])
					fmt.Println(death[55])
				}
				if countDeath == 8 {
					fmt.Println(death[56])
					fmt.Println(death[57])
					fmt.Println(death[58])
					fmt.Println(death[59])
					fmt.Println(death[60])
					fmt.Println(death[61])
					fmt.Println(death[62])
				}
				if countDeath == 9 {
					fmt.Println(death[63])
					fmt.Println(death[64])
					fmt.Println(death[65])
					fmt.Println(death[66])
					fmt.Println(death[67])
					fmt.Println(death[68])
					fmt.Println(death[69])
					fmt.Println("The word was :")
					fmt.Println(word)

					break
				}
				attempt--
				countDeath++
			}
		}
	} else {
		/// error printing about an argument missing
		fmt.Printf("\n")
		fmt.Print("If you want play to hangman, you must write words.txt or words2.txt or words3.txt \n")
		fmt.Printf("\n")
	}
}
