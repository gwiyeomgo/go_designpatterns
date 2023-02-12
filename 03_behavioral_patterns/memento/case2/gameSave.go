package main

type GameSave struct {
	blueTeamScore int
	redTeamScore  int
}

func (m *GameSave) getSavedBlueTeamScore() int {
	return m.blueTeamScore
}

func (m *GameSave) getSavedRedTeamScore() int {
	return m.redTeamScore
}
