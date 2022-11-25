package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	MOD := int(1000000007)

	var stack []int
	prevSmaller := int(0)

	dp := make([]int, len(arr))
	for i := int(0); i < len(arr); i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevSmaller = stack[len(stack)-1]
			dp[i] = dp[prevSmaller] + (i-prevSmaller)*arr[i]
		} else {
			dp[i] = (i + 1) * arr[i]
		}
		stack = append(stack, i)
	}

	sum := int(0)
	for _, val := range dp {
		sum += val
	}
	return sum % MOD
}

func main() {
	arr := []int{3, 1, 2, 4}
	result := sumSubarrayMins(arr)
	fmt.Println(result)
}
