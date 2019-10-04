package cmd

import (
	"github.com/harshzalavadiya/maven-toolbox/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cmdM2         = kingpin.Command("configure-m2", utils.DescriptionM2)
	cmdPreSDK     = kingpin.Command("configure-pre-sdk", utils.DescriptionPreSDK)
	cmdSDK        = kingpin.Command("configure-sdk", utils.DescriptionSDK)
	cmdHibernate  = kingpin.Command("configure-hibernate", utils.DescriptionHibernate)
	cmdProperties = kingpin.Command("configure-properties", utils.DescriptionProperties)
	filePath      = cmdProperties.Arg("file", utils.ArgFilePath).Required().String()
	prefix        = cmdProperties.Arg("prefix", utils.ArgEnvPrefix).Default("MTPROP").String()
)

// Execute ...
func Execute() {
	kingpin.Version("1.1.4")

	switch kingpin.Parse() {

	case cmdM2.FullCommand():
		utils.DoConfigureM2()

	case cmdPreSDK.FullCommand():
		utils.DoConfigurePreSDK()

	case cmdSDK.FullCommand():
		utils.DoConfigureSDK()

	case cmdHibernate.FullCommand():
		utils.DoConfigureHibernate()

	case cmdProperties.FullCommand():
		utils.DoConfigureProperties(*filePath, *prefix)

	}
}
