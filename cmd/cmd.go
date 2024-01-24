package cmd

import (
	"UrlExtractor/data"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "UrlExtractor",
	Short: "this is a simple url extractor",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		executeParam = data.ExecuteParam{Url: strings.Trim(args[0], "")}
	},
}

var executeParam data.ExecuteParam

func Execute() (data.ExecuteParam, error) {
	if err := rootCmd.Execute(); err != nil {
		return data.ExecuteParam{}, err
	}
	if len(executeParam.Url) < 1 {
		return data.ExecuteParam{}, fmt.Errorf("please provide a url")
	}
	return executeParam, nil
}