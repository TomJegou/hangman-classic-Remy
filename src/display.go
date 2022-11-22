package src

import (
	"bufio"
	"fmt"
	"os"
)

// Cette fonction a pour but d'afficher l'avancer du pendu après chaque erreur
func Display(attempt int) {
	colorIRed := "\033[0;91m"
	colorReset := "\033[0m"
	draw, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println((colorIRed), "The hangman can't open !", (colorReset))
	}
	Scanner := bufio.NewScanner(draw)
	Scanner.Split(bufio.ScanLines)
	var line []string
	for Scanner.Scan() {
		line = append(line, Scanner.Text())
	}
	Index := (attempt) * 8
	if attempt == 0 || attempt == 1 || attempt == 2 || attempt == 3 || attempt == 4 || attempt == 5 || attempt == 6 || attempt == 7 || attempt == 8 || attempt == 9 {
		for d := Index; d < Index+8; d++ {
			fmt.Println(line[d])
		}
	}
}
