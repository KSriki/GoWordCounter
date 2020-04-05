package main

import (
	"bufio" //read text
	"fmt"   //print formatted ouput
	"io"    // io.Reader
	"os"    // os resources
)

func main() {
	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// A scanner is used to read text from the a Reader (such as files)
	scanner := bufio.NewScanner(r)
	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)
	// Defining a counter
	wc := 0
	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	// Return the total
	return
}
