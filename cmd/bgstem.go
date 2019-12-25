package main

import (
	"log"
	"os"

	"github.com/aquilax/go-stemmer"
)

func main() {
	rules, err := stemmer.LoadRules(os.Args[1], 1)
	if err != nil {
		log.Fatal(err)
	}
	println(stemmer.Stem(os.Args[2], rules))
}
