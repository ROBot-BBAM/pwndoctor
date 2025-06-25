/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/strykethru/pwndoctor/pkg/pwndoctor"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Import called")

		pwndoctor.Init(PwnDocURL)
		pwndoctor.AutoAuth()

		var engagementList []string

		if EngagementName != "" {
			engagementList = append(engagementList, strings.Split(EngagementName, ",")...)
			if len(engagementList) > 1 {
				fmt.Println("\n[!] Alert: You are attempting to import findings to multiple reports. Aborting out of caution")
				os.Exit(1)
			}
		} else {
			fmt.Println("[!] Alert: You must supply one and only one engagement name! Aborting.")
			os.Exit(1)
		}

		pwndoctor.DoImport(EngagementName, findingsDir)

	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVarP(&PwnDocURL, "url", "u", "", "PwnDoc-NG URL (i.e. https://127.0.0.1:8443)")
	importCmd.Flags().StringVarP(&pwndocSSHHost, "ip", "i", "", "PwnDoc-NG SSH IP (for mongodb dump)")
	importCmd.Flags().StringVarP(&pwndocSSHUser, "user", "U", "ubuntu", "PwnDoc-NG SSH user (for mongodb dump)")
	importCmd.Flags().StringVarP(&EngagementName, "engagement", "e", "", "Engagment where you want to input the findings.")
	importCmd.Flags().StringVarP(&findingsDir, "engDir", "f", "", "Directory of PwnDoc engagement directory (i.e. /path/to/engagement/directory)")

}
