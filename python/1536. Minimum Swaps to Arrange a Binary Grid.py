class Solution(object):
    def minSwaps(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        n = len(grid)
        a = [0]*n
        for i in range(n):
            j = n-1
            while j>=0 and grid[i][j]==0:
                j -= 1
            a[i] = n-1-j
        
        print a
        ans = 0
        for i in range(n-1):
            need = n-i-1
            j = i
            while j<n and a[j]<need:
                j += 1
            if j==n: return -1
            tmp = a[j]
            del a[j]
            a.insert(i, tmp)
            ans += j - i
  
        return ans
            
