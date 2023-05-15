package config

import "strings"

func isQueryExpectedToReturnRows(query string) bool {
	lowerQuery := strings.ToLower(query)

	return strings.HasPrefix(lowerQuery, "select")
}
