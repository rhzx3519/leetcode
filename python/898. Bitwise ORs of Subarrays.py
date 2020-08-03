class Solution(object):
    def subarrayBitwiseORs(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        res, cur = set([]), set([])
        for a in A:
            cur = {a|b for b in cur}
            cur.add(a)
            res |= cur
        return len(res)
        
if __name__ == '__main__':
    A = [1,1,2]
    su = Solution()
    print su.subarrayBitwiseORs(A)