package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBCard struct {
	ID     uint   `json:"id"`
	IDCard string `json:"idcard"`
	Name   string `json:"name"`
	Due    string `json:"due"`
}

var db *gorm.DB

func SaveCard(card DBCard) {
	if db.NewRecord(card) {
		db.Create(&card)
	}
}

func GetCards() []DBCard {
	var cards []DBCard
	db.Find(&cards)

	return cards
}

func UpdateCard(card DBCard) {
	var newCard DBCard
	db.Where("id_card = ?", card.IDCard).First(&newCard)

	if newCard.IDCard == "" { // Create new card
		db.Create(&card)
	} else { // Update old card
		newCard.Name = card.Name
		newCard.Due = card.Due
		db.Save(&newCard)
	}
}

func OpenConnection() {
	var err error
	db, err = gorm.Open("sqlite3", "./card.db")
	if err != nil {
		// Handle error
	}
}
func AutoMigrate() {
	db.AutoMigrate(&DBCard{})
}

func CloseConnection() {
	db.Close()
}
