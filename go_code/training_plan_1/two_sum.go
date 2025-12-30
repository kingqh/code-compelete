// Two Sum
// 给定一个整数数组 `nums` 和一个目标值 `target`，
// 在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，且同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
//
// 示例：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：nums[0] + nums[1] == 9
//
// 要求：
// - 实现函数 `twoSum(nums []int, target int) []int`
// - 时间复杂度期望 O(n)，空间复杂度 O(n)
// - 返回值为长度为 2 的 slice，包含满足条件的两个下标
// - 处理包含负数、零和重复数字的情况
//
// 解题思路（简要）：
// 使用哈希表（map[int]int）记录已访问元素的值到下标映射。
// 遍历数组时，对于当前元素 v，检查 map 中是否存在 key = target - v。
// 若存在则返回对应下标与当前下标；否则将当前元素及其下标存入 map。
// 该方法能在一次遍历内找到答案，时间 O(n)，空间 O(n)。
package main
import "fmt"

func main() {
   nums := []int{2, 4, 3, 5, 9, 6, 7}
   fmt.Printf("nums is: %v, result is %v", nums, twoSum(nums, 6)) 
}

func twoSum(nums []int, target int) []int {
    intmap := make(map[int]int)
    for i, v := range nums {
        key := target - v 
        if k, ok := intmap[key]; ok {
            return []int{i, k}
        } else {
            intmap[v] = i
        }
    }
    return []int{}
}

