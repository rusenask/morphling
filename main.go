package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var morphs = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

func main() {
	// using nano seconds as a seed to create a random number
	rand.Seed(time.Now().UTC().UnixNano())
	// bufio.Scanner object reads os.Stdin (actually takes io.Reader as an input
	// source so you can use it in many ways)
	s := bufio.NewScanner(os.Stdin)
	// bufio default split function is line by line so taking only first line
	// calling Scan function tells scanner to read the next block of bytes (in
	// this case next line) from the input and returns a bool value indicating
	// whether it found anything or not. This is how we are able to use it as
	// the condition for the for loop. While there is still something - scanner
	// returns true and the loop is executed, when it reaches end of line/input
	// it returns false and the loop stops.
	for s.Scan() {
		// selects random item from morphs slice
		t := morphs[rand.Intn(len(morphs))]
		// bytes from input are stored in Bytes method of the scanner and the
		// Text method converts []byte slice into a string.
		// with the random number and current text we use strings.Replace method to
		// insert the original word where the otherWord string appears. using
		// fmt.Println to print the output to os.Stdout
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
