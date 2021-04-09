#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def partitionLabels(self, S):
        """
        :type S: str
        :rtype: List[int]
        """
        mem = {}
        for i in range(len(S)):
            mem[S[i]] = i

        ans = []
        l = r = 0
        for i in range(len(S)):
            r = max(r, mem[S[i]])
            if i==r:
                ans.append(r-l+1)
                l = i + 1
        return ans
        

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