package accumulate

// Accumulate does things
func Accumulate(input []string, fn func(string) string) []string {

	output := make([]string, len(input)) // preallocate memory for output

	for i, v := range input {
		output[i] = fn(v)
	}
	return output
}
