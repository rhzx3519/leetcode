class Solution(object):
    def unhappyFriends(self, n, preferences, pairs):
        """
        :type n: int
        :type preferences: List[List[int]]
        :type pairs: List[List[int]]
        :rtype: int
        """
        d = collections.defaultdict(list)
        for x, y in pairs:
            for i in preferences[x]:
                if i==y: break
                d[x].append(i) # 比y更亲近朋友列表
            for j in preferences[y]:
                if j==x: break
                d[y].append(j) # 比x更亲近朋友列表
        ans = 0
        for i, pf in d.iteritems():
            for j in pf:
                if j in d and i in d[j]:
                    ans += 1 # i 不开心
                    # print i
                    break
        return ans