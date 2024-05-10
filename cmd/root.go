/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	fa "github.com/schmiddim/font-awesome-golang/download"
	"github.com/schmiddim/font-awesome-golang/lib"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your generate command",
	Long:  `A longer description of your generate command`,
	Run: func(cmd *cobra.Command, args []string) {

		f := fa.FontAwesome{}
		err := f.FetchMedataJson("5.x")
		if err != nil {
			fmt.Println(err)
		}

		f.ParseMetaData()
		err = f.GenerateGoFileFromIcons()
		if err != nil {
			log.Fatal(err)
		}

	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your test command",
	Long:  `A longer description of your test command`,
	Run: func(cmd *cobra.Command, args []string) {
		s := lib.GetIconForString("Read")
		fmt.Println(s)
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "font-awesome-golang",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(testCmd)
}
