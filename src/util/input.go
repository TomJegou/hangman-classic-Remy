package util

import "fmt"

//cette fonction permet de récupérer la lettre ou le mot entré dans le terminal

func Input() string {
	var letter string
	fmt.Scanln(&letter)
	return letter
}
