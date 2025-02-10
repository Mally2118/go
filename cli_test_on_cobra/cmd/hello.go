/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// helloCmd представляет команду hello
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Выводит приветствие",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
