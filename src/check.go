package src

// Cette fonction a pour but de vérifier si la lettre qu'on demande est dans le mot à deviner si c'est le cas il return vrai sinon return faux

func CheckLetterInWord(word string, letter string) bool {
	for _, value := range word {
		if string(value) == letter {
			return true
		}
	}
	return false
}

// cette fonction vérifie si la lettre entrée n'est pas déjà essayée et fausse

func CheckUsedLetter(letters string, letter string) bool {
	for _, value := range letters {
		if string(value) == letter {
			return false
		}
	}
	return true
}
