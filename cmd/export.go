package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/strykethru/pwndoctor/pkg/pwndoctor"
)

var EngagementName string
var pwndocSSHHost string
var pwndocSSHUser string
var findingsDir string

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Do all the things to close up PwnDoc after engagement.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if findingsDir != "" {
			pwndoctor.ExportCSVLocally(findingsDir)
		} else {
			pwndoctor.Init(PwnDocURL)
			pwndoctor.AutoAuth()

			var engagementList []string

			if EngagementName != "" {
				engagementList = append(engagementList, strings.Split(EngagementName, ",")...)
			}

			pwndoctor.DoExport(engagementList, pwndocSSHUser, pwndocSSHHost)
		}

	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&PwnDocURL, "url", "u", "", "PwnDoc-NG URL (i.e. https://127.0.0.1:8443)")
	exportCmd.Flags().StringVarP(&pwndocSSHHost, "ip", "i", "", "PwnDoc-NG SSH IP (for mongodb dump)")
	exportCmd.Flags().StringVarP(&pwndocSSHUser, "user", "U", "ubuntu", "PwnDoc-NG SSH user (for mongodb dump)")
	exportCmd.Flags().StringVarP(&EngagementName, "engagement", "e", "", "Engagement Name")
	exportCmd.Flags().StringVarP(&findingsDir, "findingsDir", "f", "", "Directory of JSON files from an old PwnDoc/Report Archive (i.e. /path/to/JSON/Findings)")
}
