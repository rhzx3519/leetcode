class Solution(object):
    def largeGroupPositions(self, s):
        """
        :type s: str
        :rtype: List[List[int]]
        """
        ans = []
        start = end = 0
        for i, ch in enumerate(s):
            if i==0 or s[i]==s[i-1]:
                end = i
            else:
                # print start, end
                if end - start + 1 >= 3:
                    ans.append([start, end])
                start = end = i
        if end - start + 1 >= 3:
            ans.append([start, end])
        return ans