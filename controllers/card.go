package controllers

import (
	"report/database"
	"report/modules"

	"github.com/adlio/trello"

	"github.com/gin-gonic/gin"
)

func AllCardReview(c *gin.Context) {
	idBoard := c.Param("id")
	var mycard modules.MyCard
	var tmpReview []modules.MyCard
	var tmpReviewed []modules.MyCard

	cardsOnReview, err := Consumer.GetCardsIsOpenOnWeek(idBoard, "review-me")
	cardsOnReviewed, err := Consumer.GetCardsIsOpenOnWeek(idBoard, "Done")

	if err != nil {
	}

	for i := 0; i < len(cardsOnReview); i++ {
		tmpReview = append(tmpReview, mycard.NewCard(cardsOnReview[i]))
	}
	for i := 0; i < len(cardsOnReviewed); i++ {
		tmpReviewed = append(tmpReviewed, mycard.NewCard(cardsOnReviewed[i]))
	}

	c.JSON(200, gin.H{
		"review":      tmpReview,
		"review-done": tmpReviewed,
	})
}

func AllCardChangedDueDate(c *gin.Context) {
	database.OpenConnection()
	defer database.CloseConnection()
	idBoard := c.Param("id")
	cardsOnBoard, err := Consumer.GetCardsOnBoard(idBoard)
	if err != nil {
		// Handle error
	}

	cardsOnDB := database.GetCards()
	var cardsChangedDueDate []*trello.Card
	for i, _ := range cardsOnBoard {
		for j, _ := range cardsOnDB {
			if CheckChangeDue(cardsOnBoard[i], cardsOnDB[j]) {
				cardsChangedDueDate = append(cardsChangedDueDate, cardsOnBoard[i])
			}
		}
	}

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}

// Check due date of 2 card
func CheckChangeDue(cardsBoard *trello.Card, cardsDB database.DBCard) bool {
	if cardsBoard.ID != cardsDB.IDCard {
		return false
	}
	if cardsBoard.Due == nil && cardsDB.Due == "" {
		return false
	}
	if cardsBoard.Due != nil && cardsBoard.Due.String() == cardsDB.Due {
		return false
	}
	return true
}

// Bước tiền đề cho API. Lưu tất cả card để có data mới kiểm tra sự thay đổi ngày hết hạn của các card được
func SaveCardsOnDB(c *gin.Context) {
	database.OpenConnection()
	database.AutoMigrate()
	defer database.CloseConnection()

	idBoard := c.Param("id")

	cardsOnBoard, err := Consumer.GetCardsOnBoard(idBoard)
	if err != nil {
		// Handle error
	}

	for _, v := range cardsOnBoard {
		// Set due  for card if card on board don't have due
		due := ""
		if v.Due != nil {
			due = v.Due.String()
		}

		tmpCard := database.DBCard{
			IDCard: v.ID,
			Name:   v.Name,
			Due:    due,
		}
		database.SaveCard(tmpCard)
	}

	cards := database.GetCards()
	c.JSON(200, gin.H{
		"Cards on database": cards,
	})
}

func UpdateCards(c *gin.Context) {
	database.OpenConnection()
	defer database.CloseConnection()

	idBoard := c.Param("id")
	cardsOnBoard, err := Consumer.GetCardsOnBoard(idBoard)
	if err != nil {
		// Handle error
	}

	for _, v := range cardsOnBoard {
		due := ""
		if v.Due != nil {
			due = v.Due.String()
		}

		tmpCard := database.DBCard{
			IDCard: v.ID,
			Name:   v.Name,
			Due:    due,
		}
		database.UpdateCard(tmpCard)
	}

	c.JSON(200, gin.H{
		"Update card": database.GetCards(),
	})
}
