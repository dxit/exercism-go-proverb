package proverb

import "fmt"

const (
	rhymeModel string = "For want of a %s the %s was lost."
	finalModel string = "And all for the want of a %s."
)

// Proverb returns the full chain of strings
func Proverb(rhyme []string) []string {
	var proverb []string
	if len(rhyme) > 0 {
		prev := rhyme[0]
		for _, strophe := range rhyme[1:] {
			proverb = append(proverb, fmt.Sprintf(rhymeModel, prev, strophe))
			prev = strophe
		}
		proverb = append(proverb, fmt.Sprintf(finalModel, rhyme[0]))
	}
	return proverb
}
