class Solution(object):
    def partitionDisjoint(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        l_max = A[0]
        a_max = A[0]
        idx = 0
        for i in range(len(A)):
            a_max = max(a_max, A[i])
            if A[i] < l_max:
                l_max = a_max
                idx = i
        return idx + 1

# 思路: l_max记录左数组最大值, a_max记录当前最大值
            