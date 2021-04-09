class Solution(object):
    def rangeSum(self, nums, n, left, right):
        """
        :type nums: List[int]
        :type n: int
        :type left: int
        :type right: int
        :rtype: int
        """
        a = []
        pre = [0]*(n+1)
        for i in range(1, n+1):
            pre[i] = nums[i-1] + pre[i-1]
        # print pre
        for i in range(n+1):
            for j in range(i+1, n+1):
                a.append(pre[j] - pre[i])
        a.sort()
        return int(sum(a[left-1:right])%(1e9+7))
        
