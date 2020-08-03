class Solution(object):
    def divisorGame(self, N):
        """
        :type N: int
        :rtype: bool
        """
        f = [False]*(N+5)
        f[1] = False
        f[2] = True
        for i in range(3, N+1):
            for j in range(1, i):
                if i%j==0 and not f[i-j]:
                    f[i] = True
                    break
        return f[N]
        
# 我们定义 f[i]f[i] 表示当前数字 ii 的时候先手是处于必胜态还是必败态，\textit{true}true 表示先手必胜，\textit{false}false 表示先手必败，从前往后递推，根据我们上文的分析，枚举 ii 在 (0, i)(0,i) 中 ii 的因数 jj，看是否存在 f[i-j]f[i−j] 为必败态即可。

