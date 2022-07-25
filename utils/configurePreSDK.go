package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// DoConfigurePreSDK updates sdk endpoint in pom.xml
func DoConfigurePreSDK() {
	// Swagger SDK pom.xml
	cwd, _ := os.Getwd()
	rootFilePath := cwd + sdkRootPomXMLPath
	rootFileContents, _ := ioutil.ReadFile(rootFilePath)
	updatedRootFileContents := replaceTag(string(rootFileContents), "schemes", LookupEnv("MTPROP_SCHEMES", "httpx"))
	updatedRootFileContents = replaceTag(updatedRootFileContents, "host", LookupEnv("MTPROP_HOST", "localhostx"))
	updatedRootFileContents = replaceArtifactory(updatedRootFileContents)

	// compatibility: from artifactory to reposilite
	if LookupEnv(txArtifactoryCompat, "0") == "1" {
		updatedRootFileContents = strings.ReplaceAll(updatedRootFileContents, "-local", "")
		fmt.Println("✅ updated repository manager compatibility")
	}

	ioutil.WriteFile(rootFilePath, []byte(updatedRootFileContents), 0777)

	fmt.Println("✅ Updated " + rootFilePath)
}
