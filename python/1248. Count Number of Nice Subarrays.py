class Solution(object):
    def numberOfSubarrays(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        n = len(nums)
        for i in range(n):
            nums[i] = nums[i]&1
        memo = collections.defaultdict(int)
        memo[0] = 1 # memo【i】表示和为i的字数组个数
        s = 0
        res = 0
        for i in range(n):
            s += nums[i]
            res += memo[s-k]
            memo[s] += 1
        return res

# 思路：区间和

if __name__ == '__main__':
    nums = [2,2,2,1,2,2,1,2,2,2]
    k = 2
    su = Solution()
    print su.numberOfSubarrays(nums, k)