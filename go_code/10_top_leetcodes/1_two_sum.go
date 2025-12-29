/*

两数之和 (Two Sum)

题目描述：给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。

示例：给定 nums = [2, 7, 11, 15], target = 9，因为 nums[0] + nums[1] = 2 + 7 = 9，所以返回 [0, 1]。

解题思路：使用哈希表（在Go中可以用map）存储遍历过的数字及其索引，对于每个元素，检查target与当前元素的差是否在map中。

*/


package main

import (
    "fmt"
)

func main() {
    fmt.Print("hello world !")
    nums := []int{2, 7, 11, 15}
    s, e := twoSum(nums, 13)
    fmt.Printf("s: %d, e: %d", s, e)
}

func twoSum(nums []int, target int) (int, int) {
    // 双指针解法
    leftIndex := 0
    rightIndex := len(nums) - 1
    for rightIndex > leftIndex {
        if nums[leftIndex] + nums[rightIndex] == target {
            return leftIndex, rightIndex
        } else if nums[leftIndex] + nums[rightIndex] > target {
            rightIndex--
        } else {
            leftIndex++
        }
    }
    return -1, -1
}
