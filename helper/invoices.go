package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInvoice() string {
	timestamp := time.Now().Unix()
	randomNum := rand.Intn(1000)
	invoiceNumber := fmt.Sprintf("%d-%d", timestamp, randomNum)
	return invoiceNumber
}
