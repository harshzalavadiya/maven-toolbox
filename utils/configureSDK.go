package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hoisie/mustache"
)

// DoConfigureSDK compiles settings.xml template and overrides file
func DoConfigureSDK() {
	data := mustache.Render(sdkPomXMLTemplate, map[string]string{
		txArtifactoryURL: LookupEnv(txArtifactoryURL, "http://localhost:8080/artifactory"),
	})

	cwd, _ := os.Getwd()
	filePath := cwd + sdkPomXMLPath
	fileContents, _ := ioutil.ReadFile(filePath)
	updatedFileContents := bytes.Replace(fileContents, []byte(sdkPomTag), []byte(data+sdkPomTag), -1)

	ioutil.WriteFile(filePath, []byte(updatedFileContents), 0777)

	fmt.Println("🗸 Updated " + filePath)
}
