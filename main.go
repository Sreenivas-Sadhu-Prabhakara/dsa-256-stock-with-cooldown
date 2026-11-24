package main

import "fmt"

// MaxProfit returns the best profit trading one share repeatedly when every
// sale is followed by a one-day cooldown. Three-state machine per day:
// holding, just sold (cooling), and free to buy.
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	const minInt = -int(^uint(0)>>1) - 1
	hold := minInt // best profit while holding a share
	sold := 0      // best profit having sold today (cooldown tomorrow)
	rest := 0      // best profit flat and free to buy
	for _, p := range prices {
		prevSold := sold
		sold = hold + p // sell what we hold
		if rest-p > hold {
			hold = rest - p // buy today (only from the free state)
		}
		if prevSold > rest {
			rest = prevSold // cooldown ends: yesterday's sale becomes free cash
		}
	}
	if sold > rest {
		return sold
	}
	return rest
}

func main() {
	fmt.Println(MaxProfit([]int{1, 2, 3, 0, 2}))
}
