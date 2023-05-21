package utils

import "strings"

func queryNormalization(query string) string {
	querySplited := strings.Split(query, " ")

	return querySplited
}
