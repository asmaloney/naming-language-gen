package naming

import (
	"fmt"
	"strings"
	"testing"
)

var wordTests = []struct {
	group      string
	structures structureList
}{
	{"group1", DefaultSyllableStructures},
	{"group2", structureList{"S?CVC?"}},
	{"group3", structureList{"CVV?C", "CL?VF"}},
}

func TestMakeWord(t *testing.T) {
	lang := OrthoLanguage()

	lang.ApplyMorph = true
	lang.Phonemes["S"] = phonemeSSets["s ʃ f"]
	lang.SyllableRestrictions = restrictionSets["Hard clusters"]

	const minSyllables = 1
	const maxSyllables = 1

	for _, val := range wordTests {
		list := []string{}

		p := &WordParams{
			minSyllables,
			maxSyllables,
			val.structures,
		}

		for i := 0; i < 20; i++ {
			list = append(list, lang.makeWord(p, val.group))
		}

		fmt.Printf("Random words [%v]: %v\n", val.group, list)
	}
}

func TestGetWord(t *testing.T) {
	lang := OrthoLanguage()

	lang.ApplyMorph = true
	lang.Phonemes["C"] = consonantSets["English-ish"]
	lang.Phonemes["S"] = phonemeSSets["s ʃ f"]
	lang.SyllableRestrictions = restrictionSets["Hard clusters"]

	const minSyllables = 1
	const maxSyllables = 2

	for _, val := range wordTests {
		p := &WordParams{
			minSyllables,
			maxSyllables,
			val.structures,
		}

		for i := 0; i < 20; i++ {
			lang.GetWord(p, val.group)
		}

		list := strings.Join(lang.Words.General[val.group], ", ")

		fmt.Printf("Get words [%v]: %v\n", val.group, list)
	}
}

func TestRandomLang(t *testing.T) {
	lang := RandomLanguage(false, true)

	const minSyllables = 1
	const maxSyllables = 4

	for _, val := range wordTests {
		p := &WordParams{
			minSyllables,
			maxSyllables,
			val.structures,
		}

		for i := 0; i < 20; i++ {
			lang.GetWord(p, val.group)
		}

		list := strings.Join(lang.Words.General[val.group], ", ")

		fmt.Printf("Random lang words [%v]: %v\n", val.group, list)
	}
}
