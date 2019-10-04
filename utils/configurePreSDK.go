package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// DoConfigurePreSDK updates sdk endpoint in pom.xml
func DoConfigurePreSDK() {
	// Swagger SDK pom.xml
	cwd, _ := os.Getwd()
	rootFilePath := cwd + sdkRootPomXMLPath
	rootFileContents, _ := ioutil.ReadFile(rootFilePath)
	updatedRootFileContents := replaceTag(string(rootFileContents), "schemes", LookupEnv("MTPROP_SCHEMES", "httpx"))
	updatedRootFileContents = replaceTag(updatedRootFileContents, "host", LookupEnv("MTPROP_HOST", "localhostx"))

	ioutil.WriteFile(rootFilePath, []byte(updatedRootFileContents), 0777)

	fmt.Println("🗸 Updated " + rootFilePath)
}
