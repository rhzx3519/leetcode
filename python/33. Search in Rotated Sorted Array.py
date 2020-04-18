class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        l = 0
        r = len(nums)-1
        while l<=r:
            mid = l + (r-l)/2
            if nums[mid]==target:
                return mid
            elif nums[mid] < nums[r]:
                if nums[mid] < target and target <= nums[r]:
                    l = mid+1
                else:
                    r = mid-1
            else:
                if nums[mid] > target and target >= nums[l]:
                    r = mid-1
                else:
                    l = mid+1
        
        return -1

# 思路：如果中间的数小于最右边的数，则右半段是有序的，若中间数大于最右边数，则左半段是有序的，
# 我们只要在有序的半段里用首尾两个数组来判断目标值是否在这一区域内，这样就可以确定保留哪半边了        