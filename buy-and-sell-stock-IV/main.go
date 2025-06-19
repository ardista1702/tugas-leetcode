package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int, k int) int {

	n := len(prices)

	stockTransactionRecord := make([][]int, k+1)
	fmt.Println("hai")

	for i := range stockTransactionRecord {
		stockTransactionRecord[i] = make([]int, n)
	}

	for i := 1; i <= k; i++ {
		hold := -prices[0]
		for j := 1; j < n; j++ {
			stockTransactionRecord[i][j] = max(stockTransactionRecord[i][j-1], hold+prices[j])
			hold = max(hold, stockTransactionRecord[i-1][j]-prices[j])
		}
	}
	return stockTransactionRecord[k][n-1]
}

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 3
	result := maxProfit(prices, k)
	fmt.Println(result)

}
