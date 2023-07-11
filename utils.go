package nekos_best

// Check wether a category is valid or not
func isValidCategory(cat string) bool {
	for _, c := range categories {
		if cat == c {
			return true
		}
	}
	return false
}
