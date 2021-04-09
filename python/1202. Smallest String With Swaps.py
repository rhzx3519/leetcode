import collections

class Solution(object):
    def smallestStringWithSwaps(self, s, pairs):
        """
        :type s: str
        :type pairs: List[List[int]]
        :rtype: str
        """
        mem = {}

        def union(a, b):
            pa = find(a)
            pb = find(b)
            if pa != pb:
                mem[pb] = pa

        def find(a):
            pa = mem.get(a, a)
            if pa != a:
                mem[a] = find(pa)
            return mem.get(a, a)
        
        for a, b in pairs:
            union(a, b)
        
        # print mem
        n = len(s)
        d = collections.defaultdict(list)
        for i in range(n):
            pi = find(i)
            d[pi].append(i)
        # print d

        ans = list(s)
        for k in d.iterkeys():
            t = sorted(d[k], key=lambda x: s[x])
            print t
            for i in range(len(d[k])):
                ans[d[k][i]] = s[t[i]]

        # print ans
        return ''.join(ans)

        

if __name__ == '__main__':
    # s = "dcab"
    # pairs = [[0,3],[1,2],[0,2]]
    s = "cba"
    pairs = [[0,1],[1,2]]
    s = "dcab"
    pairs = [[0,3],[1,2]]
    su = Solution()
    su.smallestStringWithSwaps(s, pairs)


            

        
