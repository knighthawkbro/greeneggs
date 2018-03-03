package main

import (
	"bufio"
	"fmt"
	"greeneggs/array"
	"greeneggs/list"
	"log"
	"os"
	"regexp"
	"strings"
)

type set interface {
	Add(T interface{}) error // changed to error, it basically is used like a bool
	Remove() interface{}
	Get() interface{}
	Size() int
	String() string
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	arr := array.New()
	driver(arr)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning seuss function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	arr = array.New()
	seuss(arr)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	words := list.New()
	driver(words)
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning seuss function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	fmt.Println("The suess method with an list...")
	seuss(words)
}

func driver(words set) {
	fruits := []string{"orange", "grape", "kiwi", "coconut", "lime"}
	for _, fruit := range fruits {
		if err := words.Add(fruit); err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("Here's what's in our set:", words)

	fmt.Println("Let's add a banana")
	if err := words.Add("banana"); err != nil {
		log.Println(err)
	} else {
		fmt.Println("Operation Successful!")
	}
	fmt.Println("Our set looks like this", words)

	fmt.Println("Let's try to add another orange")

	if err := words.Add("orange"); err != nil {
		log.Println(err)
	} else {
		fmt.Println("Operation Successful!")
	}
	fmt.Println("Our set looks like this:", words)

	fmt.Println("Trying to add null to our set")
	if err := words.Add(nil); err != nil {
		log.Println(err)
	}
	fmt.Println("")
}

func seuss(words set) {
	file, err := os.Open("greenEggs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = reg.ReplaceAllString(word, "")
			words.Add(strings.ToLower(word))
		}
	}
	fmt.Println(words)
}
