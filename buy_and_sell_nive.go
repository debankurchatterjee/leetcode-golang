package main

import "fmt"

func buy_sell(price []int) int {
	maxProfit := 0
	for i := 0; i < len(price); i++ {
		for j := i + 1; j < len(price); j++ {
			if price[i] < price[j] {
				maxProfit = max(maxProfit, price[j]-price[i])
			}
		}
	}
	return maxProfit
}

func buy_sell_better(price []int) int {
	currentPrice := price[0]
	maxProfit := 0
	for i := 1; i < len(price); i++ {
		currentPrice = min(currentPrice, price[i])
		currentProfit := price[i] - currentPrice
		maxProfit = max(maxProfit, currentProfit)
	}
	return maxProfit
}

func main() {
	fmt.Println(buy_sell([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(buy_sell_better([]int{7, 1, 5, 3, 6, 4}))
}
