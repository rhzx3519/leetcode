class Solution(object):
    def minDeletionSize(self, A):
        """
        :type A: List[str]
        :rtype: int
        """
        def isSorted(arr):
            for i in range(1, len(arr)):
                if arr[i] < arr[i-1]:
                    return False
            return True

        ans = 0
        N = len(A)
        W = len(A[0])
        cur = ['']*N
        for j in range(W):
            cur2 = cur[:]
            for i in range(N):
                cur2[i] += A[i][j]
            # print cur2, isSorted(cur2)
            if isSorted(cur2):
                cur = cur2[:]
            else:
                ans += 1
        return ans

# 思路：贪心，如果当前A[i]加上前面的列，是有序的，则不必删除第i列，否则删除第i列