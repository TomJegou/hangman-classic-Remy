package src

//cette fonction sert a modifier le tiret par la lettre valide

func Replace(word string, wordDash string, letter string) string {
	wordDashByte := []byte(wordDash)
	letterByte := []byte(letter)
	for index := 0; index < len(word); index++ {
		if word[index] == letter[0] {
			wordDashByte[index] = letterByte[0]
		}
	}
	return string(wordDashByte)
}
