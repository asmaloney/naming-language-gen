package naming

import (
	"math/rand"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type NameParams struct {
	MinLength  int
	MaxLength  int
	WordParams *WordParams
	Joiners    string
	Group      string
}

func (lang *Language) MakeName(params *NameParams) (name string) {
	if params.MinLength <= 0 {
		params.MinLength = 5
	}

	if params.MaxLength < params.MinLength {
		params.MaxLength = params.MinLength
	} else if params.MaxLength <= 0 {
		params.MaxLength = 12
	}

	if params.WordParams.MinSyllables <= 0 {
		params.WordParams.MinSyllables = 1
	}

	if params.WordParams.MaxSyllables < params.WordParams.MinSyllables {
		params.WordParams.MaxSyllables = params.WordParams.MinSyllables
	} else if params.MaxLength <= 0 {
		params.WordParams.MaxSyllables = 2
	}

	joinersLen := len(params.Joiners)

	titleCase := cases.Title(language.Und)

	for {
		if rand.Float32() < 0.5 {
			name = titleCase.String(lang.GetWord(params.WordParams, params.Group))
		} else {
			g := ""
			if rand.Float32() < 0.6 {
				g = params.Group
			}
			w1 := titleCase.String(lang.GetWord(params.WordParams, g))
			g = ""
			if rand.Float32() < 0.6 {
				g = params.Group
			}
			w2 := titleCase.String(lang.GetWord(params.WordParams, g))
			if w1 == w2 {
				continue
			}

			if joinersLen > 0 {
				join := RandomRuneFromString(params.Joiners)

				if rand.Float32() > 0.5 {
					name = strings.Join([]string{w1, w2}, join)
				} else {
					name = strings.Join([]string{w1, lang.Words.Genitive, w2}, join)
				}
			}
		}

		if joinersLen > 0 {
			join := RandomRuneFromString(params.Joiners)

			if rand.Float32() < 0.1 {
				name = strings.Join([]string{lang.Words.Definite, name}, join)
			}
		}

		if (len(name) < params.MinLength) || (len(name) > params.MaxLength) {
			continue
		}

		used := false
		for _, name2 := range lang.Words.Names {
			if strings.Contains(name, name2) || strings.Contains(name2, name) {
				used = true
				break
			}
		}

		if used {
			continue
		}

		lang.Words.Names = append(lang.Words.Names, name)

		return name
	}
}
