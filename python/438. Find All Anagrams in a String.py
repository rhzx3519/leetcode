from collections import Counter

class Solution(object):
    def findAnagrams(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """
        res = []
        t_len = len(p)
        mem = Counter(p)
        l = 0
        for r, ch in enumerate(s):
            if mem[ch]>0:
                t_len -= 1
            mem[ch] -= 1

            if t_len==0:
                while mem[s[l]] < 0:
                    mem[s[l]] += 1
                    l += 1

                if r-l+1==len(p):
                    res.append(l)
                t_len += 1
                mem[s[l]] += 1
                l += 1
        return res
            
if __name__ == '__main__':
    s = 'cbaebabacd'
    p = 'abc'
    su = Solution()
    print su.findAnagrams(s, p)                            