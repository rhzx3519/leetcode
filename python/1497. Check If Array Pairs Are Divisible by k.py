class Solution(object):
    def canArrange(self, arr, k):
        """
        :type arr: List[int]
        :type k: int
        :rtype: bool
        """
        count = [0]*k
        for num in arr:
            r = num%k
            if r < 0:
                r += k
            count[r] += 1
        if count[0]&1: # 余数是0的个数要等于偶数个
            return False
        for i in range(1, k):
            if count[i]!=count[k-i]:
                return False
        return True

# 思路：获取k的所有余数个数， 余数x要与k-x的个数相等
# time: O(N), space: O(K)