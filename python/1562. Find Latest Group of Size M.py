class Solution(object):
    def findLatestStep(self, arr, m):
        """
        :type arr: List[int]
        :type m: int
        :rtype: int
        """
        n = len(arr)
        link = [-1]*n
        ans = -1
        cnt = 0
        for i in range(n):
            pos = arr[i] - 1
            link[pos] = pos
            L = R = pos
            if pos>0 and link[pos-1]!=-1:
                if (pos-1) - link[pos-1] + 1 == m:  # link[pos-1] -> pos-1 的连续1等于m
                    cnt -= 1
                L = link[pos-1]
            if pos<n-1 and link[pos+1]!=-1:
                if link[pos+1] - (pos+1) + 1 == m: # pos+1 -> link[pos+1] 的连续1等于m
                    cnt -= 1
                R = link[pos+1]
            link[L] = R
            link[R] = L
            if R-L+1==m:
                cnt += 1
            if cnt > 0:
                ans = i + 1
        return ans
        
# 思路：link数组记录连续1的段的左右端点，当pos = arr[i] - 1等于1时，如果pos时右端点，link[pos]表示左端点位置
# 如果pos是左端点， link[pos]表示右端点位置
# time: O(N), space: O(N)