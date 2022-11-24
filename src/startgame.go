package src

import "runtime"

// Cette fonction a pour but d'afficher des tirets à la place du mot à deviner

func Dash(word string) string {
	wordBytes := []byte(word)
	i := 0
	if runtime.GOOS == "windows" {
		i = 1
	}
	for index := 0; index < len(word)-i; index++ {
		wordBytes[index] = 95
	}
	return string(wordBytes)
}

//cette fonction sert à effectuer les affichages de chaque début de tour de boucle

func StartGame(word_dash string, used_letter string, used_word string, lettertmp string) string {
	letter := lettertmp
	letter = Accent(letter)
	letter = Convert(letter)
	return letter
}

// cette fonction convertie les lettres majuscules en lettres minuscules

func Convert(letter string) string {
	i := 0
	if runtime.GOOS == "windows" {
		i = 1
	}
	for index := 0; index < len(letter)-i; index++ {
		if letter[index] >= 65 && letter[index] <= 90 {
			letterByte := []byte(letter)
			letterByte[index] += 32
			letter = string(letterByte)
		}
	}
	return letter
}

//cette fonction permet de corriger des lettres avec accent en lettre classique pour empêcher les erreurs

func Accent(letter string) string {
	if len(letter) == 2 {
		if letter[1] == 136 || letter[1] == 137 || letter[1] == 138 || letter[1] == 139 || letter[1] == 168 || letter[1] == 169 || letter[1] == 170 || letter[1] == 171 {
			return "e"
		} else if letter[1] == 148 || letter[1] == 180 {
			return "o"
		} else if letter[1] == 128 || letter[1] == 160 {
			return "a"
		} else if letter[1] == 135 || letter[1] == 167 {
			return "c"
		} else if letter[1] == 153 || letter[1] == 185 {
			return "u"
		}
	}
	return letter
}
