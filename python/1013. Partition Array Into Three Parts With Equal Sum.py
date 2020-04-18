class Solution(object):
    def canThreePartsEqualSum(self, A):
        """
        :type A: List[int]
        :rtype: bool
        """
        s = sum(A)
        n = len(A)
        i = 0
        target = s//3
        cur = 0
        while i < n:
            cur += A[i]
            if cur==target: break
            i += 1
        if cur!=target: return False
        j = i+1
        while j < n:
            cur += A[j]
            if cur==target*2: break
            j += 1
        return j<n-1


if __name__ == '__main__':
    A = [1,-1,1,-1]
    su = Solution()
    print su.canThreePartsEqualSum(A)