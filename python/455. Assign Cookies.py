class Solution(object):
    def findContentChildren(self, g, s):
        """
        :type g: List[int]
        :type s: List[int]
        :rtype: int
        """
        g.sort()
        s.sort()
        res = 0
        i = j = 0
        while i<len(g) and j<len(s):
            if g[i]<=s[j]:
                i += 1
                j += 1
                res += 1
            else:
                j += 1
        return res
        
# 思路：简单贪心，分别对g和s从小到大排序
#       如果g[i]<=s[j]，i、j都往后移动，否则j往后移动
# time: O(N), space: O(1)