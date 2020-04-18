#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# 我们首先建立一个可以容纳ASCII的字典mem，然后记录t中元素出现的次数。接着我们遍历s中的元素c，我们将字典mem中的c元素减一，
# 如果s中的元素在t中出现过，那么我们记录此时的t_len = len(t)-1。当t_len==0时，说明此时t中的元素在s中都出现了，
# 我们此时要进入下一步判断。

# 我们要判断mem中s[l]是不是小于0，如果是的话，说明t中不应该包含s[l]，所以我们需要将它剔除出去（也就是mem[s[l++]]++）。
# 否则的话，我们记录此时的l和r。接着我们需要继续找符合条件的位置，我们将l++并且t_len++，将s[l]再加入到mem中去，
# 这样我们就将窗口向右滑动一步
# ————————————————
# 版权声明：本文为CSDN博主「coordinate_blog」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
# 原文链接：https://blog.csdn.net/qq_17550379/article/details/83014880    

from collections import Counter
class Solution(object):
    def minWindow(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        mem = Counter(t)
        t_len = len(t)
        min_l, min_r = 0, float('inf')
        l = 0
        for r, ch in enumerate(s):
            if mem[ch] > 0:
                t_len -= 1
            mem[ch] -= 1

            if t_len==0:
                while mem[s[l]] < 0:
                    mem[s[l]] += 1
                    l += 1
                if r-l < min_r-min_l:
                    min_l, min_r = l, r
                mem[s[l]] += 1
                l += 1
                t_len += 1
        return '' if min_r==float('inf') else s[min_l:min_r+1]

    def minSubstr(self, s):
        # 变种问题：给定一个字符串，找到包含该字符串所有字符的最短子串
        mem = {}
        for ch in s:
            mem[ch] = 1

        t_len = len(mem)
        min_l, min_r = 0, float('inf')
        l = 0
        for r, ch in enumerate(s):
            if mem[ch] > 0:
                t_len -= 1
            mem[ch] -= 1

            if t_len==0:
                while mem[s[l]] < 0:
                    mem[s[l]] += 1
                    l += 1
                if r-l < min_r-min_l:
                    min_l, min_r = l, r
                mem[s[l]] += 1
                l += 1
                t_len += 1
        return '' if min_r==float('inf') else s[min_l:min_r+1]



if __name__ == '__main__':
    s = 'ADOBECODEBANC'
    t = s
    su = Solution()
    print su.minSubstr(s)                    

    

