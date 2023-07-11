package nekos_best

import "golang.org/x/exp/slices"

// Check wether a category is valid or not
func isValidCategory(cat string) bool {
	return slices.Contains(categories, cat)
}
