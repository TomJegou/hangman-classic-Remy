package src

import (
	"fmt"
	"os"

	"github.com/TomJegou/hangman-classic-Remy/src/util"
)

// Cette fonction a pour but d'afficher les règles du pendu

func DisplayIntroduction() {
	colorIRed := "\033[0;91m"
	colorBlue := "\033[0;34m"
	colorReset := "\033[0m"
	data, err := os.ReadFile("start.txt")
	if err != nil {
		fmt.Println((colorIRed), "The game menu can't open !", (colorReset))
		os.Exit(1)
	}
	util.Clear()
	fmt.Println((colorBlue), string(data), (colorReset))
}
