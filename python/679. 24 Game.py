class Solution(object):
    def judgePoint24(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        ZERO = 1e-6

        def dfs(a):
            if not a:
                return False
            if len(a)==1:
                return abs(24-a[0])<=ZERO
            for i, n0 in enumerate(a):
                for j, n1 in enumerate(a):
                    if i==j:
                        continue
                    # 剩余的数
                    b = [num for idx, num in enumerate(a) if idx not in(i, j)] if len(a) > 2 else []
                    if i < j and dfs(b + [n0 + n1]): # 加，i<j避免重复计算
                        return True
                    elif dfs(b + [n0 - n1]): # 减
                        return True
                    elif i < j and dfs(b + [n0 * n1]): # 乘
                        return True
                    elif n1 > ZERO and dfs(b + [float(n0)/n1]): # 除
                        return True
            return False
        
        return dfs(nums)
                    