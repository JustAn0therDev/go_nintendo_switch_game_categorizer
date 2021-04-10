package main

import "github.com/JustAn0therDev/go_switch_game_relevance_categorizer/categorizer"

func main() {
	categorizerInstance := categorizer.Categorizer{}

	categorizerInstance.AskForNumberOfGames()
	categorizerInstance.AskForGamesName()
	categorizerInstance.PrintAllGameScores()
}