class Solution(object):
    def orderlyQueue(self, S, K):
        """
        :type S: str
        :type K: int
        :rtype: str
        """
        return ''.join(sorted(S)) if K > 1 else min([S[i:] + S[:i] for i in range(len(S))])

# 思路：当K>1, 返回全排列中最小的字符串，当K==1时，返回循环序中最小的字符串