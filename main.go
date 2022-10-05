package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	findFirstWord()
	lastWord()
	afterWord("before")
	preciseCharacter("now")
	random()
}

/* Fonctions pour trouver le premier mot, il suffit de parcourir le fichier texte a l'aide d'une boucle et de la break a la première itération
ce qui vas me permettre de print uniquement le premier mot du fichier texte */

func findFirstWord() {
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		break
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
}

/* Fonctions pour trouver l'opposé du premier mot, ici le premier mot a trouvé était le tout premier du fichier texte donc on ici il faut
trouver le dernier mot car c'est l'opposé du premier. Pour ce faire on prend une boucle qui parcourt le fichier texte et on reiniatilise la
variable crée avec le mot auquel on se trouve donc a la fin de la boucle nous serons forcèment au dernier mot ensuite cela prend le lastword et ca return
au bon emplacement */

func lastWord() {
	var lastWord string
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lastWord = fileScanner.Text()
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
	fmt.Println(lastWord)
}

/* Fonctions afterWord, d'abord on met en paramètre un mot et on crée une variable mot ensuite on fait une boucle et des que le mot sera égal au paramètre
mot alors la variable after word sera égale au mot situé apres le mot trouvé
*/

func afterWord(mot string) {
	var afterWord string
	nmot := ""
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() == mot {
			fileScanner.Scan()
			afterWord = fileScanner.Text()
			i, _ := strconv.Atoi(afterWord)
			nmot = nthWord(i)
			break
		}
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
	fmt.Println(nmot)
}

func nthWord(n int) string {
	var nthWord string
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if n == 1 {
			nthWord = fileScanner.Text()
			break
		}
		n--
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
	return nthWord
}

/* Fonctions beforeWord, d'abord on prend en paramètre un mot et nous met le mot avant ce mot */

func beforeWord(mot string) string {
	var beforeWord string
	nmot := ""
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() == mot {
			beforeWord = nmot
			break
		}
		nmot = fileScanner.Text()
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
	return beforeWord
}

/* Fonctions preciseCharacter, qui prend en paramétre un mot et vas prendre le mot avant ce mot a l'aide de la fonction before word ensuite on va
prendre le 2ème caractère du mot et le diviser par le nombre mot total du fichier texte */

func preciseCharacter(mot string) {
	nmot := beforeWord(mot)
	r := []rune(nmot)
	char := r[1]
	char /= 38
	newMot := nthWord(int(char))
	fmt.Println(newMot)
}

/* Fonction qui genere une nombre aléatoire en prenant en parametre un min et un max et donc le nombre sera generé entre le min et le max */

func random() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(38)
	fmt.Println(i)
}
