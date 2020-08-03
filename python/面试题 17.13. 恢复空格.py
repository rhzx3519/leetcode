#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def respace(self, dictionary, sentence):
        """
        :type dictionary: List[str]
        :type sentence: str
        :rtype: int
        """
        d = set(dictionary)
        n = len(sentence)
        s = sentence
        dp = [n]*(n+1)
        dp[0] = 0
        for i in range(1, n+1):
            dp[i] = dp[i-1] + 1
            for word in dictionary:
                lw = len(word)
                if lw <= i:
                    if word==s[i-lw: i]:
                        dp[i] = min(dp[i], dp[i - lw])
        return dp[n]

# 思路: dp[i] = min(dp[i], dp[i-lw]) if s[i-lw:i] in dictionary else dp[i] = dp[i-1] + 1
# time: O(M*N), space: O(N)


if __name__ == '__main__':
    # dictionary = ["looked","just","like","her","brother"]
    # sentence = "jesslookedjustliketimherbrother"

    dictionary = ['wl']
    sentence = 'lw'
    su = Solution()
    print su.respace(dictionary, sentence)
