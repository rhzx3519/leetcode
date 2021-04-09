class Solution(object):
    def totalMoney(self, n):
        """
        :type n: int
        :rtype: int
        """
        def foo(n):
            if n<=1: return n
            return n + foo(n-1)
        week = int(n/7)
        print week
        return foo(max(0, week-1))*7 + foo(7)*week + foo(n%7) + week*(n%7)

if __name__ == '__main__':
    n = 10
    su = Solution()
    print su.totalMoney(n)