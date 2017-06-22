// randomWord takes count as an argument and generates the random words.
// Input: A text file having one single word at each line. User should give only one number as an argument.
//        Bydefault it prints one word. Output: Examples are as below,

// Example 1) $ ./randomWord
// mauls

// Example 2) $ ./randomWord 3
// therology
// effuso
// spectrohelioscopic

// $ ./randomWord 3 5
// Please give Valid Input.

// Usage: ./randomWord

// Example: ./randomWord 2
// Output: apple car
//==========================================================================

package main
import (
	
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
)

// ReadFile takes the path as string and reads the file. It returns the array of lines.
func ReadFile(path string) []string {

	f, err := os.Open(path)
	if err != nil { panic(err) }
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}

// GenRandWords takes filepath and length as argument. It generates random words and returns the array.
func GenRandWords(filePath string, ln int) []string {

	// Read txt file.
	wordDataSet := ReadFile(filePath)

	// Seed the random number genrator, Just not to have same sequence of random numbers everytime.
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	var words []string

	if ln == 1 {
		words = append(words, wordDataSet[rand.Intn(len(wordDataSet))])
		return words
	}
	
	for i := 0; i <= ln; i++ {
		words = append(words, wordDataSet[rand.Intn(len(wordDataSet))])
	}
	return words
}

//==========================================================================
func main(){

	filePath := "/home/yogesh/words_alpha.txt";
	ln := len(os.Args) 
	usage := "\nUsage:\n" + os.Args[0] + " <Number to print words>\n\nExample: " + os.Args[0] + " 2\nOutput:\napple\ncar\n" 

	if ln > 2 {
		fmt.Println("Please give Valid Input.\n", usage)
		os.Exit(0)
	} 

	// Generate Random Words from user input
	words:=	GenRandWords(filePath, ln)

	// Print the Random Words
	for _, word := range words {
		fmt.Println(word)
	}
}
