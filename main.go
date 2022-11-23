package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/asmaloney/naming-language-gen/naming"
)

func main() {
	var (
		help = flag.Bool("help", false, "Displays helpful usage info")
	)

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())

	lang := naming.RandomLanguage(true, true)

	// generate some words based on our new awesome language
	p := &naming.WordParams{
		MinSyllables: 1,
		MaxSyllables: naming.RandomRange(1, 5),
		Structure:    naming.DefaultSyllableStructures,
	}

	group := "words"

	for i := 0; i < 25; i++ {
		lang.GetWord(p, group)
	}

	wordList := strings.Join(lang.Words.General[group], ", ")
	fmt.Printf("[%v]: %v\n", group, wordList)

	// generate some names
	p = &naming.WordParams{
		MinSyllables: 2,
		MaxSyllables: naming.RandomRange(2, 7),
		Structure:    naming.DefaultSyllableStructures,
	}

	params := naming.NameParams{
		MinLength:  naming.RandomRange(3, 5),
		MaxLength:  naming.RandomRange(6, 20),
		WordParams: p,
		Joiners:    "  -",
		Group:      "words",
	}

	for i := 0; i < 20; i++ {
		lang.MakeName(&params)
	}

	nameList := strings.Join(lang.Words.Names, ", ")
	fmt.Printf("[%v]: %v\n", "names", nameList)

	lang.Describe()
}
