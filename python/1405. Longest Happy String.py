class Solution(object):
    def longestDiverseString(self, a, b, c):
        """
        :type a: int
        :type b: int
        :type c: int
        :rtype: str
        """
        cnt = [[a, 'a'], [b, 'b'], [c, 'c']]
        ans = ''
        while True:
            for num in sorted(cnt, reverse=True):
                if num[0] <= 0:
                    return ans
                if len(ans)>=2 and ans[-2:]==num[1]*2: # 比较最后两个字符
                    continue
                ans += num[1] # 每次添加一个字符
                num[0] -= 1
                break
        return ans