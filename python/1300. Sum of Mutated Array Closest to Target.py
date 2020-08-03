class Solution(object):
    def findBestValue(self, arr, target):
        """
        :type arr: List[int]
        :type target: int
        :rtype: int
        """
        n = len(arr)
        arr.sort()
        pre = 0
        for i in range(n):
            cur = pre + (n-i)*arr[i]
            if cur >= target:
                value = (target - pre)//(n-i)
                sum_low = pre + value*(n-i)
                sum_high = sum_low + n-i
                return value if (target - sum_low) <= (sum_high - target) else value + 1
            pre += arr[i]
        return arr[-1]

# 思路：比较前缀和，如果当前前缀和>=target, 则value必定在value = (target - pre)//(n-i), (value, value+1)之间
# time: O(N), space: O(1)