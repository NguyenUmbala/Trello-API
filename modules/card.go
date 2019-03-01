package modules

import (
	"github.com/adlio/trello"
)

type MyCard struct {
	ID               string
	Name             string
	TimeGuessforDone int
	TimeRealForDone  int
}

func (c MyCard) NewCard(card *trello.Card) (mycard MyCard) {
	mycard.ID = card.ID
	mycard.Name = card.Name
	mycard.TimeGuessforDone = GetTimeGuessForDone(card.Name)
	mycard.TimeRealForDone = GetRealTimeOfDone(card.Name)
	return
}
