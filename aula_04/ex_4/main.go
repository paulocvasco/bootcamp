package main

import (
	"bootcamp/aula_04/ex_4/types"
	"fmt"
	"math/rand"
)

func main() {
	variavel := rand.Perm(100)
	fmt.Println(variavel)
	count(variavel)
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
	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < nums.Len()-1 {
			if !nums.Less(i, i+1) {
				nums.Swap(i, i+1)
				isDone = false
			}
			i++
		}
	}
}
