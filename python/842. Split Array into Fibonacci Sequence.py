class Solution(object):
    def splitIntoFibonacci(self, S):
        """
        :type S: str
        :rtype: List[int]
        """

        def backtrace(idx, ans):
            if idx==len(S):
                return len(ans)>2
            cur = 0
            for i in range(idx, len(S)):
                if i > idx and S[idx]=='0':
                    break
                cur = cur * 10 + (ord(S[i]) - ord('0'))
                if cur > 2**31 - 1:
                    break
                if len(ans) < 2 or ans[-1] + ans[-2]==cur:
                    ans.append(cur)
                    if backtrace(i+1, ans):
                        return True
                    ans.pop()
            return False

        ans = []
        return ans if backtrace(0, ans) else []
                