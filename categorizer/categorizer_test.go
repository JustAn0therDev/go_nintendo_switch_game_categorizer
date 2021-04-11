package categorizer

import "testing"

func TestSetNumberOfGamesShouldWork(t *testing.T) {
	numberOfGames := 3
	categorizer := Categorizer{}

	categorizer.SetNumberOfGames(numberOfGames)

	if categorizer.NumberOfGames != 3 {
		t.Errorf("Expected number of games to be %v. Got %v", numberOfGames, categorizer.NumberOfGames)
	}
}

func TestNumberOfGamesShouldNotWork(t *testing.T) {
	categorizer := Categorizer{}

	categorizer.SetNumberOfGames(0)

	if categorizer.NumberOfGames != 0 {
		t.Errorf("Expected number of games to be 0. Got %v", categorizer.NumberOfGames)
	}	
}