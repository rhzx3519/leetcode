class Solution(object):
    def maximalRectangle(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix: return 0
        m, n = len(matrix), len(matrix[0])
        for i in range(n):
            last = 0
            for j in range(m):
                if matrix[j][i]=='1':
                    last += 1
                    matrix[j][i] = last
                else:
                    matrix[j][i]=0
                    last = 0
        
        def maxRectangle(heights):
            st = []
            heights = [0] + heights + [0]
            res = 0
            for i, h in enumerate(heights):
                while st and heights[st[-1]] > h:
                    top = st.pop()
                    res = max(res, heights[top]*(i - st[-1]-1))
                st.append(i)
            return res

        res = 0
        for i in range(m):
            res = max(res, maxRectangle(matrix[i][:]))
        return res
        
if __name__ == '__main__':
    matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
    su = Solution()
    print su.maximalRectangle(matrix)                    