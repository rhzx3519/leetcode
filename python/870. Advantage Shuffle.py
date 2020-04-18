import bisect
class Solution(object):
    def advantageCount(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: List[int]
        """
        ans = []
        A.sort()
        for b in B:
            if A[0]>b or A[-1]<=b:
                ans.append(A[0])
                A.pop(0)
            else:
                idx = bisect.bisect_left(A, b)
                ans.append(A[idx])
                A.pop(idx)
        return ans



if __name__ == '__main__':
    A = [12,24,8,32]
    B = [13,25,32,11]
    su = Solution()
    su.advantageCount(A, B)