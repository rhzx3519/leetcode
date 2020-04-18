class Solution(object):
    def readBinaryWatch(self, num):
        """
        :type num: int
        :rtype: List[str]
        """
        res = []
        N = 10
        def dfs(idx, cnt, a):
            if cnt==num:
                hour = int(''.join(map(str, a[:4])), 2)
                minute = int(''.join(map(str, a[4:])), 2)
                if 0<=hour<=11 and 0<=minute<=59:
                    res.append('{}:{:>02}'.format(hour, minute))
                return
            if idx>=N or cnt>num:
                return
            for i in range(idx, N):
                a[i] = 1
                dfs(i+1, cnt+1, a)
                a[i] = 0
        
        a = [0 for _ in range(N)]
        dfs(0, 0, a)
        return res

if __name__ == '__main__':
    su = Solution()
    su.readBinaryWatch(0)                    