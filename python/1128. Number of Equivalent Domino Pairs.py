class Solution(object):
    def numEquivDominoPairs(self, dominoes):
        """
        :type dominoes: List[List[int]]
        :rtype: int
        """
        N = len(dominoes)
        ans = 0
        cnt = [0]*100
        for i in range(N):
            dominoes[i].sort()
            idx = dominoes[i][0]*10 + dominoes[i][1]
            ans += cnt[idx]
            cnt[idx] += 1
        return ans