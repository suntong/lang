package leetcode

import "fmt"

func candy0(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] += candies[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	total := 0
	for _, candy := range candies {
		total += candy + 1
	}
	return total
}

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}
	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
			for j := i; j > 0; j-- {
				// fmt.Println(j, candies[j-1], candies[j], ratings[j-1], ratings[j])
				// fmt.Println(candies[j-1] <= candies[j], ratings[j-1] > ratings[j])
				if candies[j-1] <= candies[j] && ratings[j-1] > ratings[j] {
					candies[j-1]++
				} else {
					break
				}
			}
		}
	}

	total := 0
	fmt.Println(ratings)
	for _, candy := range candies {
		fmt.Printf("%d ", candy)
		total += candy
	}
	fmt.Println()
	return total
}
