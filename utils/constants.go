package utils

import "os"

const doneMessage = "✨ Done!"

const m2SettingsXMLPath = string(os.PathSeparator) + ".m2" + string(os.PathSeparator) + "settings.xml"
const sdkPomXMLPath = string(os.PathSeparator) + "target" + string(os.PathSeparator) + "sdk" + string(os.PathSeparator) + "pom.xml"
const sdkPomTag = "<build>"
const hibernateFileName = "hibernate.cfg.xml"

const (
	txArtifactoryUsername = "ARTIFACTORY_USERNAME"
	txArtifactoryPassword = "ARTIFACTORY_PASSWORD"
	txArtifactoryURL      = "ARTIFACTORY_URL"
)

const settingsXMLTemplate = `<?xml version="1.0" encoding="UTF-8" ?>
<settings xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.1.0 http://maven.apache.org/xsd/settings-1.1.0.xsd" xmlns="http://maven.apache.org/SETTINGS/1.1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <servers>
        <server>
            <username>{{ARTIFACTORY_USERNAME}}</username>
            <password>{{ARTIFACTORY_PASSWORD}}</password>
            <id>central</id>
        </server>
        <server>
            <username>{{ARTIFACTORY_USERNAME}}</username>
            <password>{{ARTIFACTORY_PASSWORD}}</password>
            <id>snapshots</id>
        </server>
    </servers>
    <profiles>
        <profile>
            <repositories>
                <repository>
                    <snapshots>
                        <enabled>false</enabled>
                    </snapshots>
                    <id>central</id>
                    <name>libs-release</name>
                    <url>{{ARTIFACTORY_URL}}/libs-release</url>
                </repository>
                <repository>
                    <snapshots />
                    <id>snapshots</id>
                    <name>libs-snapshot</name>
                    <url>{{ARTIFACTORY_URL}}/libs-snapshot</url>
                </repository>
            </repositories>
            <pluginRepositories>
                <pluginRepository>
                    <snapshots>
                        <enabled>false</enabled>
                    </snapshots>
                    <id>central</id>
                    <name>libs-release</name>
                    <url>{{ARTIFACTORY_URL}}/libs-release</url>
                </pluginRepository>
                <pluginRepository>
                    <snapshots />
                    <id>snapshots</id>
                    <name>libs-snapshot</name>
                    <url>{{ARTIFACTORY_URL}}/libs-snapshot</url>
                </pluginRepository>
            </pluginRepositories>
            <id>artifactory</id>
        </profile>
    </profiles>
    <activeProfiles>
        <activeProfile>artifactory</activeProfile>
    </activeProfiles>
</settings>`

const sdkPomXMLTemplate = `<distributionManagement>
        <repository>
            <id>central</id>
            <name>jfrog-releases</name>
            <url>{{ARTIFACTORY_URL}}/libs-release-local</url>
        </repository>
        <snapshotRepository>
            <id>snapshots</id>
            <name>jfrog-snapshots</name>
            <url>{{ARTIFACTORY_URL}}/libs-snapshot-local</url>
        </snapshotRepository>
        </distributionManagement>
        <repositories>
        <repository>
            <snapshots>
                <enabled>false</enabled>
            </snapshots>
            <id>central</id>
            <name>libs-release</name>
            <url>{{ARTIFACTORY_URL}}/libs-release</url>
        </repository>
        <repository>
            <snapshots />
            <id>snapshots</id>
            <name>libs-snapshot</name>
            <url>{{ARTIFACTORY_URL}}/libs-snapshot</url>
        </repository>
    </repositories>
`

//Descriptions of commands
const (
	DescriptionM2         = `Update '~/.m2/settings.xml'`
	DescriptionSDK        = `Update 'target/sdk/pom.xml'`
	DescriptionHibernate  = `Finds and updates 'hibernate.cfg.xml'`
	DescriptionProperties = `Update properties file
  Args:
  <file>    relative properties file path
  <prefix>  environment variable prefix (prefix will be stripped with underscore, case insensitive)`
	ArgFilePath  = `properties path`
	ArgEnvPrefix = `Environment variable prefix`
)
