class Solution(object):
    def numPairsDivisibleBy60(self, time):
        """
        :type time: List[int]
        :rtype: int
        """
        def foo(n):
            if n<=1:
                return n
            return n + foo(n-1)

        ans = 0
        time = [t%60 for t in time]
        mem = collections.Counter(time)
        for t, num in mem.iteritems():
            if t > 30:
                continue
            x = (60 - t) % 60
            if x not in mem:
                continue
            if x == t:
                ans += foo(num-1)
            else:
                ans += mem[t] * mem[x]
        return ans