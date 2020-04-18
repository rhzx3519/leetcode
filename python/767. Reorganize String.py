#!usr/bin/env python  
#-*- coding:utf-8 _*-  
from Queue import Queue, PriorityQueue

class Comparable(object):
    """docstring for """
    def __init__(self, ch, fre):
        super(Comparable, self).__init__()
        self.ch = ch
        self.fre = fre

    def __cmp__(self, other):
        if self.fre < other.fre:
            return 1
        elif self.fre == other.fre:
            return 0
        else:
            return -1

    def __str__(self):
        return '{}: {}'.format(self.ch, self.fre)


class Solution(object):
    def reorganizeString(self, S):
        """
        :type S: str
        :rtype: str
        """
        frequence = {}
        for ch in S:
            f = frequence.setdefault(ch, 0)
            frequence[ch] = f+1
        
        max_val = 0
        pq = PriorityQueue()
        for ch, fre in frequence.iteritems():
            pq.put(Comparable(ch, fre))
            max_val = max(max_val, fre)

        if max_val > len(S) - max_val + 1:
            return ""

        arr = []
        while pq.qsize()>=2:
            c1 = pq.get()
            c2 = pq.get()
            arr.append(c1.ch)
            arr.append(c2.ch)
            c1.fre -= 1
            if c1.fre > 0:
                pq.put(c1)
            c2.fre -= 1
            if c2.fre > 0:
                pq.put(c2)



        if pq.qsize() > 0:
            arr.append(pq.get().ch)

        return ''.join([str(i) for i in arr])


if __name__ == '__main__':
    S = "baa"
    su = Solution()
    res = su.reorganizeString(S)                    
    print res

# 根据字母的出现次数将字符串中的所有字符放到一个优先队列中，每次取队头的两个字符组成返回字符串，
# 注意，统计字符出现次数的时候，当最大的字符数量超过总字符串一半时，说明无法组成相邻是不同字符的字符串

# Example 1:

# Input: S = "aab"
# Output: "aba"
# Example 2:

# Input: S = "aaab"
# Output: ""
# Note:

# S will consist of lowercase letters and have length in range [1, 500].

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/reorganize-string
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。        