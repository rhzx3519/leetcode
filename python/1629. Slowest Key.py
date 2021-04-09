class Solution(object):
    def slowestKey(self, releaseTimes, keysPressed):
        """
        :type releaseTimes: List[int]
        :type keysPressed: str
        :rtype: str
        """
        ans = None
        max_val = 0
        n = len(releaseTimes)
        for i in range(n):
            ch = keysPressed[i]
            t = releaseTimes[i] if i==0 else releaseTimes[i] - releaseTimes[i-1]
            if t > max_val:
                max_val = t
                ans = ch
            elif t == max_val and ch > ans:
                ans = ch
        return ans
