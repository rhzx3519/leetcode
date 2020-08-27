class Solution(object):
    def numSubmat(self, mat):
        """
        :type mat: List[List[int]]
        :rtype: int
        """
        if not mat:
            return 0
        m, n = len(mat), len(mat[0])
        h = [[0]*n for _ in range(m)]
        res = 0
       

        for i in range(m):
            st = []
            cnt = [0]*n
            idx = [-1]*n            
            for j in range(n-1, -1, -1):
                if i==0:
                    h[i][j] = mat[i][j]
                else:
                    h[i][j] = 1 + h[i-1][j] if mat[i][j]==1 else 0
                while st and h[i][st[-1]] > h[i][j]:    # 单调栈，记录左边第一个比i小的位置
                    idx[st.pop()] = j
                st.append(j)

            for j in range(n):
                p = idx[j]
                cnt[j] = h[i][j]*(j - p)
                if p != -1:
                    cnt[j] += cnt[p]
                res += cnt[j]

        return res
# 思路:
# time: O(m*n), space: O(m*n)