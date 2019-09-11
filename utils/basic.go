package utils

import (
	"os"
	"regexp"
)

// LookupEnv reads environment variable but with fallback support
func LookupEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func replaceHibernateVariable(contents string, key string, value string) string {
	rgxp := regexp.MustCompile(`(?m)connection.` + key + `">.+<`)
	text := `connection.` + key + `">` + value + `<`
	return rgxp.ReplaceAllString(contents, text)
}
