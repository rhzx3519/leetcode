class Solution(object):
    def balancedString(self, s):
        """
        :type s: str
        :rtype: int
        """
        n = len(s)
        avg = n/4
        cnt = collections.Counter(s)
        l = 0
        ans = n
        for r in range(n):
            cnt[s[r]] -= 1
            # 滑动窗口外的各个字母数量都小于n/4
            while l<n and cnt['Q']<=avg and cnt['W']<=avg and cnt['E']<=avg and cnt['R']<=avg:
                ans = min(ans, r-l+1)
                cnt[s[l]] += 1
                l += 1
        return ans
