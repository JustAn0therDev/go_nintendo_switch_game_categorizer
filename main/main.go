package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/JustAn0therDev/go_switch_game_relevance_categorizer/categorizer"
)

var terminalBufferMap = make(map[string]func())

func init() {
	terminalBufferMap["linux"] = func () {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	terminalBufferMap["windows"] = func () {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	var numberOfGames int
	var gameName string
	var playedBefore bool
	var medianScoreInReviewSites float32
	var howLongWillTheGameBePlayed uint16

	categorizerInstance := categorizer.Categorizer{}

	fmt.Print("Insert the number of games to compare: ")
	fmt.Scanln(&numberOfGames)

	for i := 0; i < numberOfGames; i++ {
		terminalBufferMap[runtime.GOOS]()
		fmt.Printf("Current game: %v\n", i + 1)
	
		fmt.Print("Insert the name of the game (without spaces): ")
		fmt.Scanln(&gameName)
	
		fmt.Print("Have you played this game before (1/0)? ")
		fmt.Scanln(&playedBefore)
	
		fmt.Print("What is the median score in review sites for this game? ")
		fmt.Scanln(&medianScoreInReviewSites)
	
		fmt.Print("For how long will you play the game (from one to five)? ")
		fmt.Scanln(&howLongWillTheGameBePlayed)

		categorizerInstance.AppendGameToGameSlice(gameName, playedBefore, medianScoreInReviewSites, howLongWillTheGameBePlayed)
	}

	categorizerInstance.CalculateAllGamesScore()
	gamesInformation, err := categorizerInstance.GetStringSliceWithGameScoreResults()

	if err != nil {
		log.Fatal(err)
	}

	for _, game := range gamesInformation {
		fmt.Println(game)
	}
}