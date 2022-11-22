package src

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/TomJegou/hangman-classic-Remy/src/util"
)

//cette fonction prend un mot au hasard dans le fichier words

func GetWord(mot string) string {
	rand.Seed(time.Now().UnixNano())
	difficulty := mot //ChooseDifficulty()
	var data []byte
	var err error
	var nameDB string
	if difficulty == "easy" {
		nameDB = "easy.txt"
	} else if difficulty == "medium" {
		nameDB = "medium.txt"
	} else if difficulty == "hard" {
		nameDB = "hard.txt"
	} else {
		fmt.Println("\033[31mInvalid input !\033[0m")
	}
	data, err = os.ReadFile("db/" + nameDB)
	if err != nil {
		fmt.Println("The file can't be read")
		os.Exit(1)
	}
	array := strings.SplitAfter(string(data), "\n")
	index := rand.Intn(len(array) - 1)
	word := make([]byte, len(array[index])-1)
	for index, value := range array[index] {
		if value != '\n' {
			word[index] += byte(value)
		}
	}
	return string(word) //, nameDB
}

//cette fonction permet au joueur de choisir la difficulté de jeu qu'il veut

func ChooseDifficulty() string {
	fmt.Println("Choose your difficulty : \033[32mEasy [1]\033[0m \033[33mMedium [2]\033[0m \033[31mHard [3]\033[0m")
	difficulty := util.Input()
	return difficulty
}
