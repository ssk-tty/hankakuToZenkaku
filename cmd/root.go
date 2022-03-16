package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var cfgFile string

const use = "hankakuToZenkaku"
const short = "short desc"
const long = "long desc"

/**
todo
	split into packages
	add zenMap
	delete debug log
	fix config
*/

var rootCmd = &cobra.Command{
	Use:   use,
	Short: short,
	Long:  long,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {

	fmt.Printf("args = %#v\n\n", args)

	sentence := strings.Join(args, " ")

	splitWords := strings.Split(sentence, "")

	zenMap := GetZenMap()

	result := Convert(zenMap, splitWords)

	fmt.Printf("result: %s\n", result)
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
