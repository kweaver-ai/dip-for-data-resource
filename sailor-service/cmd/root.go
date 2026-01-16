package main

import "github.com/spf13/cobra"

var confPath string

var rootCmd = &cobra.Command{
	Use:     "af-sailor service",
	Short:   "af-sailor service",
	Version: Version,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&confPath, "conf", "c", "cmd/server/config", "config path, eg: -conf config")
	rootCmd.AddCommand(serverCmd)
}
