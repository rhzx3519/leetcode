package main

/**
https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-ii/description/?envType=daily-question&envId=2024-07-11
*/
func incremovableSubarrayCount(nums []int) int64 {
    var ans int64 = 0
    length := len(nums)
    l := 0
    for l < length-1 {
        if nums[l] >= nums[l+1] {
            break
        }
        l++
    }

    if l == length-1 {
        return int64(length) * int64(length+1) / 2
    }
    ans += int64(l + 2)
    for r := length - 1; r > 0; r-- {
        if r < length-1 && nums[r] >= nums[r+1] {
            break
        }
        for l >= 0 && nums[l] >= nums[r] {
            l--
        }
        ans += int64(l + 2)
    }

    return ans
}
