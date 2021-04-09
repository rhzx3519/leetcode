#!usr/bin/env python  
#-*- coding:utf-8 _*-  


class Solution(object):
    def maxSumRangeQuery(self, nums, requests):
        """
        :type nums: List[int]
        :type requests: List[List[int]]
        :rtype: int
        """
        n = len(nums)
        diff = [0]*(n+1)
        for l, r in requests:
            diff[l] += 1
            diff[r+1] -= 1
        # print diff
        cnt = [0]*n
        cnt[0] = diff[0]
        for i in range(1, n):
            cnt[i] = cnt[i-1] + diff[i]
        cnt.sort()
        nums.sort()
        ans = 0
        for i in range(n):
            ans += cnt[i] * nums[i]
        # print ans
        return ans%(10**9+7)

# 思路：通过requests统计每个下标出现的次数，对次数多的下标分配较大的数字
# 使用差分数组记录一段区间内每个下标出现的次数，diff=[0]*(n+1),
# 例如(1, 3), diff[1]+=1, diff[3+1]-=1, cnt[1] = diff[1] + cnt[i-1]
# time: O(N), space: O(N)

if __name__ == '__main__':
    nums = [1,2,3,4,5]
    requests = [[1,3],[0,1]]
    su = Solution()
    su.maxSumRangeQuery(nums, requests)