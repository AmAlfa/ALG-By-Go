package two_sum

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target - nums[i]]; ok && v != i{
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

//m:=make(map[int]int, len(nums))
//
//    for i,v:=range nums {
//        if idx, ok := m[target-v]; ok {
//            if i>idx {
//                i,idx = idx,i
//            }
//            return []int{i,idx}
//        }
//        m[v]=i
//    }
//    return nil

func TestTwoSum(t *testing.T)  {
	arr := []int{2, 7, 1, 10}
	t.Logf("%v\n", twoSum(arr, 17))
}
