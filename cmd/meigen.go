/*
Copyright © 2022 yamato0126

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// meigenCmd represents the meigen command
var meigenCmd = &cobra.Command{
	Use:   "meigen",
	Short: "random meigen",
	Long:  `This command gets a random meigen.`,
	Run: func(cmd *cobra.Command, args []string) {
		quizFlag, err := cmd.Flags().GetBool("quiz")
		if err != nil {
			log.Printf("Flags error. %v", err)
		} else if quizFlag {
			makeMeigenQuiz()
		} else {
			printMeigen()
		}
	},
}

func init() {
	meigenCmd.Flags().BoolP("quiz", "q", false, "Quiz function")
	rootCmd.AddCommand(meigenCmd)
}

func makeMeigenQuiz() {
	var authors [3]string
	var choice [3]string = [3]string{"A", "B", "C"}
	var input string
	var ans string

	statement, author := getRandomMeigen()
	fmt.Println(statement)
	fmt.Println("Who left this meigen?")
	authors[0] = author
	for i := 1; i < 3; i++ {
		_, author := getRandomMeigen()
		authors[i] = author
	}
	randomIndex := shaffleChoices(3)
	for i := 0; i < 3; i++ {
		fmt.Println(choice[i] + ": " + authors[randomIndex[i]])
		if randomIndex[i] == 0 {
			ans = choice[i]
		}
	}
	fmt.Println("Your answer -> ")
	fmt.Scan(&input)
	if input == ans {
		fmt.Println("Yes! You are correct.")
	} else {
		fmt.Println("No! you are wrong...")
		fmt.Println("Correct answer is " + ans)
	}
}

func printMeigen() {
	statement, author := getRandomMeigen()
	fmt.Println(statement)
	fmt.Println("by " + author)
}

func getRandomMeigen() (string, string) {
	url := "https://meigen.doodlenote.net/api/json.php?c=1"
	responseBytes := getMeigenData(url)

	response := string(responseBytes)
	split := strings.Index(response, "auther")
	return response[12 : split-3], response[split+9 : len(response)-3]
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

func allKeys(m map[int]bool) []int {
	i := 0
	result := make([]int, len(m))
	for key, _ := range m {
		result[i] = key
		i++
	}
	return result
}

func pickup(min int, max int, num int) []int {
	numRange := max - min

	selected := make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	for counter := 0; counter < num; {
		n := rand.Intn(numRange) + min
		if !selected[n] {
			selected[n] = true
			counter++
		}
	}
	keys := allKeys(selected)
	return keys
}

func shaffleChoices(num int) []int {
	results := pickup(0, num, num)
	return results
}
