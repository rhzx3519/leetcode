class Solution(object):
    def subarraysDivByK(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: int
        """
        n = len(A)
        res = 0
        mem = collections.defaultdict(int)
        mem[0] = 1

        for i in range(n):
            if i>0:
                A[i] += A[i-1]
            

        for i in range(n):
            mod = A[i]%K
            if mod in mem:
                res += mem[mod]
            mem[mod] += 1
        return res

# 思路：求前缀和，前缀和mod K 相同的两个下标i, j之间的数组和可以被K整除
# time: O(N), space: O(N)