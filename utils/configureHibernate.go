package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DoConfigureHibernate compiles settings.xml template and overrides file
func DoConfigureHibernate() {

	cwd, _ := os.Getwd()
	filepath.Walk(cwd, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, hibernateFileName) {
			fileContents, _ := ioutil.ReadFile(path)
			updatedFileContents := replaceHibernateVariable(string(fileContents), "username", LookupEnv("DB_USERNAME", "postgres"))
			updatedFileContents = replaceHibernateVariable(updatedFileContents, "password", LookupEnv("DB_PASSWORD", "postgres123"))
			updatedFileContents = replaceHibernateVariable(updatedFileContents, "url", LookupEnv("DB_URL", "jdbc:postgresql://localhost:5432/biodiv"))
			ioutil.WriteFile(path, []byte(updatedFileContents), 0777)
			fmt.Println(`🗸 Updated ` + path)
		}
		return nil
	})

}
