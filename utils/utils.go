package utils

import "fmt"


func FloatToMoney(money float64) string {
	// result := strconv.FormatFloat(money, 'E', -1, 64)
	return fmt.Sprintf("%.2f", money)
}