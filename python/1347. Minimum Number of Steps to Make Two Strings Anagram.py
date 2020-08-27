import collections

class Solution(object):
    def minSteps(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        s1 = collections.Counter(s)
        t1 = collections.Counter(t)
        pos = neg = 0
        for ch, num in s1.iteritems():
            if num > t1.get(ch, 0):
                neg += num - t1.get(ch, 0)
        for ch, num in t1.iteritems():
            if num > s1.get(ch, 0):
                pos += num - s1.get(ch, 0)
        ans = max(pos, neg)
        print ans
        return ans

if __name__ == '__main__':
    s = "xxyyzz"
    t = "xxyyzz"
    su = Solution()
    su.minSteps(s, t)