package delivery

import (
	"report/modules"

	"github.com/adlio/trello"
)

func init() {
	//@ get data from trello  api
}

type Delivery struct {
	keyApp   string
	token    string
	username string
	client   *trello.Client
	user     *trello.Member
}

func NewDelivery(key, tk, id string) (delivery Delivery) {
	delivery = Delivery{
		keyApp:   key,
		token:    tk,
		username: id,
	}
	delivery.client, delivery.user, _ = modules.ProvideMemberAndCLient(key, tk, id)
	return
}

func (d Delivery) GetCardsIsOpenOnWeek(id string, nameList string) ([]*trello.Card, error) {
	cards, err := d.client.SearchCards("board:"+id+" is:open sort:created created:week list:"+nameList, trello.Defaults())
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (d Delivery) GetCardsOnBoard(idBoard string) ([]*trello.Card, error) {
	board, err := d.client.GetBoard(idBoard, trello.Defaults())
	if err != nil {
		// Handle error
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		// Handle error
	}

	return cards, nil
}
