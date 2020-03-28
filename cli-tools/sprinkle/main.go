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

var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets " + otherWord,
	otherWord + "hq",
}

func main() {
	//fmt.Println(time.Now().UTC().UnixNano())
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		phrase := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(phrase, otherWord, s.Text(), -1))

	}

}
