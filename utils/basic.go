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

func replaceTag(contents string, key string, value string) string {
	rgxp := regexp.MustCompile(`<` + key + `>.+<`)
	text := `<` + key + `>` + value + `<`
	return rgxp.ReplaceAllString(contents, text)
}

func replaceArtifactory(contents string) string {
	rgxp := regexp.MustCompile(`<url>http.+/artifactory`)
	artifactoryURL := LookupEnv(txArtifactoryURL, "http://localhost:8080/artifactory")
	text := `<url>` + artifactoryURL
	return rgxp.ReplaceAllString(contents, text)
}
