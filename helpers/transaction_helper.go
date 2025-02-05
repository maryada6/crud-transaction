package helpers

import (
	"crud-transaction/db"
	"crud-transaction/models"
)

func CalculateSum(transactionID int64) float64 {
	var transaction models.Transaction
	db.DB.First(&transaction, transactionID)

	total := transaction.Amount

	var childTransactions []models.Transaction
	db.DB.Where("parent_id = ?", transactionID).Find(&childTransactions)

	for _, child := range childTransactions {
		total += CalculateSum(child.ID)
	}

	return total
}
