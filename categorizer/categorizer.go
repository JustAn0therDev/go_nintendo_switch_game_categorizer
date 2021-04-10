package categorizer

import (
	"fmt"
)


type Categorizer struct {
	gamesSplit []game
	numberOfGames int
}

type game struct {
	GameName string
	PlayedBefore bool
	MedianScoreInReviewSites float32
	AmountOfTimeTheGameWillBePlayerFromOneToFive uint16
	FinalScore float32
}

func (categorizerInstance *Categorizer) AskForNumberOfGames() {
	fmt.Print("Insert the number of games: ")
	fmt.Scanln(&categorizerInstance.numberOfGames)
}

func (categorizerInstance *Categorizer) AskForGamesName() {
	var gameName string
	var playedBefore bool
	var medianScoreInReviewSites float32
	var howLongWillTheGameBePlayed uint16

	for i := 0; i < categorizerInstance.numberOfGames; i++ {
		fmt.Printf("Current game: %v\n", i + 1)

		fmt.Print("Insert the name of the game (without spaces): ")
		fmt.Scanln(&gameName)

		fmt.Print("Have you played this game before (1/0)? ")
		fmt.Scanln(&playedBefore)

		fmt.Print("What is the median score in review sites for this game? ")
		fmt.Scanln(&medianScoreInReviewSites)

		fmt.Print("For how long will you play the game (from one to five)? ")
		fmt.Scanln(&howLongWillTheGameBePlayed)

		game := game{
			GameName: gameName,
			PlayedBefore: playedBefore, 
			MedianScoreInReviewSites: medianScoreInReviewSites, 
			AmountOfTimeTheGameWillBePlayerFromOneToFive: howLongWillTheGameBePlayed}

		categorizerInstance.gamesSplit = append(categorizerInstance.gamesSplit, categorizerInstance.getGameScoreCalculated(&game))
	}
}

func (categorizerInstance *Categorizer) getGameScoreCalculated(currentGame *game) game {
	var total float32

	if currentGame.PlayedBefore {
		total -= 2
	}

	total += currentGame.MedianScoreInReviewSites

	total += float32(currentGame.AmountOfTimeTheGameWillBePlayerFromOneToFive)

	currentGame.FinalScore = total

	return *currentGame
}

func (categorizerInstance *Categorizer) PrintAllGameScores() {
	for i := 0; i < len(categorizerInstance.gamesSplit); i++ {
		fmt.Printf("%v's score: %v\n", categorizerInstance.gamesSplit[i].GameName, categorizerInstance.gamesSplit[i].FinalScore)
	}
}