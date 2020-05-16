// Package acronym converts an acronym to its abbreviation.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes the inputed words and abbreviates to their first letters capitalized.
func Abbreviate(s string) string {
	abb := ""
	re := regexp.MustCompile(`[A-Za-z']*`)
	wordSlice := re.FindAllString(s, -1)
	if len(wordSlice) > 0 {
		for _, word := range wordSlice {
			if word != "" {
				abb += strings.ToUpper(string(word[0]))
			}
		}
	}
	return abb
}
