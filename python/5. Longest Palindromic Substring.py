class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        arr = ['#']
        for x in s:
            arr.append(x)
            arr.append('#')
        dp = [1 for x in range(len(arr))]
        far_pos = 0
        ci = 0
        i, j = 0, 0
        while i < len(arr):
            # key is here
            if far_pos < i:
                j = 1
            else:
                j = min(far_pos - i + 1, dp[2*ci - i]) #2*ci-i is the synmetric pos of i against ci
            while i+j < len(arr) and i-j >= 0 and arr[i+j] == arr[i-j]:
                j += 1
            dp[i] = j
            
            if i + dp[i] - 1 > far_pos:
                far_pos = i + dp[i] - 1
                ci = i
            i += 1
        
        res = ''
        max_len = 0
        length = 0
        i = 0
        while i < len(arr):
            length = dp[i] & ~1
            if length > max_len:
                max_len = length
                j = i - length >> 1
                res = s[j:j+max_len]
            i += 2
            
        i = 1
        while i < len(arr):
            length = (dp[i] - 1 & ~1) + 1
            if length > max_len:
                max_len = length
                j = i - length >> 1
                res = s[j:j+max_len]
            i += 2

        return res
           
        
s = Solution()
s.longestPalindrome('')





