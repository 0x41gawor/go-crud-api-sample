package repo

import "strings"

func removeSpacesFromStringSlice(input []string) []string {
	result := make([]string, len(input))
	for i, str := range input {
		result[i] = strings.ReplaceAll(str, " ", "")
	}
	return result
}
