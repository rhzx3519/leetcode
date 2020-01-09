class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ''
        pre = strs[0]
        i = 1
        size = len(pre)
        while i < len(strs):
            t = strs[i]
            j = 0
            while j < min(size, len(t)):
                if pre[j] != t[j]:
                    break
                j += 1
            size = min(size, j)
            i += 1
            
        return pre[0:size]