// Package twofer will return the name of the person I'm sharing with.
package twofer

// ShareWith will return the name of the person I'm sharing with or "you" if no name provided.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
