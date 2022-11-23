package src

import (
	"fmt"
)

//cette fonction rassemble toutes les fonctions de création du hangman et le génère

func Hangman(inputChan <-chan string, responseChan chan<- string, levelChan <-chan string, attemptChan chan<- int, wordChan chan<- string, quitChan <-chan bool) {
	for {
		select {
		case <-quitChan:
			return
		default:
			fmt.Println("Hangman activated")
			var word string
			for level := range levelChan {
				if level != "" {
					word = GetWord(level)
					break
				}
			}
			fmt.Println(word)
			wordDash := Dash(word)
			responseChan <- wordDash
			for content := range inputChan {
				if content != "" && content == "b0c9713aa009f4fcf39920d0d7eda80714b0c44ff2f98205278be112c755ca45e5386cbe7a9fca360ad22f06e45f80a8b8f23838725d15f889e202f5cea26359" {
					break
				}
			}
			var attempt int
			usedLetters := ""
			usedWords := ""
			attempt = 10
			attemptChan <- attempt
			fmt.Println("bienvenu dans la boucle")
			for content := range inputChan {
				select {
				case <-quitChan:
					return
				default:
					fmt.Println("Début de la boucle de hang")
					letter := StartGame(wordDash, usedLetters, usedWords, content)
					if !CheckUsedLetter(usedLetters, letter) {
						if letter != "" {
							fmt.Println("\033[31mYou already try it once !\033[0m")
						}
					} else if len(letter) > 1 {
						if letter == word {
							responseChan <- "50536101b1c465eafbecc8fca26eeb18a2ac8a2f83570bade315c5a112363cdfd820acad2ab234f91d43f0db8fed0cec400a1109ad8f99c21b5b74f59e8bb00d"
							attemptChan <- attempt
							break
						} else {
							attempt--
							Display(attempt)
							if attempt == 0 {
								responseChan <- "889ce65f137b3b9aa1005f417d7972c948b8bb6360cbdd4118cb05a29d37905744fc0dbc3d17c1de02689d837bfea5bb8114a994f9c1a53dddb993139ab2974c"
								attemptChan <- attempt
								wordChan <- word
								break
							}
						}
					} else if CheckLetterInWord(word, letter) {
						wordDash = Replace(word, wordDash, letter)
						if wordDash == word {
							responseChan <- "50536101b1c465eafbecc8fca26eeb18a2ac8a2f83570bade315c5a112363cdfd820acad2ab234f91d43f0db8fed0cec400a1109ad8f99c21b5b74f59e8bb00d"
							attemptChan <- attempt
							break
						}
					} else if letter == "" {
						responseChan <- wordDash
						attemptChan <- attempt
						continue
					} else {
						usedLetters += letter
						attempt--
						Display(attempt)
						if attempt == 0 {
							responseChan <- "889ce65f137b3b9aa1005f417d7972c948b8bb6360cbdd4118cb05a29d37905744fc0dbc3d17c1de02689d837bfea5bb8114a994f9c1a53dddb993139ab2974c"
							attemptChan <- attempt
							wordChan <- word
							break
						}
					}
					responseChan <- wordDash
					attemptChan <- attempt
					fmt.Println("Fin de l'envoie")
				}
			}
		}
	}
}

// Cette fonction a pour but d'afficher des tirets à la place du mot à deviner

func Dash(word string) string {
	wordBytes := []byte(word)
	for index := 0; index < len(word); index++ {
		wordBytes[index] = 95
	}
	return string(wordBytes)
}
