package handlers

import (
	"crud-transaction/db"
	"crud-transaction/models"
)

func calculateSum(transactionID int64) float64 {
	var transaction models.Transaction
	db.GetDB().First(&transaction, transactionID)

	total := transaction.Amount

	var childTransactions []models.Transaction
	db.GetDB().Where("parent_id = ?", transactionID).Find(&childTransactions)

	for _, child := range childTransactions {
		total += calculateSum(child.ID)
	}

	return total
}
