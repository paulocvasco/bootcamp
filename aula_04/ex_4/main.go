package main

import (
	"bootcamp/aula_04/ex_4/types"
	"fmt"
	"math/rand"
)

func main() {
	variavel := rand.Perm(100)
	fmt.Println(variavel)
}

func insertion(nums types.Values) {
	var i = 1
	for i < nums.Len() {
		var j = i

		for j >= 1 && nums[j] < nums[j-1] {
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j--
		}

		i++
	}
}

func selection(nums types.Values) {
	for i := 0; i < nums.Len(); i++ {
		var minIdx = i
		for j := i; j < nums.Len(); j++ {
			if nums.Less(j, minIdx) {
				minIdx = j
			}
		}
		nums.Swap(i, minIdx)
	}
}

func count(nums types.Values) {
	var max = nums[0]

	var i = 1
	for i < nums.Len() {
		if nums[i] > max {
			max = nums[i]
		}

		i++
	}

	var indices = make([]int, max+1)

	var j = 0
	for j < nums.Len() {
		indices[nums[j]]++

		j++
	}

	var k = 1
	for k < len(indices) {
		indices[k] += indices[k-1]

		k++
	}

	var result = make([]int, nums.Len())

	var m = 0
	for m < nums.Len() {
		result[indices[nums[m]]-1] = nums[m]
		indices[nums[m]]--

		m++
	}
}
