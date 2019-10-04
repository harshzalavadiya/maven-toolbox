package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// DoConfigureProperties updates properties file
func DoConfigureProperties(filePathPre string, prefix string) {

	cwd, _ := os.Getwd()
	filePath := cwd + string(os.PathSeparator) + filePathPre
	props, err := Read(filePath, "")
	if err != nil {
		fmt.Println(err)
	}

	fileContents, _ := ioutil.ReadFile(filePath)
	fileContentsString := string(fileContents)

	for _, e := range os.Environ() {
		canUpdate, k, v := findInProperties(props, e, prefix)
		if canUpdate {
			fileContentsString = updateProperty(fileContentsString, k, v)
			fmt.Println(`Found ` + k)
		}
	}

	ioutil.WriteFile(filePath, []byte(fileContentsString), 0777)
	fmt.Println(`🗸 Updated ` + filePath)
}

func findInProperties(props map[string]string, e string, prefix string) (bool, string, string) {
	env := strings.Split(e, "=")
	for k := range props {
		envProp := strings.ToLower(env[0])
		fileProp := strings.ToLower(prefix + "_" + strings.Replace(k, ".", "_", -1))
		if envProp == fileProp {
			return true, k, env[1]
		}
	}
	return false, "", ""
}

func updateProperty(contents string, key string, value string) string {
	rgxp := regexp.MustCompile(`(?m)` + key + `=.+`)
	text := key + `=` + value
	return rgxp.ReplaceAllString(contents, text)
}
