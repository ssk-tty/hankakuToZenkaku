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
	split into package
 	add zenMap
*/

var rootCmd = &cobra.Command{
	Use:   use,
	Short: short,
	Long:  long,
	Run: func(cmd *cobra.Command, args []string) {
		// go run main.go "123 ccc" aaa

		fmt.Printf("args = %#v\n\n", args)

		sentence := strings.Join(args, " ")

		splitWords := strings.Split(sentence, "")

		result := convert(splitWords)

		fmt.Printf("result: %s\n", result)
	},
}

func convert(words []string) []byte {
	zenMap := getZenMap()

	var result []byte

	for _, word := range words {
		w := hanToZen(zenMap, word)
		fmt.Printf("old: %s, new: %s\n", word, w)
		result = append(result, w...)
	}

	return result
}

func hanToZen(wordMap map[string]string, word string) string {
	if val, ok := wordMap[word]; ok {
		return val
	}

	return word
}

func getZenMap() map[string]string {
	zenMap := make(map[string]string)

	zenMap["1"] = "１"
	zenMap["2"] = "２"
	zenMap["3"] = "３"
	zenMap["4"] = "４"
	zenMap["5"] = "５"
	zenMap["6"] = "６"
	zenMap["7"] = "７"
	zenMap["8"] = "８"
	zenMap["9"] = "９"
	zenMap["0"] = "０"
	zenMap[" "] = "　"

	return zenMap
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
