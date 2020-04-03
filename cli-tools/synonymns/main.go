package main

import (
	"bufio"
	"fmt"
	"go-blueprint-book/cli-tools/thesaurus"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed"+word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldnt find any synonyms for" + word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
