/*
Copyright © 2022 yamato0126

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// meigenCmd represents the meigen command
var meigenCmd = &cobra.Command{
	Use:   "meigen",
	Short: "random meigen",
	Long:  `This command gets a random meigen`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomMeigen()
	},
}

func init() {
	rootCmd.AddCommand(meigenCmd)
}

func getRandomMeigen() {
	url := "https://meigen.doodlenote.net/api/json.php?c=1"
	responseBytes := getMeigenData(url)

	response := string(responseBytes)
	split := strings.Index(response, "auther")
	fmt.Println(response[12 : split-3])
	fmt.Println("by " + response[split+9:len(response)-3])
}

func getMeigenData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not request a meigen. %v", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}
