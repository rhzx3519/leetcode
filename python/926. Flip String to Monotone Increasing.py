class Solution(object):
    def minFlipsMonoIncr(self, S):
        """
        :type S: str
        :rtype: int
        """
        m = S.count('0')
        t = [m] # m 表示分界点之后的0
        for ch in S:
            if ch=='0': # 分界之前1的数量不变，分界点之后0的数量减1
                m -= 1
            else:
                m += 1 # 分界点之前的1数量加1, 分界点之后0的数量不变
            t.append(m)
        return min(t)


#遍历字符串，找到一个分界点，使得该分界点之前1的个数和分界点之后0的个数之和最小，把分界点之前的1变成0，之后的0变成1
