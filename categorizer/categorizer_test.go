package categorizer

import (
	"testing"
)

func TestAppendGameToGameSliceShouldWorkWithTwoItems(t *testing.T) {
	var err error
	categorizer := Categorizer{}

	err = categorizer.AppendGameToGameSlice("The Legend of Zelda", true, 10, 5)
	logFailedIfErrorExists(err, &t, "AppendGameToGameSlice")

	err = categorizer.AppendGameToGameSlice("Pokemon Shield", false, 8.5, 5)
	logFailedIfErrorExists(err, &t, "AppendGameToGameSlice")
}

func TestAppendGameToGameSliceShouldNotWork(t *testing.T) {
	categorizer := Categorizer{}

	err := categorizer.AppendGameToGameSlice("", false, 2, 5)

	if err == nil {
		t.Error("expected error on AppendGameToGameSlice function call with empty string as game name")
	}
}

func TestAppendGameToGameSliceShouldWork(t *testing.T) {
	categorizer := Categorizer{}

	categorizer.AppendGameToGameSlice("Pokemon Shield", false, 8.5, 5)

	if len(categorizer.GamesSlice) != 1 {
		t.Errorf("expected number of games to be 1. Got: %v", len(categorizer.GamesSlice))
	}
}

func TestCalculateAllGamesScoreShouldWork(t *testing.T) {
	var err error
	categorizer := Categorizer{}

	err = categorizer.AppendGameToGameSlice("The Legend of Zelda", true, 10, 5)
	logFailedIfErrorExists(err, &t, "AppendGameToGameSlice")

	err = categorizer.AppendGameToGameSlice("Pokemon Shield", false, 8.5, 5)
	logFailedIfErrorExists(err, &t, "AppendGameToGameSlice")

	categorizer.CalculateAllGamesScore()

	for _, game := range categorizer.GamesSlice {
		if game.FinalScore == 0 {
			t.Errorf("expected final score to have been calculated by the function. Game: %v", game.GameName)
		}
	}
}

func TestFormattedStringsShouldWork(t *testing.T) {
	var err error
	categorizer := Categorizer{}

	err = categorizer.AppendGameToGameSlice("The Legend of Zelda", true, 10, 5)
	logFailedIfErrorExists(err, &t, "AppendGameToGameSlice")

	err = categorizer.AppendGameToGameSlice("Pokemon Shield", false, 8.5, 5)
	logFailedIfErrorExists(err, &t, "AppendGameToGameSlice")

	categorizer.CalculateAllGamesScore()
	formattedStringsSlice, err := categorizer.GetStringSliceWithGameScoreResults()

	logFailedIfErrorExists(err, &t, "GetStringSliceWithGameScoreResults")

	for _, gameInfo := range formattedStringsSlice {
		if gameInfo == "" {
			t.Error("expected game information to return from GetStringSliceWithGameScoreResults")
		}
	}
}

func logFailedIfErrorExists(err error, t **testing.T, funcName string) {
	if err != nil {
		(*t).Errorf("expected no error return from %v. Got: %v", funcName, err)
	}
}