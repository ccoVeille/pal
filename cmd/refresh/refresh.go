package refresh

import (
	"fmt"

	"github.com/jaytyrrell13/pal/cmd/make"
	"github.com/jaytyrrell13/pal/pkg"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var RefreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Delete `~/.pal` file and run `make` command",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunRefreshCmd()
		cobra.CheckErr(err)
	},
}

func RunRefreshCmd() error {
	AppFs := afero.NewOsFs()

	aliasFilePath, aliasFilePathErr := pkg.AliasFilePath()
	if aliasFilePathErr != nil {
		return aliasFilePathErr
	}

	if pkg.FileMissing(AppFs, aliasFilePath) {
		fmt.Println("~/.pal file is missing")
		return nil
	}

	fmt.Println("Removing ~/.pal file")
	removeFileErr := pkg.RemoveFile(AppFs, aliasFilePath)
	if removeFileErr != nil {
		return removeFileErr
	}

	fmt.Println("Running `make` command")
	return make.RunMakeCmd()
}
