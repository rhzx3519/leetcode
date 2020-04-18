class Solution(object):
    def maxDepthAfterSplit(self, seq):
        """
        :type seq: str
        :rtype: List[int]
        """
        res = []
        depth = 0
        for ch in seq:
            if ch=='(':
                depth += 1
                res.append(depth%2)
            else:
                res.append(depth%2)
                depth -= 1
        return res



if __name__ == '__main__':
    seq = "()(())()"
    su = Solution()
    print su.maxDepthAfterSplit(seq)

# 思路： 根据深度将奇数深度的分到A，偶数的分到B