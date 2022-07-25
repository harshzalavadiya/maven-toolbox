package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hoisie/mustache"
)

// DoConfigureM2 compiles settings.xml template and overrides file
func DoConfigureM2() {
	data := mustache.Render(settingsXMLTemplate, map[string]string{
		txArtifactoryUsername: LookupEnv(txArtifactoryUsername, "admin"),
		txArtifactoryPassword: LookupEnv(txArtifactoryPassword, "password"),
		txArtifactoryURL:      LookupEnv(txArtifactoryURL, "http://localhost:8080/artifactory"),
	})

	userHome, _ := os.UserHomeDir()
	filePath := userHome + m2SettingsXMLPath

	ioutil.WriteFile(filePath, []byte(data), 0777)

	fmt.Println("✅ Updated " + filePath)
}
