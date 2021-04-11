package categorizer

import (
	"errors"
	"fmt"
)

type Categorizer struct {
	GamesSlice []game
}

type game struct {
	GameName string
	PlayedBefore bool
	MedianScoreInReviewSites float32
	AmountOfTimeTheGameWillBePlayerFromOneToFive uint16
	FinalScore float32
}

// Creates a new Game instance and appends it to its GamesSlice property
func (categorizerInstance *Categorizer) AppendGameToGameSlice(gameName string, playedBefore bool, medianScoreInReviewSites float32, howLongWillTheGameBePlayed uint16) error {
	if gameName == "" {
		return errors.New("the name of the game cannot be empty")
	}

	game := game{
		GameName: gameName,
		PlayedBefore: playedBefore, 
		MedianScoreInReviewSites: medianScoreInReviewSites, 
		AmountOfTimeTheGameWillBePlayerFromOneToFive: howLongWillTheGameBePlayed}

	categorizerInstance.GamesSlice = append(categorizerInstance.GamesSlice, game)

	return nil
}

// Calculates the score for each game inside the GamesSplit slice.
func (categorizerInstance *Categorizer) CalculateAllGamesScore() {
	gameSliceSize := len(categorizerInstance.GamesSlice)
	for i := 0; i < gameSliceSize; i++ {
		categorizerInstance.GamesSlice[i].CalculateGameScore()
	}
}

// Returns a list of formatted strings containing an ""information breakdown"" of each game in the GamesSlice property.
func (categorizerInstance *Categorizer) GetSortedStringSliceWithGameScoreResults() ([]string, error) {
	gameSliceSize := len(categorizerInstance.GamesSlice)
	var formattedInformationAboutEachGame []string

	if gameSliceSize <= 0 {
		return formattedInformationAboutEachGame,
		errors.New("there is no game list in the instance. Please set at least one game by calling 'AppendGameToGameSlice' before calling this function")
	}

	categorizerInstance.GamesSlice = sortGamesSlice(categorizerInstance.GamesSlice)

	for i := 0; i < gameSliceSize; i++ {
	 	formattedInformationAboutEachGame = append(formattedInformationAboutEachGame, 
			fmt.Sprintf("%v's score: %v", categorizerInstance.GamesSlice[i].GameName, categorizerInstance.GamesSlice[i].FinalScore))
	}

	return formattedInformationAboutEachGame, nil
}

func sortGamesSlice(gamesSlice []game) []game {
	var gameNameHistory []string
	gameSliceSize := len(gamesSlice)
	var sortedGamesSlice []game
	var currentGameWithBiggestScore = game{}

	for len(sortedGamesSlice) < gameSliceSize {
		for i := 0; i < gameSliceSize; i++ {
			if gamesSlice[i].FinalScore > currentGameWithBiggestScore.FinalScore && !stringExistsInSlice(gameNameHistory, gamesSlice[i].GameName) {
				currentGameWithBiggestScore = gamesSlice[i]
				gameNameHistory = append(gameNameHistory, currentGameWithBiggestScore.GameName)
			}
		}

		sortedGamesSlice = append(sortedGamesSlice, currentGameWithBiggestScore)
		currentGameWithBiggestScore = game{}
	}

	return sortedGamesSlice
}

func stringExistsInSlice(slice []string, s string) bool {
	sliceSize := len(slice)
	
	for i := 0; i < sliceSize; i++ {
		if (slice[i] == s)  {
			return true;
		}
	}

	return false
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