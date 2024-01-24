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
		executeParam.Url = strings.TrimSpace(args[0])
	},
}

var (
	executeParam data.ExecuteParam
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&executeParam.FileToWrite, "output", "o", "", "file to write the output to")
	rootCmd.PersistentFlags().StringVarP(&executeParam.InputFile, "input", "i", "", "input file to read the urls from")
	rootCmd.Flags().IntVar(&executeParam.Depth, "depth", 1, "depth of the crawling")
	rootCmd.Flags().StringVar(&executeParam.Regex, "regex", "", "regex filter links")
}

func Execute() (data.ExecuteParam, error) {
	if err := rootCmd.Execute(); err != nil {
		return data.ExecuteParam{}, err
	}
	if len(executeParam.Url) < 1 {
		return data.ExecuteParam{}, fmt.Errorf("please provide a url")
	}
	return executeParam, nil
}
