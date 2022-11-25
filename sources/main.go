package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	MOD := int(1000000007)

	stack := NewStack()
	prevSmaller := int(0)

	dp := make([]int, len(arr))
	for i := range arr {
		top, _ := stack.Peek()
		for !stack.IsEmpty() && arr[top] >= arr[i] {
			stack.Pop()
		}
		if !stack.IsEmpty() {
			prevSmaller, _ = stack.Peek()
			dp[i] = dp[prevSmaller] + (i-prevSmaller)*arr[i]
		} else {
			dp[i] = (i + 1) * arr[i]
		}
		stack.Push(i)
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
