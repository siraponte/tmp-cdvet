package utils

// ContainsString checks if a string s is in a slice of strings.
// It returns true if the string is in the slice, false otherwise.
func ContainsString(slice []string, s string) bool {
	for _, entry := range slice {
		if entry == s {
			return true
		}
	}
	return false
}
