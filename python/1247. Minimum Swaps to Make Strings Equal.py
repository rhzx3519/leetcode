class Solution(object):
    def minimumSwap(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: int
        """
        n = len(s1)
        cnt1 = cnt2 = 0
        for i in range(n):
            if s1[i]==s2[i]: # 相同不处理
                continue
            if s1[i]=='x': # 不同x数量+1
                cnt1 += 1
            else:
                cnt2 += 1 # 不同y数量+1
        if (cnt1 + cnt2)&1: # 要满足
            return -1
        return cnt1//2 + cnt2//2 + (cnt1%2)*2


