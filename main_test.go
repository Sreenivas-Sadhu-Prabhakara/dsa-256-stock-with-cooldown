package main

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 0, 2}, 3}, // buy sell cool buy sell
		{[]int{1}, 0},
		{[]int{}, 0},
		{[]int{2, 1}, 0}, // falling market: never trade
		{[]int{1, 2}, 1},
		{[]int{1, 2, 4}, 3},     // hold through, single sale
		{[]int{2, 10, 1, 3}, 8}, // sell at 10; the day-2 dip is unreachable (cooldown)
		{[]int{6, 1, 3, 2, 4, 7}, 6},
	}
	for _, c := range cases {
		if got := MaxProfit(c.in); got != c.want {
			t.Errorf("MaxProfit(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
