package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hoisie/mustache"
)

// DoConfigureSDK compiles settings.xml template and overrides file
func DoConfigureSDK() {

	template := sdkPomXMLTemplate

	// compatibility: from artifactory to reposilite
	if LookupEnv(txArtifactoryCompat, "0") == "1" {
		template = strings.ReplaceAll(template, "-local", "")
		fmt.Println("✅ updated repository manager compatibility")
	}

	data := mustache.Render(template, map[string]string{
		txArtifactoryURL: LookupEnv(txArtifactoryURL, "http://localhost:8080/artifactory"),
	})

	cwd, _ := os.Getwd()
	filePath := cwd + sdkPomXMLPath
	fileContents, _ := ioutil.ReadFile(filePath)
	updatedFileContents := bytes.Replace(fileContents, []byte(sdkPomTag), []byte(data+sdkPomTag), -1)

	ioutil.WriteFile(filePath, []byte(updatedFileContents), 0777)

	fmt.Println("✅ Updated " + filePath)
}
