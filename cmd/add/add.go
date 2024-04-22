package add

import (
	"os"

	"github.com/jaytyrrell13/pal/pkg"
	"github.com/jaytyrrell13/pal/pkg/prompts"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	nameFlag string
	pathFlag string
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create an alias for an additional directory",
	Run: func(cmd *cobra.Command, args []string) {
		AppFs := afero.NewOsFs()
		if pkg.FileMissing(AppFs, pkg.AliasFilePath()) {
			cobra.CheckErr("~/.pal file is missing, please run make command first")
		}

		name := nameFlag
		path := pathFlag

		if name == "" {
			name = prompts.StringPrompt("What is the name of the alias?")
		}

		if path == "" {
			path = prompts.StringPrompt("What is the path for the alias?")
		}

		saveExtraDirErr := pkg.SaveExtraDir(path)
		cobra.CheckErr(saveExtraDirErr)

		aliasesFile, openAliasesFileErr := os.OpenFile(pkg.AliasFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o755)
		cobra.CheckErr(openAliasesFileErr)

		c, readConfigFileErr := pkg.ReadFile(AppFs, pkg.ConfigFilePath())
		cobra.CheckErr(readConfigFileErr)

		jsonConfig, fromJsonErr := pkg.FromJson(c)
		cobra.CheckErr(fromJsonErr)

		output := pkg.MakeAliasCommands(name, path, jsonConfig)

		if _, err := aliasesFile.Write([]byte(output)); err != nil {
			aliasesFile.Close()
			cobra.CheckErr(err)
		}

		if err := aliasesFile.Close(); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	AddCmd.Flags().StringVarP(&nameFlag, "name", "n", "", "Name of the additional alias")
	AddCmd.Flags().StringVarP(&pathFlag, "path", "p", "", "Path to your additional directory")
}
