#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def partitionLabels(self, S):
        """
        :type S: str
        :rtype: List[int]
        """
        n = len(S)
        lastList = [0]*26
        for i in range(n-1, -1, -1):
            ch = ord(S[i]) - ord('a')
            if lastList[ch] == 0:
                lastList[ch] = i

        res = []
        l = r = 0
        last = 0
        for i in range(n):
            ch = ord(S[i]) - ord('a')
            lastIdx = lastList[ch]
            r = max(r, lastIdx)

            if l==r:
                res.append(r-last+1)
                l += 1
                r = last = l
            else:
                l += 1

        return res

if __name__ == '__main__':
    S = "ababcbacadefegdehijhklij"
    su = Solution()
    su.partitionLabels(S)                    


# 输入: S = "ababcbacadefegdehijhklij"
# 输出: [9,7,8]
# 解释:
# 划分结果为 "ababcbaca", "defegde", "hijhklij"。
# 每个字母最多出现在一个片段中。
# 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。