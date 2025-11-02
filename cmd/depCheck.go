/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/luviz/lcmd/libs/dependencies"
	"github.com/luviz/lcmd/libs/writer"
	"github.com/spf13/cobra"
)

// depCheckCmd represents the depCheck command
var depCheckCmd = &cobra.Command{
	Use:   "depCheck",
	Short: "Check to see if dependency are installed on system",
	Run: func(cmd *cobra.Command, args []string) {
		res := dependencies.CheckDependencies()
		writer.New(res, output).Write()
	},
}

func init() {
	rootCmd.AddCommand(depCheckCmd)
	depCheckCmd.Flags().StringVarP(&output, "output", "o", "text", "setting output supports text, json, yaml")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// depCheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// depCheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
