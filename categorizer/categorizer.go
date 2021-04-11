package categorizer

import (
	"fmt"
	"log"
)

type Categorizer struct {
	GamesSplit []game
	NumberOfGames int
}

type game struct {
	GameName string
	PlayedBefore bool
	MedianScoreInReviewSites float32
	AmountOfTimeTheGameWillBePlayerFromOneToFive uint16
	FinalScore float32
}

// Set the NumberOfGames property of the Categorizer instance
func (categorizerInstance *Categorizer) SetNumberOfGames(numberOfGames int) {
	categorizerInstance.NumberOfGames = numberOfGames
}

// Asks for each property in the Game struct for a value from 
// STDIN then sets the game.
func (categorizerInstance *Categorizer) AppendGameToGameSlice(gameName string, playedBefore bool, medianScoreInReviewSites float32, howLongWillTheGameBePlayed uint16) {
	game := game{
		GameName: gameName,
		PlayedBefore: playedBefore, 
		MedianScoreInReviewSites: medianScoreInReviewSites, 
		AmountOfTimeTheGameWillBePlayerFromOneToFive: howLongWillTheGameBePlayed}

	categorizerInstance.GamesSplit = append(categorizerInstance.GamesSplit, game)
}

// Calculates the score for each game inside the GamesSplit slice.
func (categorizerInstance *Categorizer) CalculateAllGamesScore() {
	gameSplitSize := len(categorizerInstance.GamesSplit)
	for i := 0; i < gameSplitSize; i++ {
		categorizerInstance.GamesSplit[i].CalculateGameScore()
	}
}

// Prints the results to STDOUT.
// If the GamesSplit slice is not present or does not have more than 0 items, 
// log.Fatalf is called.
func (categorizerInstance *Categorizer) ReturnFormattedStringsForEachGameAndTheirResults() []string {
	gameSplitSize := len(categorizerInstance.GamesSplit)
	formattedInformationAboutEachGame := make([]string, gameSplitSize)

	if gameSplitSize <= 0 {
		log.Fatalf("there is no game list in the instance. Please set at least one game by calling 'AppendGameToGameSlice' before calling this function")
	}

	for i := 0; i < gameSplitSize; i++ {
	 	formattedInformationAboutEachGame = append(formattedInformationAboutEachGame, 
			fmt.Sprintf("%v's score: %v", categorizerInstance.GamesSplit[i].GameName, categorizerInstance.GamesSplit[i].FinalScore))
	}

	return formattedInformationAboutEachGame
}

// Calculates the score for the current game instance
func (currentGame *game) CalculateGameScore() {
	var total float32

	if currentGame.PlayedBefore {
		total -= 2
	}

	total += currentGame.MedianScoreInReviewSites

	total += float32(currentGame.AmountOfTimeTheGameWillBePlayerFromOneToFive)

	currentGame.FinalScore = total
}