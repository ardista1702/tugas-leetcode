package main

import "fmt"



func max(a,b int)int {
	if a > b {
		return  a
	}
	return  b
}

func backtrack(prices []int,day int,k int,holding bool) int {
	if day == len(prices) || k == 0 {
		return 0
	}
	if holding{
		skipToday := backtrack(prices,day + 1,k,true) //profit ketika kita melewatkan hari ini
		sellToday := prices[day] + backtrack(prices,day + 1,k-1,false)
		return  max(skipToday,sellToday)
	}else{
		skipToday := backtrack(prices,day + 1,k,false)
		buyToday := -prices[day] + backtrack(prices,day + 1,k,true)
		return  max(skipToday,buyToday)
	}
}

func maxProfit(prices []int ,k int) int {
	return  backtrack(prices,0,k,false)
}

func main() {
	prices := []int{3,2}
	k := 1
	res := maxProfit(prices,k)
	fmt.Println(res)
}