class Solution(object):
    def knightDialer(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n==1: return 10
        MOD = 10**9 + 7
        a, b, c, d = 1, 4, 2, 2
        for i in range(1, n):
            ta, tb, tc, td = a, b, c, d

            a = tc
            b = 2*tc + 2*td
            c = 2*ta + tb
            d = tb  # 2可以由7和9通过两种方式到达，8可以由1和3通过两种方式到达，所以总的到达数量就是tb
        return (a + b + c + d)%MOD


# {5},a:{0},b:{1,3,7,9},c:{4,6},d:{2,8}