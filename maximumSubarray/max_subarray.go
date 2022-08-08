package main

const INF = 1<<32 - 1

func sumArray(nums []int, start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += nums[i]
	}
	return sum
}

func MaxSubArray1(nums []int) int {
	N := len(nums)
	maxSum := -INF
	for start := 0; start < N; start++ {
		for end := start; end < N; end++ {
			ans := sumArray(nums, start, end)
			if ans > maxSum {
				maxSum = ans
			}
		}
	}
	return maxSum
}

// Divide and Conquer
// ans(nums, n) =
//  max {a[1] + ans(nums, n-1)[2] + ans(nums, n-1)[1]},
//  max {a[1] , a[1] + ans(nums, n-1)[2]}
func max(n1, n2 int) int {
	if n2 > n1 {
		return n2
	} else {
		return n1
	}
}

func maxSubArrayHelper(nums []int, start int) (int, int) {
	if len(nums)-1 < start {
		return -INF, -INF
	} else {
		maxSum, maxConSum := maxSubArrayHelper(nums, start+1)
		newMaxConSum := max(nums[start], nums[start]+maxConSum)
		return max(newMaxConSum, maxSum), newMaxConSum
	}
}

func MaxSubArray2(nums []int) int {
	ans, _ := maxSubArrayHelper(nums, 0)
	return ans
}

func MaxSubArray3(nums []int) int {
	maxSubCon := nums[0]
	maxSubSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+maxSubCon > nums[i] {
			maxSubCon = nums[i] + maxSubCon
		} else {
			maxSubCon = nums[i]
		}
		if maxSubCon > maxSubSum {
			maxSubSum = maxSubCon
		}

	}
	return maxSubSum
}
