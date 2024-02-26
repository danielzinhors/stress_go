package cmd

import (
	"os"

	"github.com/danielzinhors/stress_go/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress_go",
	Short: "A stress tester em Golang",
	Long:  "A stress tester elaborado para o desafio final da pos-goexpert da Full Cycle",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.PersistentFlags().GetString("url")
		requests, _ := cmd.PersistentFlags().GetInt64("requests")
		concurrency, _ := cmd.PersistentFlags().GetInt64("concurrency")
		internal.RunStressTester(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("url", "u", "", "URL para testar")
	rootCmd.PersistentFlags().Int64P("requests", "r", 0, "total requests para realizar testes")
	rootCmd.PersistentFlags().Int64P("concurrency", "c", 0, "numero maximo requeste em concorrencias")
	rootCmd.MarkPersistentFlagRequired("url")
	rootCmd.MarkPersistentFlagRequired("requests")
	rootCmd.MarkPersistentFlagRequired("concurrency")
}
