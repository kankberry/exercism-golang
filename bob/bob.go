// Package bob returns strings that bob likes to say.
package bob

import "strings"

// Hey returns strings based on input.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case strings.HasSuffix(remark, "?") && strings.ToUpper(remark) == remark && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?") && remark[len(remark)-1] == '?':
		return "Sure."
	case strings.ToUpper(remark) == remark && strings.ToUpper(remark) != strings.ToLower(remark):
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
