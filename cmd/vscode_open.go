package cmd

import (
	"path"
	"strings"

	"github.com/alfred_tools/alfred"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vscodeOpenCmd = &cobra.Command{
	Use:   "vscode_open",
	Short: "vscode_open",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			wf.SendFeedback()
		}()
		var query string
		if len(args) > 0 {
			query = args[0]
		}
		alfred.WalkDir(viper.GetString("VSCODE_SEARCH_DIR"), 3, func(p string) {
			if strings.Contains(p, query) {
				_, file := path.Split(p)
				wf.NewItem(file).Subtitle(p).Arg(p).Valid(true)
			}
		})

		wf.SendFeedback()
	},
}
