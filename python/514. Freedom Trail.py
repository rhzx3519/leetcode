class Solution(object):
    def findRotateSteps(self, ring, key):
        """
        :type ring: str
        :type key: str
        :rtype: int
        """
        n = len(ring)
        mem = collections.defaultdict(list)
        for i, ch in enumerate(ring):
            mem[ch].append(i)
        d = {}

        def dfs(idx, pos):
            if (idx, pos) in d:
                return d[(idx, pos)]
            if idx==len(key):
                return 0
            k = key[idx]
            ans = float('inf')
            for ni in mem[k]:
                rotate = min(len(ring) - abs(ni - pos), abs(ni - pos))
                ans = min(ans, dfs(idx+1, ni) + rotate + 1)
            d[(idx, pos)] = ans
            return ans
            
        return dfs(0, 0)    

# 思路：带备忘录的dfs，
# time: O(N*NlogN), space: O(N + K)        


