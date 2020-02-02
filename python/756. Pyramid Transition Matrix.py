class Solution(object):
    def pyramidTransition(self, bottom, allowed):
        """
        :type bottom: str
        :type allowed: List[str]
        :rtype: bool
        """
        allowed = set(allowed)
        def dfs(bottom, nextStr):
            if len(bottom)==1:
                return True
            if len(bottom) - len(nextStr)>1:
                s = bottom[len(nextStr): len(nextStr)+2]
                for i in range(7):
                    ch = chr(ord('A') + i)
                    if s+ch in allowed and dfs(bottom, nextStr+ch):
                        return True
                return False
            else:
                return dfs(nextStr, "")
        
        return dfs(bottom, "")

if __name__ == '__main__':
    bottom = 'AABA'
    allowed = ["AAA", "AAB", "ABA", "ABB", "BAC"]
    su = Solution()
    print su.pyramidTransition(bottom, allowed)        