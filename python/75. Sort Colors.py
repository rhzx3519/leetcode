class Solution(object):
    def sortColors(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        i = 0
        l = 0
        r = n-1
        while i <= min(n-1, r):
            if nums[i]==0:
                if i==l:
                    i += 1
                else:
                    nums[i], nums[l] = nums[l], nums[i]
                l += 1
            elif nums[i]==2:
                nums[i], nums[r] = nums[r], nums[i]
                r -= 1
            else:
                i += 1

# 思路：使用l作为给0留的交换位置，r作为给2留的交换位置，当nums[i]等于0时，
#       交换nums[i]和nums[l]；当nums[i]等于2时，交换nums[i]和nums[r]
# 时间复杂度: O(n), 空间复杂度: O(1)
                
            