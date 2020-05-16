// Package raindrops converst raindrops to string sounds
package raindrops

import "strconv"

// Convert returns a string given a number/raindrop
func Convert(num int) string {
	switch {
	case num%105 == 0:
		return "PlingPlangPlong"
	case num%15 == 0:
		return "PlingPlang"
	case num%21 == 0:
		return "PlingPlong"
	case num%35 == 0:
		return "PlangPlong"
	case num%3 == 0:
		return "Pling"
	case num%5 == 0:
		return "Plang"
	case num%7 == 0:
		return "Plong"
	default:
		return strconv.Itoa(num)
	}
}
