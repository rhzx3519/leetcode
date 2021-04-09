import collections
class Solution(object):
    def subarraysWithKDistinct(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: int
        """
        return self.haveMostK(A, K) - self.haveMostK(A, K - 1)

    def haveMostK(self, A, K):
        n = len(A)
        l = 0
        ans = 0
        count = collections.defaultdict(int)
        for r in range(n):
            count[A[r]] += 1
            while l <= r and len(count) > K:
                count[A[l]] -= 1
                if count[A[l]] == 0: del count[A[l]]
                l += 1
            ans += r - l
        # print ans
        return ans

# 思路：恰有K个不同的数字的子数组数目 = 最多有K个不同数字的子数组数目 - 最多有K-1个不同数字的子数组数目


if __name__ == '__main__':
    # A = [1,2,1,2,3]
    # K = 2
    # A = [1,2,1,3,4]
    # K = 3
    # A = [1,2]
    # K = 1
    su = Solution()
    print su.subarraysWithKDistinct(A, K)