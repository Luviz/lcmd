/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var output string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lcmd",
	Short: "Quality of life utilities",
	Long:  `A collection of utilities to speed up some work flows`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("output", "o", "setting output supports text, json, yaml")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
