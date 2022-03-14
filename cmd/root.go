package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

const use = "hankakuToZenkaku"
const short = "short desc"
const long = "long desc"

var rootCmd = &cobra.Command{
	Use:   use,
	Short: short,
	Long:  long,
	Run: func(cmd *cobra.Command, args []string) {
		// go run main.go aaaa ccc
		// args = []string{"aaaa", "ccc"}

		fmt.Printf("args = %#v\n\n", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hankakuToZenkaku.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
