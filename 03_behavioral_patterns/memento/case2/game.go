package main

type Game struct {
	blueTeamScore int
	redTeamScore  int
}

func (e *Game) save() *GameSave {
	return &GameSave{blueTeamScore: e.blueTeamScore, redTeamScore: e.redTeamScore}
}

func (e *Game) restore(m *GameSave) {
	e.blueTeamScore = m.getSavedBlueTeamScore()
	e.redTeamScore = m.getSavedRedTeamScore()
}

func (e *Game) setBlueTeamScore(blueTeamScore int) {
	e.blueTeamScore = blueTeamScore
}

func (e *Game) setRedTeamScore(redTeamScore int) {
	e.redTeamScore = redTeamScore
}

func (e *Game) getBlueTeamScore() int {
	return e.blueTeamScore
}

func (e *Game) getRedTeamScore() int {
	return e.redTeamScore
}
