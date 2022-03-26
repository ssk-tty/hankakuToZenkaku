package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	. "hankakuToZenkaku/convertor"
	"os"
	"strings"
)

var cfgFile string

const use = "hankakuToZenkaku"
const short = "short desc"
const long = "long desc"

/**
todo
	update zenMap
	fix config
*/

var rootCmd = &cobra.Command{
	Use:   use,
	Short: short,
	Long:  long,
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
	sentence := strings.Join(args, " ")

	splitWords := strings.Split(sentence, "")

	wordConvertor := NewWordConvertor(splitWords)

	result := wordConvertor.ToZen()

	fmt.Printf("%s\n", result)
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
