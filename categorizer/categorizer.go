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
func (categorizerInstance *Categorizer) GetStringSliceWithGameScoreResults() ([]string, error) {
	gameSliceSize := len(categorizerInstance.GamesSlice)
	var formattedInformationAboutEachGame []string

	if gameSliceSize <= 0 {
		return formattedInformationAboutEachGame,
		errors.New("there is no game list in the instance. Please set at least one game by calling 'AppendGameToGameSlice' before calling this function")
	}

	for i := 0; i < gameSliceSize; i++ {
	 	formattedInformationAboutEachGame = append(formattedInformationAboutEachGame, 
			fmt.Sprintf("%v's score: %v", categorizerInstance.GamesSlice[i].GameName, categorizerInstance.GamesSlice[i].FinalScore))
	}

	return formattedInformationAboutEachGame, nil
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