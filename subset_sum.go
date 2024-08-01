package main

import "fmt"

func main() {
	// f2(0, []int{2, 3, 4}, make([]int, 0))
	x := make([]int, 0)
	f3(0, 0, []int{2, 3, 4}, make([]int, 0), &x)
	fmt.Println(x)
}

func f2(idx int, nums []int, ds []int) {
	if idx == len(nums) {
		s := 0
		for i := 0; i < len(ds); i++ {
			s += ds[i]
		}
		fmt.Println(ds, s)
		return
	}
	f2(idx+1, nums, append(ds, nums[idx]))
	f2(idx+1, nums, ds)
}

// optimised calculate the sum in the recursion itself
func f3(idx, sum int, nums []int, ds []int, coll *[]int) {
	if idx == len(nums) {
		*coll = append(*coll, sum)
		fmt.Println(ds, sum)
		return
	}
	f3(idx+1, sum+nums[idx], nums, append(ds, nums[idx]), coll)
	f3(idx+1, sum, nums, ds, coll)
}
