package src

import (
	"fmt"
	"os"

	"github.com/TomJegou/hangman-classic-Remy/src/util"
)

// cette fonction renvoie le résultat du programme

func Result(attempt int, word string, result string, nameDB string) {
	if attempt <= 0 {
		fmt.Printf("\033[31mSorry, the word was %v\033[0m \n", word)
	} else {
		fmt.Println(result)
		points := Points(attempt, nameDB)
		fmt.Printf("\033[32mCongrats ! You have %v points\033[0m \n", points)
		fmt.Print("Do you want to add a word ? \033[32mYes [1]\033[0m \033[31mNo [other key]\033[0m ")
		if util.Input() == "1" {
			fmt.Printf("Add a word in %v : ", nameDB)
			AddWord(nameDB)
		}
	}
}

//cette fonction permet d'ajouter un mot à la base de donnée du niveau du joueur

func AddWord(nameDB string) {
	data, err := os.ReadFile(nameDB)
	if err != nil {
		fmt.Println("The file can't be read")
	}
	word := string(data) + util.Input() + "\n"
	wordByte := []byte(word)
	os.WriteFile(nameDB, wordByte, 0644)
}

// cette fonction sert à compter les points du joueur selon son nombre d'essai restant et le niveau choisi

func Points(attempt int, nameDB string) int {
	if nameDB == "hard" {
		attempt *= 3
	} else if nameDB == "medium" {
		attempt *= 2
	}
	return attempt
}
